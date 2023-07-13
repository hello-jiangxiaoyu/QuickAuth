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
	secrets, err := o.svc.ListAppSecrets(c.Param("appId"))
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
		resp.ErrorRequest(c, err, "init app secret req err")
		return
	}

	secret, err := o.svc.CreateAppSecret(in.ToModel())
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
// @Param		appId	path	string	true	"app id"
// @Success		200
// @Router		/api/quick/apps/{appId}/secrets/{secretId} [delete]
func (o Controller) deleteAppSecret(c *gin.Context) {
	appId := c.Param("appId")
	secretId := c.Param("secretId")
	if err := o.svc.DeleteAppSecret(appId, secretId); err != nil {
		resp.ErrorUnknown(c, err, "delete app secret err")
		return
	}
	resp.Success(c)
}
