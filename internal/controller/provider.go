package controller

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// @Summary	provider info
// @Schemes
// @Description	list provider info
// @Tags		provider
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		vhost		header	string	false	"tenant host"
// @Success		200
// @Router		/api/quick/providers [get]
func (o Controller) listProvider(c *gin.Context) {
	var in request.ProviderReq
	if err := o.SetCtx(c).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid provider request param")
		return
	}
	providers, err := o.svc.GetLoginProviderInfo(in.Tenant.ID)
	if err != nil {
		resp.ErrorSelect(c, err, "get provider list err")
		return
	}

	resp.SuccessArray(c, len(providers), providers)
}

// @Summary	get provider details
// @Schemes
// @Description	get provider details
// @Tags		provider
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		vhost		header	string	false	"tenant host"
// @Param		providerId	path	integer	true	"provider id"
// @Success		200
// @Router		/api/quick/providers/{providerId} [get]
func (o Controller) getProvider(c *gin.Context) {
	var in request.ProviderReq
	if err := o.SetCtx(c).BindUri(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid provider request param")
		return
	}
	provider, err := o.svc.GetProvider(in.Tenant.ID, in.ProviderId)
	if err != nil {
		resp.ErrorSelect(c, err, "get provider err")
		return
	}

	resp.SuccessWithData(c, provider)
}

// @Summary	create provider
// @Schemes
// @Description	create provider
// @Tags		provider
// @Param		X-User-ID	header	string				false	"user id"
// @Param		X-Pool-ID	header	string				false	"user pool id"
// @Param		vhost		header	string				false	"tenant host"
// @Param		bd			body	request.ProviderReq	true	"body"
// @Success		200
// @Router		/api/quick/providers [post]
func (o Controller) createProvider(c *gin.Context) {
	var in request.ProviderReq
	if err := o.SetCtx(c).BindJson(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid provider request param")
		return
	}

	provider, err := o.svc.CreateProvider(in.ToModel())
	if err != nil {
		resp.ErrorSelect(c, err, "create provider err")
		return
	}

	resp.SuccessWithData(c, provider)
}

// @Summary	modify provider
// @Schemes
// @Description	modify provider
// @Tags		provider
// @Param		providerId	path	integer				true	"provider id"
// @Param		vhost		header	string				false	"tenant host"
// @Param		bd			body	request.ProviderReq	true	"body"
// @Success		200
// @Router		/api/quick/providers/{providerId} [put]
func (o Controller) modifyProvider(c *gin.Context) {
	var in request.ProviderReq
	if err := o.SetCtx(c).BindUriAndJson(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid provider request param")
		return
	}

	if err := o.svc.ModifyProvider(in.ProviderId, in.ToModel()); err != nil {
		resp.ErrorSelect(c, err, "modify provider err")
		return
	}

	resp.Success(c)
}

// @Summary	delete provider
// @Schemes
// @Description	delete provider
// @Tags		provider
// @Param		providerId	path	integer	true	"provider id"
// @Param		vhost		header	string	false	"tenant host"
// @Success		200
// @Router		/api/quick/providers/{providerId} [delete]
func (o Controller) deleteProvider(c *gin.Context) {
	var in request.ProviderReq
	if err := o.SetCtx(c).BindUri(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid provider request param")
		return
	}
	if err := o.svc.DeleteProvider(in.Tenant.ID, in.ProviderId); err != nil {
		resp.ErrorSelect(c, err, "delete provider err")
		return
	}

	resp.Success(c)
}
