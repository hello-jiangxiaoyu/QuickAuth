package oauth

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/response"
	"QuickAuth/internal/global"
	"QuickAuth/internal/server/controller/internal"
	"QuickAuth/internal/server/service"
	"QuickAuth/internal/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"net/url"
)

func NewOauth2Router(e *gin.Engine) {
	var o oauth
	r := e.Group("/quick/v1")
	{
		r.GET("/.well-known/openid-configuration", o.getOIDC)
		r.GET("/.well-known/jwks.json", o.getJwks)
		r.POST("/login", o.login)
	}
	e.GET("/v1/health", func(c *gin.Context) { c.String(http.StatusOK, "ok") })
}

type oauth struct{}

// @Summary	login a user
// @Schemes
// @Description	login using username and password
// @Tags		login
// @Param		username	formData	string	true	"username"
// @Param		password	formData	string	true	"password"
// @Param		next		query		string	false	"next"
// @Success		302
// @Success		200
// @Router		/v1/login [post]
func (o *oauth) login(c *gin.Context) {
	var in request.Login
	session := sessions.Default(c)
	su := session.Get("user")
	if su != nil {
		response.DoNothing(c, "user is already logged in, nothing to do")
		return
	}
	if err := internal.NewApi(c).BindQuery(&in).BindForm(&in).SetTenant(&in.Tenant).Error; err != nil {
		response.ErrorRequest(c)
		global.Log.Error("api err: ", zap.Error(err))
		return
	}

	user, err := service.GetUser(&in)
	if err != nil {
		response.ErrorNotFound(c, "no such user")
		global.Log.Error("login err, no such user: ", zap.Error(err))
		return
	}
	if !utils.CheckPasswordHash(in.Password, *user.Password) {
		response.ErrorForbidden(c, "user name or password is incorrect")
		return
	}

	session.Set("tenant", in.Tenant.Name)
	session.Set("user", user.Username)
	session.Set("userId", user.ID)
	if err = session.Save(); err != nil {
		response.ErrorSaveSession(c)
		global.Log.Error("login save session err: ", zap.Error(err))
		return
	}
	if next := c.Query("next"); next != "" {
		c.Redirect(http.StatusFound, next)
		return
	}
	c.Status(http.StatusOK)
}

// @Summary	oauth2 authorize
// @Schemes
// @Description	oauth2 authorize
// @Tags		oauth2
// @Param		client_id		query	string	true	"client_id"
// @Param		scope			query	string	true	"scope"
// @Param		response_type	query	string	true	"response_type"
// @Param		redirect_uri	query	string	true	"redirect_uri"
// @Param		state			query	string	false	"state"
// @Param		nonce			query	string	false	"nonce"
// @Success		302
// @Success		200
// @Router		/v1/oauth2/auth [get]
func (o *oauth) getAuthCode(c *gin.Context) {
	var in request.Auth
	if err := internal.NewApi(c).BindQuery(&in).SetTenant(&in.Tenant).Error; err != nil {
		response.ErrorRequest(c)
		global.Log.Error("api err: ", zap.Error(err))
		return
	}

	session := sessions.Default(c)
	userId, ok := session.Get("userId").(string)
	if !ok || userId == "" {
		response.ErrorForbidden(c, "invalid user_id")
		return
	}
	if ok = service.IsRedirectUriValid(in.ClientID, in.RedirectUri); !ok {
		response.ErrorForbidden(c, "Invalid redirect_uri.")
		return
	}

	if in.ResponseType == internal.Oauth2ResponseTypeCode {
		code, state, err := service.CreateAccessCode(in.ClientID, userId)
		if err != nil {
			response.ErrorSqlModify(c, "failed to create access code")
			global.Log.Error("get access code err: " + err.Error())
			return
		}
		query := url.Values{}
		query.Add("code", code)
		session.Set("state", state)
		if err = session.Save(); err != nil {
			response.ErrorSaveSession(c)
			global.Log.Error("oauth save session err: ", zap.Error(err))
			return
		}
		location := fmt.Sprintf("%s?%s", in.RedirectUri, query.Encode())
		c.Redirect(http.StatusFound, location)
		return
	}

	if in.ResponseType == internal.Oauth2ResponseTypeToken {
		token := ""
		query := url.Values{}
		query.Add("access_token", token)
		location := fmt.Sprintf("%s?%s", in.RedirectUri, query.Encode())
		c.Redirect(http.StatusFound, location)
		return
	}

	response.ErrorRequestWithMsg(c, "Invalid response_type.")
}

// @Summary	oauth2 token
// @Schemes
// @Description	oauth2 token
// @Tags		oauth2
// @Param		client_id		query		string	true	"client_id"
// @Param		client_secret	query		string	false	"client_secret"
// @Param		code			query		string	false	"code"
// @Param		grant_type		query		string	true	"grant_type"
// @Param		redirect_uri	query		string	false	"redirect_uri"
// @Param		state			query		string	false	"state"
// @Param		nonce			query		string	false	"nonce"
// @Success		200
// @Router		/v1/oauth2/token [get]
func (o *oauth) getToken(c *gin.Context) {
	var in request.Token
	if err := internal.NewApi(c).BindQuery(&in).SetTenant(&in.Tenant).Error; err != nil {
		response.ErrorRequest(c)
		global.Log.Error("api err: ", zap.Error(err))
		return
	}

	client, err := service.GetClientById(in.ClientID)
	if err != nil {
		response.ErrorRequestWithMsg(c, "no such client")
		global.Log.Error("get client err: ", zap.Error(err))
		return
	}
	in.Client = *client
	handler, err := getTokenHandler(in.GrantType)
	if err != nil {
		response.ErrorRequestWithMsg(c, err.Error())
		return
	}

	token, err := handler(&in)
	if err != nil {
		switch err {
		case service.ErrorCodeExpired:
			response.ErrorForbidden(c, err.Error())
		default:
			response.ErrorUnknown(c, "failed to get token.")
		}
		global.Log.Error("get token err: ", zap.Error(err))
		return
	}

	response.SuccessWithData(c, token) // success!
}
