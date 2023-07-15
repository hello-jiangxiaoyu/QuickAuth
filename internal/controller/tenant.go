package controller

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// @Summary	provider info
// @Schemes
// @Description	list provider info
// @Tags		tenant
// @Param		appId	path	string	true	"app id"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants [get]
func (o Controller) listTenant(c *gin.Context) {
	var in request.TenantReq
	if err := o.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid tenant request param")
		return
	}

	tenants, err := o.svc.ListTenant(in.AppID)
	if err != nil {
		resp.ErrorSelect(c, err, "get tenant list err")
		return
	}

	resp.SuccessArray(c, len(tenants), tenants)
}

// @Summary	get provider details
// @Schemes
// @Description	get provider details
// @Tags		tenant
// @Param		appId	path	string	true	"app id"
// @Param		vhost	header	string	false	"tenant host"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants/current [get]
func (o Controller) getTenant(c *gin.Context) {
	var in request.TenantReq
	if err := o.SetCtx(c).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid tenant request param")
		return
	}

	resp.SuccessWithData(c, in.Tenant)
}

// @Summary	get provider details
// @Schemes
// @Description	get provider details
// @Tags		tenant
// @Param		appId	path	string				true	"app id"
// @Param		vhost	header	string				false	"tenant host"
// @Param		bd		body	request.TenantReq	true	"body"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants [post]
func (o Controller) createTenant(c *gin.Context) {
	var in request.TenantReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid tenant request param")
		return
	}

	tenant, err := o.svc.CreatTenant(in.ToModel())
	if err != nil {
		resp.ErrorSelect(c, err, "create tenant err")
		return
	}

	resp.SuccessWithData(c, tenant)
}

// @Summary	get provider details
// @Schemes
// @Description	get provider details
// @Tags		tenant
// @Param		appId	path	string				true	"app id"
// @Param		vhost	header	string				false	"tenant host"
// @Param		bd		body	request.TenantReq	true	"body"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants [put]
func (o Controller) modifyTenant(c *gin.Context) {
	var in request.TenantReq
	if err := o.SetCtx(c).BindUriAndJson(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid tenant request param")
		return
	}

	if err := o.svc.ModifyTenant(in.ToModel()); err != nil {
		resp.ErrorSelect(c, err, "modify tenant err")
		return
	}

	resp.Success(c)
}

// @Summary	get provider details
// @Schemes
// @Description	get provider details
// @Tags		tenant
// @Param		appId		path	string	true	"app id"
// @Param		vhost		header	string	false	"tenant host"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants [delete]
func (o Controller) deleteTenant(c *gin.Context) {
	var in request.TenantReq
	if err := o.SetCtx(c).BindUri(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid tenant request param")
		return
	}

	if err := o.svc.DeleteTenant(in.AppID, in.Tenant.ID); err != nil {
		resp.ErrorSelect(c, err, "delete tenant err")
		return
	}

	resp.Success(c)
}
