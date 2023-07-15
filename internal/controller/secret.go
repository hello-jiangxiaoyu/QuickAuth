package controller

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// @Summary	app secret info
// @Schemes
// @Description	list app secret
// @Tags		app
// @Param		appId	path	string	true	"app id"
// @Success		200
// @Router		/api/quick/apps/{appId}/secrets [get]
func (o Controller) listAppSecret(c *gin.Context) {
	var in request.AppSecretReq
	if err := o.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid app secret request param")
		return
	}

	secrets, err := o.svc.ListAppSecrets(in.AppId)
	if err != nil {
		resp.ErrorSelect(c, err, "list app secret err")
		return
	}
	resp.SuccessArray(c, len(secrets), secrets)
}

// @Summary	create app secret
// @Schemes
// @Description	create app secret
// @Tags		app
// @Param		appId	path	string					true	"app id"
// @Param		bd		body	request.AppSecretReq	true	"body"
// @Success		200
// @Router		/api/quick/apps/{appId}/secrets [post]
func (o Controller) createAppSecret(c *gin.Context) {
	var in request.AppSecretReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid app secret request param")
		return
	}

	secret, err := o.svc.CreateAppSecret(in.ToModel())
	if err != nil {
		resp.ErrorUnknown(c, err, "create app secret err")
		return
	}
	resp.SuccessWithData(c, secret)
}

// @Summary	create app secret
// @Schemes
// @Description	create app secret
// @Tags		app
// @Param		appId		path	string					true	"app id"
// @Param		secretId	path	string					true	"secret id"
// @Param		bd			body	request.AppSecretReq	true	"body"
// @Success		200
// @Router		/api/quick/apps/{appId}/secrets/{secretId} [put]
func (o Controller) modifyAppSecret(c *gin.Context) {
	var in request.AppSecretReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid app secret request param")
		return
	}

	secret, err := o.svc.ModifyAppSecret(in.ToModel())
	if err != nil {
		resp.ErrorUnknown(c, err, "create app secret err")
		return
	}
	resp.SuccessWithData(c, secret)
}

// @Summary	delete app secret
// @Schemes
// @Description	delete app secret
// @Tags		app
// @Param		appId		path	string	true	"app id"
// @Param		secretId	path	string	true	"secret id"
// @Success		200
// @Router		/api/quick/apps/{appId}/secrets/{secretId} [delete]
func (o Controller) deleteAppSecret(c *gin.Context) {
	var in request.AppSecretReq
	if err := o.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid app secret request param")
		return
	}

	if err := o.svc.DeleteAppSecret(in.AppId, in.SecretId); err != nil {
		resp.ErrorUnknown(c, err, "delete app secret err")
		return
	}
	resp.Success(c)
}
