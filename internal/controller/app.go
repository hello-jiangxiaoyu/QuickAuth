package controller

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/internal/model"
	"QuickAuth/internal/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

var sg singleflight.Group

// @Summary	apps info
// @Schemes
// @Description	list apps
// @Tags		app
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Success		200
// @Router		/api/quick/apps [get]
func (o Controller) listApp(c *gin.Context) {
	apps, err, _ := sg.Do("get-app-list", func() (interface{}, error) {
		return o.svc.ListApps()
	})
	if err != nil {
		resp.ErrorSelect(c, err, "list apps err", true)
		return
	}
	resp.SuccessArray(c, len(apps.([]model.App)), apps)
}

// @Summary	apps info
// @Schemes
// @Description	list apps
// @Tags		app
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		appId		path	string	true	"app id"
// @Success		200
// @Router		/api/quick/apps/{appId} [get]
func (o Controller) getApp(c *gin.Context) {
	var in request.AppReq
	if err := o.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	app, err := o.svc.GetAppDetail(in.AppId)
	if err != nil {
		resp.ErrorSelect(c, err, "get app err")
		return
	}
	resp.SuccessWithData(c, app.Dto())
}

// @Summary	create app
// @Schemes
// @Description	create app
// @Tags		app
// @Param		X-User-ID	header	string			false	"user id"
// @Param		X-Pool-ID	header	string			false	"user pool id"
// @Param		bd			body	request.AppReq	true	"body"
// @Success		200
// @Router		/api/quick/apps [post]
func (o Controller) createApp(c *gin.Context) {
	var in request.AppReq
	if err := o.SetCtx(c).BindJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	app, err := o.svc.CreateApp(in.ToModel(), in.Host, in.PoolId)
	if err != nil {
		resp.ErrorCreate(c, err, "create app err")
		return
	}
	resp.SuccessWithData(c, app)
}

// @Summary	modify app
// @Schemes
// @Description	modify app
// @Tags		app
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		appId		path	string			true	"app id"
// @Param		bd			body	request.AppReq	true	"body"
// @Success		200
// @Router		/api/quick/apps/{appId} [put]
func (o Controller) modifyApp(c *gin.Context) {
	var in request.AppReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := o.svc.ModifyApp(in.AppId, in.ToModel()); err != nil {
		resp.ErrorUpdate(c, err, "modify app err")
		return
	}
	resp.Success(c)
}

// @Summary	delete app
// @Schemes
// @Description	delete app
// @Tags		app
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		appId		path	string	true	"app id"
// @Success		200
// @Router		/api/quick/apps/{appId} [delete]
func (o Controller) deleteApp(c *gin.Context) {
	if err := o.svc.DeleteApp(c.Param("appId")); err != nil {
		if err == service.ErrorDeleteDefaultApp {
			resp.ErrorUnknown(c, err, err.Error())
		} else {
			resp.ErrorDelete(c, err, "delete app err")
		}
		return
	}
	resp.Success(c)
}
