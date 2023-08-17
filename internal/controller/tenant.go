package controller

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/internal/model"
	"QuickAuth/pkg/utils"
	"github.com/gin-gonic/gin"
)

// @Summary	get tenant list
// @Schemes
// @Description	get tenant list
// @Tags		tenant
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		appId		path	string	true	"app id"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants [get]
func (o Controller) listTenant(c *gin.Context) {
	var in request.TenantReq
	if err := o.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err, true)
		return
	}

	tenants, err := o.svc.ListTenant(in.AppID)
	if err != nil {
		resp.ErrorSelect(c, err, "list tenant err", true)
		return
	}

	resp.SuccessArrayData(c, len(tenants), utils.DtoFilter(tenants, model.TenantsDto))
}

// @Summary	get tenant details
// @Schemes
// @Description	get tenant details
// @Tags		tenant
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		appId		path	string	true	"app id"
// @Param		tenantId	path	integer	true	"tenant id"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants/{tenantId} [get]
func (o Controller) getTenant(c *gin.Context) {
	var in request.TenantReq
	if err := o.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	tenant, err := o.svc.GetTenant(in.AppID, in.TenantID)
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
// @Param		X-User-ID	header	string				false	"user id"
// @Param		X-Pool-ID	header	string				false	"user pool id"
// @Param		appId		path	string				true	"app id"
// @Param		bd			body	request.TenantReq	true	"body"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants [post]
func (o Controller) createTenant(c *gin.Context) {
	var in request.TenantReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	tenant, err := o.svc.CreatTenant(in.ToModel())
	if err != nil {
		resp.ErrorCreate(c, err, "create tenant err")
		return
	}

	resp.SuccessWithData(c, tenant)
}

// @Summary	modify tenant
// @Schemes
// @Description	modify tenant
// @Tags		tenant
// @Param		X-User-ID	header	string				false	"user id"
// @Param		X-Pool-ID	header	string				false	"user pool id"
// @Param		appId		path	string				true	"app id"
// @Param		tenantId	path	integer				true	"tenant id"
// @Param		bd			body	request.TenantReq	true	"body"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants/{tenantId} [put]
func (o Controller) modifyTenant(c *gin.Context) {
	var in request.TenantReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := o.svc.ModifyTenant(in.TenantID, in.ToModel()); err != nil {
		resp.ErrorUpdate(c, err, "modify tenant err")
		return
	}

	resp.Success(c)
}

// @Summary	delete tenant
// @Schemes
// @Description	delete tenant
// @Tags		tenant
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		appId		path	string	true	"app id"
// @Param		tenantId	path	integer	true	"tenant id"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants/{tenantId} [delete]
func (o Controller) deleteTenant(c *gin.Context) {
	var in request.TenantReq
	if err := o.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := o.svc.DeleteTenant(in.AppID, in.TenantID); err != nil {
		resp.ErrorDelete(c, err, "delete tenant err")
		return
	}

	resp.Success(c)
}
