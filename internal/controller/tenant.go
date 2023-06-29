package controller

import (
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/pkg/model"
	"github.com/gin-gonic/gin"
)

// @Summary	provider info
// @Schemes
// @Description	list provider info
// @Tags		tenant
// @Param		clientId	path	string	true	"client id"
// @Success		200
// @Router		/api/quick/clients/{clientId}/tenants [get]
func (o Controller) listTenant(c *gin.Context) {
	tenants, err := o.svc.ListTenant(c.Param("clientId"))
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
// @Param		clientId	path	string	true	"client id"
// @Param		tenantId	path	string	true	"tenant id"
// @Success		200
// @Router		/api/quick/clients/{clientId}/tenants/{tenantId} [get]
func (o Controller) getTenant(c *gin.Context) {
	clientId := c.Param("clientId")
	tenantId := c.Param("tenantId")
	tenant, err := o.svc.GetTenant(clientId, tenantId)
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
// @Param		clientId	path	string	true	"client id"
// @Param		bd			body	model.Tenant	true	"body"
// @Success		200
// @Router		/api/quick/clients/{clientId}/tenants [post]
func (o Controller) createTenant(c *gin.Context) {
	var in model.Tenant
	if err := o.SetCtx(c).BindJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init tenant err")
		return
	}
	in.ClientID = c.Param("clientId")
	tenant, err := o.svc.CreatTenant(in)
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
// @Param		clientId	path	string	true	"client id"
// @Param		tenantId	path	string	true	"tenant id"
// @Param		bd			body	model.Tenant	true	"body"
// @Success		200
// @Router		/api/quick/clients/{clientId}/tenants/{tenantId} [put]
func (o Controller) modifyTenant(c *gin.Context) {
	var in model.Tenant
	if err := o.SetCtx(c).BindJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init tenant err")
		return
	}
	in.ID = c.Param("tenantId")
	in.ClientID = c.Param("clientId")
	if err := o.svc.ModifyTenant(in); err != nil {
		resp.ErrorSelect(c, err, "modify tenant err")
		return
	}

	resp.Success(c)
}

// @Summary	get provider details
// @Schemes
// @Description	get provider details
// @Tags		tenant
// @Param		clientId	path	string	true	"client id"
// @Param		tenantId	path	string	true	"tenant id"
// @Success		200
// @Router		/api/quick/clients/{clientId}/tenants/{tenantId} [delete]
func (o Controller) deleteTenant(c *gin.Context) {
	clientId := c.Param("clientId")
	tenantId := c.Param("tenantId")
	if err := o.svc.DeleteTenant(clientId, tenantId); err != nil {
		resp.ErrorSelect(c, err, "delete tenant err")
		return
	}

	resp.Success(c)
}
