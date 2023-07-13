package controller

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// @Summary	get redirect uri list
// @Schemes
// @Description	get redirect uri list
// @Tags		app
// @Param		appId		path		string	true	"app id"
// @Param		tenantId	path		string	true	"tenant id"
// @Success		200			{object}	[]string
// @Router		/api/quick/apps/{appId}/tenants/{tenantId}/redirect-uri [get]
func (o Controller) listRedirectUri(c *gin.Context) {
	uris, err := o.svc.ListRedirectUri(c.Param("appId"), c.Param("tenantId"))
	if err != nil {
		resp.ErrorSelect(c, err, "list redirect uri err")
		return
	}
	resp.SuccessArray(c, len(uris), uris)
}

// @Summary	create app redirect uri
// @Schemes
// @Description	create app redirect uri
// @Tags		app
// @Param		appId		path		string	true	"app id"
// @Param		tenantId	path		string	true	"tenant id"
// @Param		bd			body	request.RedirectUriReq	true	"body"
// @Success		200
// @Router		/api/quick/apps/{appId}/redirect-uri [post]
func (o Controller) createRedirectUri(c *gin.Context) {
	var in request.RedirectUriReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init redirect uri req err")
		return
	}

	if err := o.svc.CreateRedirectUri(in.AppId, in.TenantId, in.Uri); err != nil {
		resp.ErrorUnknown(c, err, "create redirect uri err")
		return
	}
	resp.Success(c)
}

// @Summary	modify app redirect uri
// @Schemes
// @Description	modify app
// @Tags		app
// @Param		appId		path	string					true	"app id"
// @Param		tenantId	path	string					true	"tenant id"
// @Param		uriId		path	string					true	"uri id"
// @Param		bd			body	request.RedirectUriReq	true	"body"
// @Success		200
// @Router		/api/quick/apps/{appId}/redirect-uri/{uriId} [put]
func (o Controller) modifyRedirectUri(c *gin.Context) {
	var in request.RedirectUriReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init redirect uri req err")
		return
	}
	if err := o.svc.ModifyRedirectUri(in.AppId, in.TenantId, in.UriId, in.Uri); err != nil {
		resp.ErrorUnknown(c, err, "modify redirect uri err")
		return
	}
	resp.Success(c)
}

// @Summary	delete app
// @Schemes
// @Description	delete app
// @Tags		app
// @Param		appId		path	string	true	"app id"
// @Param		tenantId	path	string	true	"tenant id"
// @Param		uri			path	string	true	"uri name"
// @Success		200
// @Router		/api/quick/apps/{appId}/redirect-uri/{uri} [delete]
func (o Controller) deleteRedirectUri(c *gin.Context) {
	appId := c.Param("appId")
	tenantId := c.Param("tenantId")
	uri := c.Param("uri")
	if err := o.svc.DeleteRedirectUri(appId, tenantId, uri); err != nil {
		resp.ErrorUnknown(c, err, "delete redirect uri err")
		return
	}
	resp.Success(c)
}
