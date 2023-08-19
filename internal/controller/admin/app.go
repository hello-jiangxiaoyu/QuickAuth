package admin

import (
	"QuickAuth/internal/controller/internal"
	"QuickAuth/internal/endpoint/model"
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/internal/service"
	"QuickAuth/internal/service/admin"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

var sg singleflight.Group

type Route struct {
	svc *service.Service
	internal.Api
}

func NewAdminRoute(svc *service.Service) *Route {
	return &Route{svc: svc}
}

// ListApp	swagger
// @Summary	apps info
// @Schemes
// @Description	list apps
// @Tags		app
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Success		200
// @Router		/api/quick/apps [get]
func (a Route) ListApp(c *gin.Context) {
	apps, err, _ := sg.Do("get-app-list", func() (interface{}, error) {
		return a.svc.ListApps()
	})
	if err != nil {
		resp.ErrorSelect(c, err, "list apps err", true)
		return
	}
	resp.SuccessArrayData(c, len(apps.([]model.App)), apps)
}

// GetApp	swagger
// @Summary	apps info
// @Schemes
// @Description	list apps
// @Tags		app
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		appId		path	string	true	"app id"
// @Success		200
// @Router		/api/quick/apps/{appId} [get]
func (a Route) GetApp(c *gin.Context) {
	var in request.AppReq
	if err := a.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	app, err := a.svc.GetAppDetail(in.AppId)
	if err != nil {
		resp.ErrorSelect(c, err, "get app err")
		return
	}
	resp.SuccessWithData(c, app.Dto())
}

// CreateApp	swagger
// @Summary	create app
// @Schemes
// @Description	create app
// @Tags		app
// @Param		X-User-ID	header	string			false	"user id"
// @Param		X-Pool-ID	header	string			false	"user pool id"
// @Param		bd			body	request.AppReq	true	"body"
// @Success		200
// @Router		/api/quick/apps [post]
func (a Route) CreateApp(c *gin.Context) {
	var in request.AppReq
	if err := a.SetCtx(c).BindJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	app, err := a.svc.CreateApp(in.ToModel(), in.Host, in.PoolId)
	if err != nil {
		resp.ErrorCreate(c, err, "create app err")
		return
	}
	resp.SuccessWithData(c, app)
}

// ModifyApp	swagger
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
func (a Route) ModifyApp(c *gin.Context) {
	var in request.AppReq
	if err := a.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := a.svc.ModifyApp(in.AppId, in.ToModel()); err != nil {
		resp.ErrorUpdate(c, err, "modify app err")
		return
	}
	resp.Success(c)
}

// DeleteApp	swagger
// @Summary	delete app
// @Schemes
// @Description	delete app
// @Tags		app
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		appId		path	string	true	"app id"
// @Success		200
// @Router		/api/quick/apps/{appId} [delete]
func (a Route) DeleteApp(c *gin.Context) {
	if err := a.svc.DeleteApp(c.Param("appId")); err != nil {
		if err == admin.ErrorDeleteDefaultApp {
			resp.ErrorUnknown(c, err, err.Error())
		} else {
			resp.ErrorDelete(c, err, "delete app err")
		}
		return
	}
	resp.Success(c)
}
