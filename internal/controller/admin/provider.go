package admin

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// ListProvider	swagger
// @Description	list provider info
// @Tags		provider
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		vhost		header	string	false	"tenant host"
// @Success		200
// @Router		/api/quick/providers [get]
func (a Route) ListProvider(c *gin.Context) {
	var in request.ProviderReq
	if err := a.SetCtx(c).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	providers, err := a.svc.GetLoginProviderInfo(in.Tenant.ID)
	if err != nil {
		resp.ErrorSelect(c, err, "list provider err")
		return
	}

	resp.SuccessArrayData(c, len(providers), providers)
}

// GetProvider	swagger
// @Description	get provider details
// @Tags		provider
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		vhost		header	string	false	"tenant host"
// @Param		providerId	path	integer	true	"provider id"
// @Success		200
// @Router		/api/quick/providers/{providerId} [get]
func (a Route) GetProvider(c *gin.Context) {
	var in request.ProviderReq
	if err := a.SetCtx(c).BindUri(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	provider, err := a.svc.GetProvider(in.Tenant.ID, in.ProviderId)
	if err != nil {
		resp.ErrorSelect(c, err, "get provider err")
		return
	}

	resp.SuccessWithData(c, provider)
}

// CreateProvider	swagger
// @Description	create provider
// @Tags		provider
// @Param		X-User-ID	header	string				false	"user id"
// @Param		X-Pool-ID	header	string				false	"user pool id"
// @Param		vhost		header	string				false	"tenant host"
// @Param		bd			body	request.ProviderReq	true	"body"
// @Success		200
// @Router		/api/quick/providers [post]
func (a Route) CreateProvider(c *gin.Context) {
	var in request.ProviderReq
	if err := a.SetCtx(c).BindJson(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	provider, err := a.svc.CreateProvider(in.ToModel())
	if err != nil {
		resp.ErrorCreate(c, err, "create provider err")
		return
	}

	resp.SuccessWithData(c, provider)
}

// ModifyProvider	swagger
// @Description	modify provider
// @Tags		provider
// @Param		providerId	path	integer				true	"provider id"
// @Param		vhost		header	string				false	"tenant host"
// @Param		bd			body	request.ProviderReq	true	"body"
// @Success		200
// @Router		/api/quick/providers/{providerId} [put]
func (a Route) ModifyProvider(c *gin.Context) {
	var in request.ProviderReq
	if err := a.SetCtx(c).BindUriAndJson(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := a.svc.ModifyProvider(in.ProviderId, in.ToModel()); err != nil {
		resp.ErrorUpdate(c, err, "modify provider err")
		return
	}

	resp.Success(c)
}

// DeleteProvider	swagger
// @Description	delete provider
// @Tags		provider
// @Param		providerId	path	integer	true	"provider id"
// @Param		vhost		header	string	false	"tenant host"
// @Success		200
// @Router		/api/quick/providers/{providerId} [delete]
func (a Route) DeleteProvider(c *gin.Context) {
	var in request.ProviderReq
	if err := a.SetCtx(c).BindUri(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := a.svc.DeleteProvider(in.Tenant.ID, in.ProviderId); err != nil {
		resp.ErrorDelete(c, err, "delete provider err")
		return
	}

	resp.Success(c)
}
