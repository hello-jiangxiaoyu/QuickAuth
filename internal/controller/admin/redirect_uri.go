package admin

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// ListRedirectUri	swagger
// @Summary	get redirect uri list
// @Schemes
// @Description	get redirect uri list
// @Tags		tenant
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		vhost		header	string	false	"tenant host"
// @Success		200		{object}	[]string
// @Router		/api/quick/redirect-uri [get]
func (a Route) ListRedirectUri(c *gin.Context) {
	var in request.RedirectUriReq
	if err := a.SetCtx(c).BindUri(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, true)
		return
	}

	uris, err := a.svc.ListRedirectUri(in.Tenant.ID)
	if err != nil {
		resp.ErrorSelect(c, err, "list redirect uri err", true)
		return
	}
	resp.SuccessArrayData(c, len(uris), uris)
}

// CreateRedirectUri	swagger
// @Summary	create app redirect uri
// @Schemes
// @Description	create app redirect uri
// @Tags		tenant
// @Param		X-User-ID	header	string					false	"user id"
// @Param		X-Pool-ID	header	string					false	"user pool id"
// @Param		vhost		header	string					false	"tenant host"
// @Param		bd			body	request.RedirectUriReq	true	"body"
// @Success		200
// @Router		/api/quick/redirect-uri [post]
func (a Route) CreateRedirectUri(c *gin.Context) {
	var in request.RedirectUriReq
	if err := a.SetCtx(c).BindUriAndJson(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := a.svc.CreateRedirectUri(in.Tenant.ID, in.Uri); err != nil {
		resp.ErrorUpdate(c, err, "create redirect uri err")
		return
	}
	resp.Success(c)
}

// ModifyRedirectUri	swagger
// @Summary	modify app redirect uri
// @Schemes
// @Description	modify app
// @Tags		tenant
// @Param		X-User-ID	header	string					false	"user id"
// @Param		X-Pool-ID	header	string					false	"user pool id"
// @Param		vhost		header	string					false	"tenant host"
// @Param		uriId		path	string					true	"uri id"
// @Param		bd			body	request.RedirectUriReq	true	"body"
// @Success		200
// @Router		/api/quick/redirect-uri/{uriId} [put]
func (a Route) ModifyRedirectUri(c *gin.Context) {
	var in request.RedirectUriReq
	if err := a.SetCtx(c).BindUriAndJson(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := a.svc.ModifyRedirectUri(in.Tenant.ID, in.UriId, in.Uri); err != nil {
		resp.ErrorUpdate(c, err, "modify redirect uri err")
		return
	}
	resp.Success(c)
}

// DeleteRedirectUri	swagger
// @Summary	delete app
// @Schemes
// @Description	delete app
// @Tags		tenant
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		vhost		header	string	false	"tenant host"
// @Param		uri			path	string	true	"uri name"
// @Success		200
// @Router		/api/quick/redirect-uri/{uri} [delete]
func (a Route) DeleteRedirectUri(c *gin.Context) {
	var in request.RedirectUriReq
	if err := a.SetCtx(c).BindUri(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	uri := c.Param("uri")
	if err := a.svc.DeleteRedirectUri(in.Tenant.ID, uri); err != nil {
		resp.ErrorUpdate(c, err, "delete redirect uri err")
		return
	}
	resp.Success(c)
}
