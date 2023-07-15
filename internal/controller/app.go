package controller

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// @Summary	apps info
// @Schemes
// @Description	list apps
// @Tags		app
// @Success		200
// @Router		/api/quick/apps [get]
func (o Controller) listApp(c *gin.Context) {
	apps, err := o.svc.ListApps()
	if err != nil {
		resp.ErrorSelect(c, err, "list apps err")
		return
	}
	resp.SuccessArray(c, len(apps), apps)
}

// @Summary	apps info
// @Schemes
// @Description	list apps
// @Tags		app
// @Param		appId	path	string	true	"app id"
// @Success		200
// @Router		/api/quick/apps/{appId} [get]
func (o Controller) getApp(c *gin.Context) {
	var in request.AppReq
	if err := o.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid app request param")
		return
	}

	app, err := o.svc.GetApp(in.AppId)
	if err != nil {
		resp.ErrorUnknown(c, err, "no such app")
		return
	}
	resp.SuccessWithData(c, app)
}

// @Summary	create app
// @Schemes
// @Description	create app
// @Tags		app
// @Param		bd		body	request.AppReq	true	"body"
// @Success		200
// @Router		/api/quick/apps [post]
func (o Controller) createApp(c *gin.Context) {
	var in request.AppReq
	if err := o.SetCtx(c).BindJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid app request param")
		return
	}

	app, err := o.svc.CreateApp(in.ToModel())
	if err != nil {
		resp.ErrorUnknown(c, err, "create app err")
		return
	}
	resp.SuccessWithData(c, app)
}

// @Summary	modify app
// @Schemes
// @Description	modify app
// @Tags		app
// @Param		appId	path	string			true	"app id"
// @Param		bd		body	request.AppReq	true	"body"
// @Success		200
// @Router		/api/quick/apps/{appId} [put]
func (o Controller) modifyApp(c *gin.Context) {
	var in request.AppReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid app request param")
		return
	}

	if err := o.svc.ModifyApp(in.AppId, in.ToModel()); err != nil {
		resp.ErrorUnknown(c, err, "modify app err")
		return
	}
	resp.Success(c)
}

// @Summary	delete app
// @Schemes
// @Description	delete app
// @Tags		app
// @Param		appId	path	string	true	"app id"
// @Success		200
// @Router		/api/quick/apps/{appId} [delete]
func (o Controller) deleteApp(c *gin.Context) {
	if err := o.svc.DeleteApp(c.Param("appId")); err != nil {
		resp.ErrorUnknown(c, err, "delete app err")
		return
	}
	resp.Success(c)
}
