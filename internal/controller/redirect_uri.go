package controller

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// @Summary	get redirect uri list
// @Schemes
// @Description	get redirect uri list
// @Tags		tenant
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		vhost		header	string	false	"tenant host"
// @Success		200		{object}	[]string
// @Router		/api/quick/redirect-uri [get]
func (o Controller) listRedirectUri(c *gin.Context) {
	var in request.RedirectUriReq
	if err := o.SetCtx(c).BindUri(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid redirect-uri request param", true)
		return
	}

	uris, err := o.svc.ListRedirectUri(in.Tenant.ID)
	if err != nil {
		resp.ErrorSelect(c, err, "list redirect uri err", true)
		return
	}
	resp.SuccessArray(c, len(uris), uris)
}

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
func (o Controller) createRedirectUri(c *gin.Context) {
	var in request.RedirectUriReq
	if err := o.SetCtx(c).BindUriAndJson(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid redirect-uri request param")
		return
	}

	if err := o.svc.CreateRedirectUri(in.Tenant.ID, in.Uri); err != nil {
		resp.ErrorUnknown(c, err, "create redirect uri err")
		return
	}
	resp.Success(c)
}

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
func (o Controller) modifyRedirectUri(c *gin.Context) {
	var in request.RedirectUriReq
	if err := o.SetCtx(c).BindUriAndJson(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid redirect-uri request param")
		return
	}
	if err := o.svc.ModifyRedirectUri(in.Tenant.ID, in.UriId, in.Uri); err != nil {
		resp.ErrorUnknown(c, err, "modify redirect uri err")
		return
	}
	resp.Success(c)
}

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
func (o Controller) deleteRedirectUri(c *gin.Context) {
	var in request.RedirectUriReq
	if err := o.SetCtx(c).BindUri(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid redirect-uri request param")
		return
	}

	uri := c.Param("uri")
	if err := o.svc.DeleteRedirectUri(in.Tenant.ID, uri); err != nil {
		resp.ErrorUnknown(c, err, "delete redirect uri err")
		return
	}
	resp.Success(c)
}
