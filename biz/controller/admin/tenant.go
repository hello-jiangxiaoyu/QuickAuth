package admin

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/pkg/utils"
	"github.com/gin-gonic/gin"
)

// ListTenant	swagger
// @Description	get tenant list
// @Tags		tenant
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		appId		path	string	true	"app id"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants [get]
func (a Route) ListTenant(c *gin.Context) {
	var in request.TenantReq
	if err := a.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err, true)
		return
	}

	tenants, err := a.svc.ListTenant(in.AppID)
	if err != nil {
		resp.ErrorSelect(c, err, "list tenant err", true)
		return
	}

	resp.SuccessArrayData(c, len(tenants), utils.DtoFilter(tenants, model.TenantsDto))
}

// GetTenant	swagger
// @Description	get tenant details
// @Tags		tenant
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		appId		path	string	true	"app id"
// @Param		tenantId	path	integer	true	"tenant id"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants/{tenantId} [get]
func (a Route) GetTenant(c *gin.Context) {
	var in request.TenantReq
	if err := a.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	tenant, err := a.svc.GetTenant(in.AppID, in.TenantID)
	if err != nil {
		resp.ErrorSelect(c, err, "get tenant err")
		return
	}

	resp.SuccessWithData(c, tenant)
}

// CreateTenant	swagger
// @Description	create tenant
// @Tags		tenant
// @Param		X-User-ID	header	string				false	"user id"
// @Param		X-Pool-ID	header	string				false	"user pool id"
// @Param		appId		path	string				true	"app id"
// @Param		bd			body	request.TenantReq	true	"body"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants [post]
func (a Route) CreateTenant(c *gin.Context) {
	var in request.TenantReq
	if err := a.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	tenant, err := a.svc.CreatTenant(in.ToModel())
	if err != nil {
		resp.ErrorCreate(c, err, "create tenant err")
		return
	}

	resp.SuccessWithData(c, tenant)
}

// ModifyTenant	swagger
// @Description	modify tenant
// @Tags		tenant
// @Param		X-User-ID	header	string				false	"user id"
// @Param		X-Pool-ID	header	string				false	"user pool id"
// @Param		appId		path	string				true	"app id"
// @Param		tenantId	path	integer				true	"tenant id"
// @Param		bd			body	request.TenantReq	true	"body"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants/{tenantId} [put]
func (a Route) ModifyTenant(c *gin.Context) {
	var in request.TenantReq
	if err := a.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := a.svc.ModifyTenant(in.TenantID, in.ToModel()); err != nil {
		resp.ErrorUpdate(c, err, "modify tenant err")
		return
	}

	resp.Success(c)
}

// DeleteTenant	swagger
// @Description	delete tenant
// @Tags		tenant
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		appId		path	string	true	"app id"
// @Param		tenantId	path	integer	true	"tenant id"
// @Success		200
// @Router		/api/quick/apps/{appId}/tenants/{tenantId} [delete]
func (a Route) DeleteTenant(c *gin.Context) {
	var in request.TenantReq
	if err := a.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := a.svc.DeleteTenant(in.AppID, in.TenantID); err != nil {
		resp.ErrorDelete(c, err, "delete tenant err")
		return
	}

	resp.Success(c)
}