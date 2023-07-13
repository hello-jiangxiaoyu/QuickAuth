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
	tenants, err := o.svc.ListTenant(c.Param("appId"))
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
// @Param		appId		path	string	true	"app id"
// @Param		tenantId	path	string	true	"tenant id"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants/{tenantId} [get]
func (o Controller) getTenant(c *gin.Context) {
	appId := c.Param("appId")
	tenantId := c.Param("tenantId")
	tenant, err := o.svc.GetTenant(appId, tenantId)
	if err != nil {
		resp.ErrorSelect(c, err, "get tenant err")
		return
	}

	resp.SuccessWithData(c, tenant)
}

// @Summary	get provider details
// @Schemes
// @Description	get provider details
// @Tags		tenant
// @Param		appId	path	string				true	"app id"
// @Param		bd		body	request.TenantReq	true	"body"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants [post]
func (o Controller) createTenant(c *gin.Context) {
	var in request.TenantReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init tenant err")
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
// @Param		appId		path	string				true	"app id"
// @Param		tenantId	path	string				true	"tenant id"
// @Param		bd			body	request.TenantReq	true	"body"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants/{tenantId} [put]
func (o Controller) modifyTenant(c *gin.Context) {
	var in request.TenantReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init tenant err")
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
// @Param		tenantId	path	string	true	"tenant id"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants/{tenantId} [delete]
func (o Controller) deleteTenant(c *gin.Context) {
	appId := c.Param("appId")
	tenantId := c.Param("tenantId")
	if err := o.svc.DeleteTenant(appId, tenantId); err != nil {
		resp.ErrorSelect(c, err, "delete tenant err")
		return
	}

	resp.Success(c)
}
