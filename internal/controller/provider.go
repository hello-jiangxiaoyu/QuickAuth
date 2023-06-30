package controller

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/pkg/model"
	"github.com/gin-gonic/gin"
)

// @Summary	provider info
// @Schemes
// @Description	list provider info
// @Tags		provider
// @Success		200
// @Router		/api/quick/providers [get]
func (o Controller) listProvider(c *gin.Context) {
	var tenant model.Tenant
	if err := o.SetCtx(c).SetTenant(&tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "init provider err")
		return
	}
	providers, err := o.svc.GetLoginProviderInfo(tenant.ID)
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
// @Param		providerId	path	string	true	"client id"
// @Success		200
// @Router		/api/quick/providers/{providerId} [get]
func (o Controller) getProvider(c *gin.Context) {
	var tenant model.Tenant
	if err := o.SetCtx(c).SetTenant(&tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "init provider err")
		return
	}
	provider, err := o.svc.GetProviderByType(tenant.ID, c.Param("providerId"))
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
// @Param		providerId	path	string				true	"client id"
// @Param		bd			body	request.ProviderReq	true	"body"
// @Success		200
// @Router		/api/quick/providers/{providerId} [post]
func (o Controller) createProvider(c *gin.Context) {
	var in request.ProviderReq
	var tenant model.Tenant
	if err := o.SetCtx(c).BindJson(&in).SetTenant(&tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "init provider err")
		return
	}
	in.TenantID = tenant.ID
	provider, err := o.svc.CreateProvider(in.ToModel())
	provider.TenantID = tenant.ID
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
// @Param		providerId	path	string				true	"client id"
// @Param		bd			body	request.ProviderReq	true	"body"
// @Success		200
// @Router		/api/quick/providers/{providerId} [put]
func (o Controller) modifyProvider(c *gin.Context) {
	var tenant model.Tenant
	var in request.ProviderReq
	if err := o.SetCtx(c).BindJson(&in).SetTenant(&tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "init provider err")
		return
	}
	in.TenantID = tenant.ID
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
// @Param		providerId	path	string	true	"client id"
// @Success		200
// @Router		/api/quick/providers/{providerId} [delete]
func (o Controller) deleteProvider(c *gin.Context) {
	var tenant model.Tenant
	if err := o.SetCtx(c).SetTenant(&tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "init provider err")
		return
	}
	if err := o.svc.DeleteProvider(tenant.ID, c.Param("providerId")); err != nil {
		resp.ErrorSelect(c, err, "delete provider err")
		return
	}

	resp.Success(c)
}
