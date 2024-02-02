package oauth

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/admin"
	"QuickAuth/biz/service/oauth"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"net/url"
)

type Controller struct {
}

func NewOAuth2Route() Controller {
	return Controller{}
}

// GetAuthCode	swagger
// @Description	oauth2 authorize
// @Tags		oidc
// @Param		client_id		query	string	true	"client_id"
// @Param		scope			query	string	true	"scope"
// @Param		response_type	query	string	true	"response_type"
// @Param		redirect_uri	query	string	true	"redirect_uri"
// @Param		state			query	string	false	"state"
// @Param		nonce			query	string	false	"nonce"
// @Success		302
// @Router		/api/quick/oauth2/auth [get]
func (o Controller) GetAuthCode(c *gin.Context) {
	var in request.Auth
	var user jwt.MapClaims
	if err := internal.New(c).BindQuery(&in).SetUser(&user).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if ok, err := admin.IsRedirectUriValid(in.ClientID, in.Tenant.ID, in.RedirectUri); err != nil {
		resp.ErrorSelect(c, err, "no such uri.")
		return
	} else if !ok {
		resp.ErrorForbidden(c, "Invalid redirect_uri.")
		return
	}

	userId, err := user.GetSubject()
	if err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if in.ResponseType == internal.Oauth2ResponseTypeCode {
		code, err := oauth.CreateAccessCode(in.ClientID, userId)
		if err != nil {
			resp.ErrorUpdate(c, err, "failed to create access code.")
			return
		}
		query := url.Values{}
		query.Add("code", code)
		c.Redirect(http.StatusFound, fmt.Sprintf("%s?%s", in.RedirectUri, query.Encode()))
		return
	}

	if in.ResponseType == internal.Oauth2ResponseTypeToken {
		token := ""
		query := url.Values{}
		query.Add("access_token", token)
		c.Redirect(http.StatusFound, fmt.Sprintf("%s?%s", in.RedirectUri, query.Encode()))
		return
	}

	resp.ErrorRequestWithMsg(c, "Invalid response_type.")
}

// GetToken	swagger
// @Description	oauth2 token
// @Tags		oidc
// @Param		client_id		query		string	true	"client_id"
// @Param		client_secret	query		string	false	"client_secret"
// @Param		code			query		string	false	"code"
// @Param		grant_type		query		string	true	"grant_type"
// @Param		redirect_uri	query		string	false	"redirect_uri"
// @Param		state			query		string	false	"state"
// @Param		nonce			query		string	false	"nonce"
// @Success		200
// @Router		/api/quick/oauth2/token [post]
func (o Controller) GetToken(c *gin.Context) {
	var in request.Token
	if err := internal.New(c).BindQuery(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	app, err := admin.GetApp(in.ClientID)
	if err != nil {
		resp.ErrorRequestWithErr(c, err, "no such app")
		return
	}
	in.App = *app
	handler, err := o.getTokenHandler(in.GrantType)
	if err != nil {
		resp.ErrorRequestWithErr(c, err, err.Error())
		return
	}

	token, err := handler(&in)
	if errors.Is(err, oauth.ErrorCodeExpired) {
		resp.ErrorForbidden(c, err.Error())
		return
	} else if err != nil {
		resp.ErrorUnknown(c, err, "failed to get token.")
		return
	}

	resp.SuccessWithData(c, token) // success!
}
