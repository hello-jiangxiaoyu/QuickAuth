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
// @Param		providerId	path	string	true	"app id"
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

// @Summary	get provider details
// @Schemes
// @Description	get provider details
// @Tags		provider
// @Param		providerId	path	string				true	"app id"
// @Param		bd			body	request.ProviderReq	true	"body"
// @Success		200
// @Router		/api/quick/providers/{providerId} [post]
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

// @Summary	get provider details
// @Schemes
// @Description	get provider details
// @Tags		provider
// @Param		providerId	path	string				true	"app id"
// @Param		bd			body	request.ProviderReq	true	"body"
// @Success		200
// @Router		/api/quick/providers/{providerId} [put]
func (o Controller) modifyProvider(c *gin.Context) {
	var in request.ProviderReq
	if err := o.SetCtx(c).BindJson(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid provider request param")
		return
	}

	if err := o.svc.ModifyProvider(in.ToModel()); err != nil {
		resp.ErrorSelect(c, err, "modify provider err")
		return
	}

	resp.Success(c)
}

// @Summary	get provider details
// @Schemes
// @Description	get provider details
// @Tags		provider
// @Param		providerId	path	string	true	"app id"
// @Success		200
// @Router		/api/quick/providers/{providerId} [delete]
func (o Controller) deleteProvider(c *gin.Context) {
	var in request.ProviderReq
	if err := o.SetCtx(c).BindUriAndJson(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid provider request param")
		return
	}
	if err := o.svc.DeleteProvider(in.Tenant.ID, in.ProviderId); err != nil {
		resp.ErrorSelect(c, err, "delete provider err")
		return
	}

	resp.Success(c)
}
