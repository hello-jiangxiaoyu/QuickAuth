package controller

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/internal/server/controller/internal"
	"QuickAuth/internal/server/service"
	"QuickAuth/pkg/tools/safe"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
)

type Controller struct {
	internal.Api
	svc *service.Service
}

func NewOAuth2Api(svc *service.Service) Controller {
	return Controller{svc: svc}
}

// @Summary	login a user
// @Schemes
// @Description	login using username and password
// @Tags		login
// @Param		username	formData	string	true	"username"
// @Param		password	formData	string	true	"password"
// @Param		next		query		string	false	"next"
// @Success		302
// @Success		200
// @Router		/api/quick/login [post]
func (o Controller) login(c *gin.Context) {
	var in request.Login
	session := sessions.Default(c)
	su := session.Get("user")
	if su != nil {
		resp.DoNothing(c, "user is already logged in, nothing to do")
		return
	}
	if err := o.SetCtx(c).BindQuery(&in).BindForm(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "init login req err")
		return
	}

	user, err := o.svc.GetUser(&in)
	if err != nil {
		resp.ErrorNotFound(c, err, "no such user")
		return
	}
	if !safe.CheckPasswordHash(in.Password, *user.Password) {
		resp.ErrorForbidden(c, "user name or password is incorrect")
		return
	}

	session.Set("tenant", in.Tenant.Name)
	session.Set("user", user.Username)
	session.Set("userId", user.ID)
	if err = session.Save(); err != nil {
		resp.ErrorSaveSession(c, errors.Wrap(err, "login err"))
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
// @Router		/api/quick/oauth2/auth [get]
func (o Controller) getAuthCode(c *gin.Context) {
	var in request.Auth
	if err := o.SetCtx(c).BindQuery(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "init auth para err")
		return
	}

	session := sessions.Default(c)
	userId, ok := session.Get("userId").(string)
	if !ok || userId == "" {
		resp.ErrorForbidden(c, "invalid user_id")
		return
	}
	if ok, err := o.svc.IsRedirectUriValid(in.ClientID, in.RedirectUri); err != nil {
		resp.ErrorSelect(c, err, "get redirect uri err.")
		return
	} else if !ok {
		resp.ErrorForbidden(c, "Invalid redirect_uri.")
		return
	}

	if in.ResponseType == internal.Oauth2ResponseTypeCode {
		code, state, err := o.svc.CreateAccessCode(in.ClientID, userId)
		if err != nil {
			resp.ErrorSqlModify(c, err, "failed to create access code.")
			return
		}
		query := url.Values{}
		query.Add("code", code)
		session.Set("state", state)
		if err = session.Save(); err != nil {
			resp.ErrorSaveSession(c, errors.Wrap(err, "auth err"))
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

	resp.ErrorRequestWithMsg(c, nil, "Invalid response_type.")
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
// @Router		/v1/oauth2/token [post]
func (o Controller) getToken(c *gin.Context) {
	var in request.Token
	if err := o.SetCtx(c).BindQuery(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "init token req para err")
		return
	}

	client, err := o.svc.GetClientById(in.ClientID)
	if err != nil {
		resp.ErrorRequestWithMsg(c, err, "no such client")
		return
	}
	in.Client = *client
	handler, err := o.getTokenHandler(in.GrantType)
	if err != nil {
		resp.ErrorRequestWithMsg(c, err, err.Error())
		return
	}

	token, err := handler(&in)
	if err != nil {
		switch err {
		case service.ErrorCodeExpired:
			resp.ErrorForbidden(c, err.Error())
		default:
			resp.ErrorUnknown(c, err, "failed to get token.")
		}
		return
	}

	resp.SuccessWithData(c, token) // success!
}
