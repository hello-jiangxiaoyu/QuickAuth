package controller

import (
	"QuickAuth/internal/controller/internal"
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/internal/service"
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

// @Summary	oauth2 authorize
// @Schemes
// @Description	oauth2 authorize
// @Tags		oidc
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
		resp.ErrorRequest(c, err, "invalid auth request param")
		return
	}

	session := sessions.Default(c)
	userId, ok := session.Get("userId").(int64)
	if !ok || userId == 0 {
		resp.ErrorForbidden(c, "invalid user_id")
		return
	}
	if ok, err := o.svc.IsRedirectUriValid(in.ClientID, in.Tenant.ID, in.RedirectUri); err != nil {
		resp.ErrorSelect(c, err, "no such uri.")
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
func (o Controller) getToken(c *gin.Context) {
	var in request.Token
	if err := o.SetCtx(c).BindQuery(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid token request param")
		return
	}

	app, err := o.svc.GetApp(in.ClientID)
	if err != nil {
		resp.ErrorRequestWithMsg(c, err, "no such app")
		return
	}
	in.App = *app
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
