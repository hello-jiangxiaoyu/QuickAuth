package admin

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/admin"
	"github.com/gin-gonic/gin"
)

// ListRedirectUri
// @Summary	get redirect uri list
// @Tags	tenant
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		header	string	false	"tenant host"
// @Success	200		{object}	[]string
// @Router	/api/quick/redirect-uri [get]
func (a Route) ListRedirectUri(c *gin.Context) {
	var in request.RedirectUriReq
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	uris, err := admin.ListRedirectUri(in.Tenant.ID)
	if err != nil {
		resp.ErrorSelect(c, err, "list redirect uri err")
		return
	}
	resp.SuccessArrayData(c, len(uris), uris)
}

// CreateRedirectUri
// @Summary	create app redirect uri
// @Tags	tenant
// @Param	X-User-ID	header	string					false	"user id"
// @Param	X-Pool-ID	header	string					false	"user pool id"
// @Param	vhost		header	string					false	"tenant host"
// @Param	bd			body	request.RedirectUriReq	true	"body"
// @Success	200
// @Router	/api/quick/redirect-uri [post]
func (a Route) CreateRedirectUri(c *gin.Context) {
	var in request.RedirectUriReq
	if err := internal.BindUriAndJson(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := admin.CreateRedirectUri(in.Tenant.ID, in.Uri); err != nil {
		resp.ErrorUpdate(c, err, "create redirect uri err")
		return
	}
	resp.Success(c)
}

// ModifyRedirectUri
// @Summary	modify app
// @Tags	tenant
// @Param	X-User-ID	header	string					false	"user id"
// @Param	X-Pool-ID	header	string					false	"user pool id"
// @Param	vhost		header	string					false	"tenant host"
// @Param	uriId		path	string					true	"uri id"
// @Param	bd			body	request.RedirectUriReq	true	"body"
// @Success	200
// @Router	/api/quick/redirect-uri/{uriId} [put]
func (a Route) ModifyRedirectUri(c *gin.Context) {
	var in request.RedirectUriReq
	if err := internal.BindUriAndJson(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := admin.ModifyRedirectUri(in.Tenant.ID, in.UriId, in.Uri); err != nil {
		resp.ErrorUpdate(c, err, "modify redirect uri err")
		return
	}
	resp.Success(c)
}

// DeleteRedirectUri
// @Summary	delete app
// @Tags	tenant
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		header	string	false	"tenant host"
// @Param	uri			path	string	true	"uri name"
// @Success	200
// @Router	/api/quick/redirect-uri/{uri} [delete]
func (a Route) DeleteRedirectUri(c *gin.Context) {
	var in request.RedirectUriReq
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	uri := c.Param("uri")
	if err := admin.DeleteRedirectUri(in.Tenant.ID, uri); err != nil {
		resp.ErrorUpdate(c, err, "delete redirect uri err")
		return
	}
	resp.Success(c)
}
