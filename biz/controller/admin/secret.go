package admin

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/admin"
	"github.com/gin-gonic/gin"
)

// ListAppSecret
// @Summary	list app secret
// @Tags	app
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	appId		path	string	true	"app id"
// @Success	200
// @Router	/api/quick/apps/{appId}/secrets [get]
func ListAppSecret(c *gin.Context) {
	var in request.AppSecretReq
	if err := internal.BindUri(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	secrets, err := admin.ListAppSecrets(in.AppId)
	if err != nil {
		resp.ErrorSelect(c, err, "list app secret err")
		return
	}
	resp.SuccessArrayData(c, len(secrets), secrets)
}

// CreateAppSecret
// @Summary	create app secret
// @Tags	app
// @Param	X-User-ID	header	string					false	"user id"
// @Param	X-Pool-ID	header	string					false	"user pool id"
// @Param	vhost		query	string					false	"virtual host"
// @Param	appId		path	string					true	"app id"
// @Param	bd			body	request.AppSecretReq	true	"body"
// @Success	200
// @Router	/api/quick/apps/{appId}/secrets [post]
func CreateAppSecret(c *gin.Context) {
	var in request.AppSecretReq
	if err := internal.BindUriAndJson(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	secret, err := admin.CreateAppSecret(in.AppId, in.ToModel())
	if err != nil {
		resp.ErrorCreate(c, err, "create app secret err")
		return
	}
	resp.SuccessWithData(c, secret)
}

// ModifyAppSecret
// @Summary	create app secret
// @Tags	app
// @Param	X-User-ID	header	string					false	"user id"
// @Param	X-Pool-ID	header	string					false	"user pool id"
// @Param	vhost		query	string					false	"virtual host"
// @Param	appId		path	string					true	"app id"
// @Param	secretId	path	integer					true	"secret id"
// @Param	bd			body	request.AppSecretReq	true	"body"
// @Success	200
// @Router	/api/quick/apps/{appId}/secrets/{secretId} [put]
func ModifyAppSecret(c *gin.Context) {
	var in request.AppSecretReq
	if err := internal.BindUriAndJson(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	secret, err := admin.ModifyAppSecret(in.SecretId, in.ToModel())
	if err != nil {
		resp.ErrorUpdate(c, err, "modify app secret err")
		return
	}
	resp.SuccessWithData(c, secret)
}

// DeleteAppSecret
// @Summary	delete app secret
// @Tags	app
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	appId		path	string	true	"app id"
// @Param	secretId	path	integer	true	"secret id"
// @Success	200
// @Router	/api/quick/apps/{appId}/secrets/{secretId} [delete]
func DeleteAppSecret(c *gin.Context) {
	var in request.AppSecretReq
	if err := internal.BindUri(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := admin.DeleteAppSecret(in.AppId, in.SecretId); err != nil {
		resp.ErrorDelete(c, err, "delete app secret err")
		return
	}
	resp.Success(c)
}
