package oauth

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
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
	if err := o.SetCtx(c).BindUri(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if state, err := c.Cookie(resp.CookieState); err != nil { // 校验state，防止CSRF攻击
		resp.ErrorUnknown(c, err, "get state cookie err")
		return
	} else if state == "" { // todo: state check
		resp.ErrorRequestWithMsg(c, "invalid state")
		return
	}

	c.SetCookie(resp.CookieState, "", -1, "/api/quick/login", "", false, true) // 删除state
	provider, err := o.svc.GetProviderById(in.Tenant.ID, in.ProviderId)
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
	idProvider.SetHttpClient(http.DefaultClient)
	token, err := idProvider.GetToken(in.Code)
	if err != nil {
		resp.ErrorUnknown(c, err, fmt.Sprintf("get %d token err", in.ProviderId))
		return
	}
	if !token.Valid() {
		resp.ErrorForbidden(c, "token is not valid")
		return
	}
	userInfo, err := idProvider.GetUserInfo(token)
	if err != nil {
		resp.ErrorUnknown(c, err, fmt.Sprintf("get %s user info err", in.ProviderId))
		return
	}

	tokenStr, err := o.svc.CreateProviderToken(in.Tenant.App, in.Tenant, userInfo, "")
	if err != nil {
		resp.ErrorUnknown(c, err, "create provider id token err")
		return
	}
	c.SetCookie(resp.CookieIDToken, tokenStr, int(in.Tenant.IDExpire), "/api/quick", "", false, true)
	resp.Success(c) // todo: redirect to next by state
}
