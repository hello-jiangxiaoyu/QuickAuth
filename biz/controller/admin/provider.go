package admin

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/admin"
	"github.com/gin-gonic/gin"
)

// ListLoginProviderInfo
// @Summary	list provider info
// @Tags	provider
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		header	string	false	"tenant host"
// @Success	200
// @Router	/api/quick/providers [get]
func ListLoginProviderInfo(c *gin.Context) {
	var in request.ProviderReq
	if err := internal.New(c).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	providers, err := admin.GetLoginProviderInfo(in.Tenant.ID)
	if err != nil {
		resp.ErrorSelect(c, err, "list provider err")
		return
	}

	resp.SuccessArrayData(c, len(providers), providers)
}

// GetProvider
// @Summary	get provider details
// @Tags		provider
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		vhost		header	string	false	"tenant host"
// @Param		providerId	path	integer	true	"provider id"
// @Success		200
// @Router		/api/quick/providers/{providerId} [get]
func GetProvider(c *gin.Context) {
	var in request.ProviderReq
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	provider, err := admin.GetProviderById(in.Tenant.ID, in.ProviderId)
	if err != nil {
		resp.ErrorSelect(c, err, "get provider err")
		return
	}

	resp.SuccessWithData(c, provider)
}

// CreateProvider
// @Summary	create provider
// @Tags	provider
// @Param	X-User-ID	header	string				false	"user id"
// @Param	X-Pool-ID	header	string				false	"user pool id"
// @Param	vhost		header	string				false	"tenant host"
// @Param	bd			body	request.ProviderReq	true	"body"
// @Success	200
// @Router	/api/quick/providers [post]
func CreateProvider(c *gin.Context) {
	var in request.ProviderReq
	if err := internal.BindJson(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	provider, err := admin.CreateProvider(in.ToModel())
	if err != nil {
		resp.ErrorCreate(c, err, "create provider err")
		return
	}

	resp.SuccessWithData(c, provider)
}

// ModifyProvider
// @Summary	modify provider
// @Tags	provider
// @Param	providerId	path	integer				true	"provider id"
// @Param	vhost		header	string				false	"tenant host"
// @Param	bd			body	request.ProviderReq	true	"body"
// @Success	200
// @Router	/api/quick/providers/{providerId} [put]
func ModifyProvider(c *gin.Context) {
	var in request.ProviderReq
	if err := internal.BindUriAndJson(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := admin.ModifyProvider(in.ProviderId, in.ToModel()); err != nil {
		resp.ErrorUpdate(c, err, "modify provider err")
		return
	}

	resp.Success(c)
}

// DeleteProvider
// @Summary	delete provider
// @Tags	provider
// @Param	providerId	path	integer	true	"provider id"
// @Param	vhost		header	string	false	"tenant host"
// @Success	200
// @Router	/api/quick/providers/{providerId} [delete]
func DeleteProvider(c *gin.Context) {
	var in request.ProviderReq
	if err := internal.BindUri(c, &in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	if err := admin.DeleteProvider(in.Tenant.ID, in.ProviderId); err != nil {
		resp.ErrorDelete(c, err, "delete provider err")
		return
	}

	resp.Success(c)
}
