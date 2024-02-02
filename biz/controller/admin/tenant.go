package admin

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/admin"
	"QuickAuth/pkg/utils"
	"github.com/gin-gonic/gin"
)

// ListTenant
// @Summary	get tenant list
// @Tags	tenant
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	appId		path	string	true	"app id"
// @Success	200
// @Router	/api/quick/apps/{appId}/tenants [get]
func ListTenant(c *gin.Context) {
	var in request.TenantReq
	if err := internal.BindUri(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	tenants, err := admin.ListTenant(in.AppID)
	if err != nil {
		resp.ErrorSelect(c, err, "list tenant err")
		return
	}

	resp.SuccessArrayData(c, len(tenants), utils.DtoFilter(tenants, model.TenantsDto))
}

// GetTenant
// @Summary	get tenant details
// @Tags	tenant
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	appId		path	string	true	"app id"
// @Param	tenantId	path	integer	true	"tenant id"
// @Success	200
// @Router	/api/quick/apps/{appId}/tenants/{tenantId} [get]
func GetTenant(c *gin.Context) {
	var in request.TenantReq
	if err := internal.BindUri(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	tenant, err := admin.GetTenant(in.AppID, in.TenantID)
	if err != nil {
		resp.ErrorSelect(c, err, "get tenant err")
		return
	}

	resp.SuccessWithData(c, tenant)
}

// CreateTenant
// @Summary	create tenant
// @Tags	tenant
// @Param	X-User-ID	header	string				false	"user id"
// @Param	X-Pool-ID	header	string				false	"user pool id"
// @Param	appId		path	string				true	"app id"
// @Param	bd			body	request.TenantReq	true	"body"
// @Success	200
// @Router	/api/quick/apps/{appId}/tenants [post]
func CreateTenant(c *gin.Context) {
	var in request.TenantReq
	if err := internal.BindUriAndJson(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	tenant, err := admin.CreatTenant(in.ToModel())
	if err != nil {
		resp.ErrorCreate(c, err, "create tenant err")
		return
	}

	resp.SuccessWithData(c, tenant)
}

// ModifyTenant
// @Summary	modify tenant
// @Tags	tenant
// @Param	X-User-ID	header	string				false	"user id"
// @Param	X-Pool-ID	header	string				false	"user pool id"
// @Param	appId		path	string				true	"app id"
// @Param	tenantId	path	integer				true	"tenant id"
// @Param	bd			body	request.TenantReq	true	"body"
// @Success	200
// @Router	/api/quick/apps/{appId}/tenants/{tenantId} [put]
func ModifyTenant(c *gin.Context) {
	var in request.TenantReq
	if err := internal.BindUriAndJson(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := admin.ModifyTenant(in.TenantID, in.ToModel()); err != nil {
		resp.ErrorUpdate(c, err, "modify tenant err")
		return
	}

	resp.Success(c)
}

// DeleteTenant
// @Summary	delete tenant
// @Tags	tenant
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	appId		path	string	true	"app id"
// @Param	tenantId	path	integer	true	"tenant id"
// @Success	200
// @Router	/api/quick/apps/{appId}/tenants/{tenantId} [delete]
func DeleteTenant(c *gin.Context) {
	var in request.TenantReq
	if err := internal.BindUri(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := admin.DeleteTenant(in.AppID, in.TenantID); err != nil {
		resp.ErrorDelete(c, err, "delete tenant err")
		return
	}

	resp.Success(c)
}
