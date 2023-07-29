package controller

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/pkg/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary	get tenant list
// @Schemes
// @Description	get tenant list
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

	tenants, err, _ := sg.Do("get-tenant-list-"+in.AppID, func() (interface{}, error) {
		return o.svc.ListTenant(in.AppID)
	})
	if err != nil {
		resp.ErrorSelect(c, err, "get tenant list err")
		return
	}

	resp.SuccessArray(c, len(tenants.([]model.Tenant)), tenants)
}

// @Summary	get tenant details
// @Schemes
// @Description	get tenant details
// @Tags		tenant
// @Param		appId		path	string	true	"app id"
// @Param		tenantId	path	integer	true	"tenant id"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants/{tenantId} [get]
func (o Controller) getTenant(c *gin.Context) {
	var in request.TenantReq
	if err := o.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid tenant request param")
		return
	}

	tenant, err, _ := sg.Do("get-tenant-"+strconv.FormatInt(in.TenantID, 10), func() (interface{}, error) {
		return o.svc.GetTenant(in.AppID, in.TenantID)
	})
	if err != nil {
		resp.ErrorSelect(c, err, "get tenant err")
		return
	}

	resp.SuccessWithData(c, tenant)
}

// @Summary	create tenant
// @Schemes
// @Description	create tenant
// @Tags		tenant
// @Param		appId	path	string				true	"app id"
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
		resp.ErrorSqlCreate(c, err, "create tenant err")
		return
	}

	resp.SuccessWithData(c, tenant)
}

// @Summary	modify tenant
// @Schemes
// @Description	modify tenant
// @Tags		tenant
// @Param		appId		path	string				true	"app id"
// @Param		tenantId	path	integer				true	"tenant id"
// @Param		bd			body	request.TenantReq	true	"body"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants/{tenantId} [put]
func (o Controller) modifyTenant(c *gin.Context) {
	var in request.TenantReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid tenant request param")
		return
	}

	if err := o.svc.ModifyTenant(in.TenantID, in.ToModel()); err != nil {
		resp.ErrorSelect(c, err, "modify tenant err")
		return
	}

	resp.Success(c)
}

// @Summary	delete tenant
// @Schemes
// @Description	delete tenant
// @Tags		tenant
// @Param		appId		path	string	true	"app id"
// @Param		tenantId	path	integer	true	"tenant id"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants/{tenantId} [delete]
func (o Controller) deleteTenant(c *gin.Context) {
	var in request.TenantReq
	if err := o.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid tenant request param")
		return
	}

	if err := o.svc.DeleteTenant(in.AppID, in.TenantID); err != nil {
		resp.ErrorSelect(c, err, "delete tenant err")
		return
	}

	resp.Success(c)
}
