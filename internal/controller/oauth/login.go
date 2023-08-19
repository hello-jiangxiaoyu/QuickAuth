package oauth

import (
	"QuickAuth/internal/endpoint/model"
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/pkg/idp"
	"QuickAuth/pkg/safe"
	"QuickAuth/pkg/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login	swagger
// @Summary	login a user
// @Schemes
// @Description	login using username and password
// @Tags		login
// @Param		username	formData	string	true	"username"
// @Param		password	formData	string	true	"password"
// @Param		next		query		string	false	"next"
// @Success		200
// @Router		/api/quick/login [post]
func (o Controller) Login(c *gin.Context) {
	var in request.Login
	if cookie, err := c.Cookie(resp.IDToken); err == nil && cookie != "" {
		resp.DoNothing(c, "user is already logged in, nothing to do")
		return
	}
	if err := o.SetCtx(c).BindForm(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	// 用户名和密码检查
	user, err := o.svc.GetUserByName(in.Tenant.UserPoolID, in.UserName)
	if err != nil {
		resp.ErrorNotFound(c, err, "no such user")
		return
	}
	if !safe.CheckPasswordHash(in.Password, user.Password) {
		resp.ErrorForbidden(c, "user name or password is incorrect")
		return
	}

	// 生成包含用户信息的id_token
	token, err := o.svc.CreateIdToken(in.Tenant.App, in.Tenant, *user, "")
	if err != nil {
		resp.ErrorUnknown(c, err, "create id token err")
		return
	}
	c.SetCookie(resp.IDToken, token, int(in.Tenant.IDExpire), "/api/quick", "", false, true)
	if next := c.Query("next"); next != "" {
		c.Redirect(http.StatusFound, next)
		return
	}
	resp.Success(c)
}

// Logout	swagger
// @Summary	logout current user
// @Schemes
// @Description	logout current user
// @Tags		login
// @Success		200
// @Router		/api/quick/logout [get]
func (o Controller) Logout(c *gin.Context) {
	c.SetCookie(resp.IDToken, "", -1, "/", "", false, true)
	resp.Success(c)
}

// ProviderCallback	swagger
// @Summary	provider callback
// @Schemes
// @Description	login third provider callback
// @Tags		login
// @Param		provider	path	string	true	"provider name"
// @Param		code		query	string	true	"code"
// @Success		200
// @Router		/api/quick/login/providers/{provider} [get]
func (o Controller) ProviderCallback(c *gin.Context) {
	var in request.LoginProvider
	if err := o.SetCtx(c).BindUri(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
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
		resp.ErrorSelect(c, err, "no such provider")
		return
	}

	idProvider := idp.GetIdProvider(provider.Type, provider.ClientID, provider.ClientSecret, "")
	if idProvider == nil {
		resp.ErrorRequestWithMsg(c, nil, "no such provider")
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
	if err = session.Save(); err != nil {
		resp.ErrorSaveSession(c, err)
		return
	}

	resp.Success(c) // todo: redirect to next by state
}

// Register	swagger
// @Summary	login a user
// @Schemes
// @Description	login using username and password
// @Tags		login
// @Param		username	formData	string	true	"username"
// @Param		password	formData	string	true	"password"
// @Param		next		query		string	false	"next"
// @Success		200
// @Router		/api/quick/register [post]
func (o Controller) Register(c *gin.Context) {
	var in request.Login
	if err := o.SetCtx(c).BindForm(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	var err error
	in.Password, err = safe.HashPassword(in.Password)
	if err != nil {
		resp.ErrorUnknown(c, err, "hash password err")
		return
	}

	user, err := o.svc.CreateUser(&model.User{
		UserPoolID: in.Tenant.UserPoolID,
		Username:   in.UserName,
		Password:   in.Password,
	})
	if err != nil {
		resp.ErrorCreate(c, err, "create user err")
		return
	}

	resp.SuccessWithData(c, user.Dto())
}
