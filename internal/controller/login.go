package controller

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/pkg/idp"
	"QuickAuth/pkg/tools/safe"
	"QuickAuth/pkg/tools/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

// @Summary	login a user
// @Schemes
// @Description	login using username and password
// @Tags		login
// @Param		username	formData	string	true	"username"
// @Param		password	formData	string	true	"password"
// @Param		next		query		string	false	"next"
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

	user, err := o.svc.GetUserByName(in.Tenant.UserPoolID, in.UserName)
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

// @Summary	logout current user
// @Schemes
// @Description	logout current user
// @Tags		login
// @Success		200
// @Router		/api/quick/logout [get]
func (o Controller) logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		resp.ErrorNoLogin(c)
		return
	}
	session.Delete("user")
	if err := session.Save(); err != nil {
		resp.ErrorSaveSession(c, err)
		return
	}
	resp.Success(c)
}

// @Summary	provider callback
// @Schemes
// @Description	login third provider callback
// @Tags		login
// @Param		provider	path	string	true	"provider name"
// @Param		code		query	string	true	"code"
// @Success		200
// @Router		/api/quick/login/providers/{provider} [get]
func (o Controller) providerCallback(c *gin.Context) {
	var in request.LoginProvider
	if err := o.SetCtx(c).BindUri(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "init provider login err")
		return
	}
	session := sessions.Default(c)
	if state := session.Get("state"); state == nil { // todo: state check
		resp.ErrorRequestWithMsg(c, nil, "invalid state")
		return
	}
	session.Delete("state")
	utils.DeferErr(session.Save)
	provider, err := o.svc.GetProviderByType(in.Tenant.ID, in.ProviderName)
	if err != nil {
		resp.ErrorRequestWithMsg(c, err, "no such provider")
		return
	}

	idProvider := idp.GetIdProvider(provider.Type, provider.ClientID, provider.ClientSecret, "")
	if idProvider == nil {
		resp.ErrorRequest(c, nil, "no such provider")
		return
	}
	idProvider.SetHttpClient(http.DefaultClient)
	token, err := idProvider.GetToken(in.Code)
	if err != nil {
		resp.ErrorUnknown(c, err, fmt.Sprintf("get %s token err", in.ProviderName))
		return
	}
	if !token.Valid() {
		resp.ErrorForbidden(c, "token is not valid")
		return
	}

	userInfo, err := idProvider.GetUserInfo(token)
	if err != nil {
		resp.ErrorUnknown(c, err, fmt.Sprintf("get %s user info err", in.ProviderName))
		return
	}
	session.Set("userId", userInfo.Id)
	resp.Success(c) // todo: redirect to next by state
}
