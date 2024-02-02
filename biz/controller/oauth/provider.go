package oauth

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/admin"
	"QuickAuth/biz/service/oauth"
	"QuickAuth/pkg/idp"
	"QuickAuth/pkg/safe"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ProviderLogin	swagger
// @Description	login third provider
// @Tags		login
// @Param		provider	path	string	true	"provider name"
// @Param		next		query	string	false	"next"
// @Success		200
// @Router		/api/quick/login/providers/{providerId} [get]
func (o Controller) ProviderLogin(c *gin.Context) {
	state := safe.RandHex(31)
	c.SetCookie(resp.CookieState, state, 60*5, "/api/quick/login", "", false, true)
}

// ProviderCallback	swagger
// @Description	login third provider callback
// @Tags		login
// @Param		provider	path	string	true	"provider name"
// @Param		code		query	string	true	"code"
// @Param		next		query	string	false	"next"
// @Success		200
// @Router		/api/quick/login/providers/{providerId}/callback [get]
func (o Controller) ProviderCallback(c *gin.Context) {
	var in request.LoginProvider
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := oauth.CheckState(c); err != nil {
		resp.ErrorRequestWithErr(c, err, "invalid state")
		return
	}

	c.SetCookie(resp.CookieState, "", -1, "/api/quick/login", "", false, true) // 删除state
	provider, err := admin.GetProviderById(in.Tenant.ID, in.ProviderId)
	if err != nil {
		resp.ErrorSelect(c, err, "no such provider")
		return
	}

	// 向第三方id系统获取用户信息
	idProvider := idp.GetIdProvider(provider, "")
	if idProvider == nil {
		resp.ErrorRequestWithMsg(c, "provider type "+provider.Type+" does not exist")
		return
	}
	token, err := idProvider.GetToken(in.Code)
	if err != nil {
		resp.ErrorUnknown(c, err, fmt.Sprintf("get %d token err", in.ProviderId))
		return
	}
	if !token.Valid() {
		resp.ErrorForbidden(c, provider.Type+" token is invalid")
		return
	}
	userInfo, err := idProvider.GetUserInfo(token)
	if err != nil {
		resp.ErrorUnknown(c, err, fmt.Sprintf("get %d user info err", in.ProviderId))
		return
	}

	// 登录成功，生成id_token
	tokenStr, err := oauth.CreateProviderToken(in.Tenant.App, in.Tenant, userInfo, "")
	if err != nil {
		resp.ErrorUnknown(c, err, "create provider id token err")
		return
	}
	c.SetCookie(resp.CookieIDToken, tokenStr, int(in.Tenant.IDExpire), "/api/quick", "", false, true)
	if next := c.Query("next"); next != "" {
		c.Redirect(http.StatusFound, next)
		return
	}

	resp.Success(c) // todo: redirect to next by state
}
