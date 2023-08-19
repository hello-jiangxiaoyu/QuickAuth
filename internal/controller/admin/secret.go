package admin

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// ListAppSecret	swagger
// @Summary	app secret info
// @Schemes
// @Description	list app secret
// @Tags		app
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		appId		path	string	true	"app id"
// @Success		200
// @Router		/api/quick/apps/{appId}/secrets [get]
func (a Route) ListAppSecret(c *gin.Context) {
	var in request.AppSecretReq
	if err := a.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err, true)
		return
	}

	secrets, err := a.svc.ListAppSecrets(in.AppId)
	if err != nil {
		resp.ErrorSelect(c, err, "list app secret err", true)
		return
	}
	resp.SuccessArrayData(c, len(secrets), secrets)
}

// CreateAppSecret	swagger
// @Summary	create app secret
// @Schemes
// @Description	create app secret
// @Tags		app
// @Param		X-User-ID	header	string					false	"user id"
// @Param		X-Pool-ID	header	string					false	"user pool id"
// @Param		appId		path	string					true	"app id"
// @Param		bd			body	request.AppSecretReq	true	"body"
// @Success		200
// @Router		/api/quick/apps/{appId}/secrets [post]
func (a Route) CreateAppSecret(c *gin.Context) {
	var in request.AppSecretReq
	if err := a.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	secret, err := a.svc.CreateAppSecret(in.AppId, in.ToModel())
	if err != nil {
		resp.ErrorCreate(c, err, "create app secret err")
		return
	}
	resp.SuccessWithData(c, secret)
}

// ModifyAppSecret	swagger
// @Summary	create app secret
// @Schemes
// @Description	create app secret
// @Tags		app
// @Param		X-User-ID	header	string					false	"user id"
// @Param		X-Pool-ID	header	string					false	"user pool id"
// @Param		appId		path	string					true	"app id"
// @Param		secretId	path	integer					true	"secret id"
// @Param		bd			body	request.AppSecretReq	true	"body"
// @Success		200
// @Router		/api/quick/apps/{appId}/secrets/{secretId} [put]
func (a Route) ModifyAppSecret(c *gin.Context) {
	var in request.AppSecretReq
	if err := a.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	secret, err := a.svc.ModifyAppSecret(in.SecretId, in.ToModel())
	if err != nil {
		resp.ErrorUpdate(c, err, "modify app secret err")
		return
	}
	resp.SuccessWithData(c, secret)
}

// DeleteAppSecret	swagger
// @Summary	delete app secret
// @Schemes
// @Description	delete app secret
// @Tags		app
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		appId		path	string	true	"app id"
// @Param		secretId	path	integer	true	"secret id"
// @Success		200
// @Router		/api/quick/apps/{appId}/secrets/{secretId} [delete]
func (a Route) DeleteAppSecret(c *gin.Context) {
	var in request.AppSecretReq
	if err := a.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := a.svc.DeleteAppSecret(in.AppId, in.SecretId); err != nil {
		resp.ErrorDelete(c, err, "delete app secret err")
		return
	}
	resp.Success(c)
}
