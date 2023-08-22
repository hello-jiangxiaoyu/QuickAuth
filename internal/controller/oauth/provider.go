package oauth

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/pkg/idp"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
	if state, err := c.Cookie(resp.CookieState); err != nil { // 校验state，防止CSRF攻击
		resp.ErrorUnknown(c, err, "get state cookie err")
		return
	} else if state == "" { // todo: state check
		resp.ErrorRequestWithMsg(c, "invalid state")
		return
	}

	c.SetCookie(resp.CookieState, "", -1, "/api/quick/login", "", false, true)
	provider, err := o.svc.GetProviderByType(in.Tenant.ID, in.ProviderName)
	if err != nil {
		resp.ErrorSelect(c, err, "no such provider")
		return
	}

	// 向第三方id系统获取用户信息
	idProvider := idp.GetIdProvider(provider, "")
	if idProvider == nil {
		resp.ErrorRequestWithMsg(c, "no such provider")
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

	tokenStr, err := o.svc.CreateProviderToken(in.Tenant.App, in.Tenant, userInfo, "")
	if err != nil {
		resp.ErrorUnknown(c, err, "create provider id token err")
		return
	}
	c.SetCookie(resp.CookieIDToken, tokenStr, int(in.Tenant.IDExpire), "/api/quick", "", false, true)
	resp.Success(c) // todo: redirect to next by state
}
