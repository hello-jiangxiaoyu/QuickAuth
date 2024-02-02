package admin

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/admin"
	"errors"
	"github.com/gin-gonic/gin"
)

type Route struct {
}

func NewAdminRoute() *Route {
	return &Route{}
}

// ListApp
// @Summary	list apps
// @Tags	app
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Success	200
// @Router	/api/quick/apps [get]
func (a Route) ListApp(c *gin.Context) {
	apps, err := admin.ListApps()
	if err != nil {
		resp.ErrorSelect(c, err, "list apps err")
		return
	}
	resp.SuccessArrayData(c, len(apps), apps)
}

// GetApp
// @Summary	get app
// @Tags	app
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	appId		path	string	true	"app id"
// @Success	200
// @Router	/api/quick/apps/{appId} [get]
func (a Route) GetApp(c *gin.Context) {
	var in request.AppReq
	if err := internal.BindUri(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	app, err := admin.GetAppDetail(in.AppId)
	if err != nil {
		resp.ErrorSelect(c, err, "get app err")
		return
	}
	resp.SuccessWithData(c, app.Dto())
}

// CreateApp
// @Summary	create app
// @Tags	app
// @Param	X-User-ID	header	string			false	"user id"
// @Param	X-Pool-ID	header	string			false	"user pool id"
// @Param	bd			body	request.AppReq	true	"body"
// @Success	200
// @Router	/api/quick/apps [post]
func (a Route) CreateApp(c *gin.Context) {
	var in request.AppReq
	if err := internal.BindJson(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	app, err := admin.CreateApp(in.ToModel(), in.Host, in.PoolId)
	if err != nil {
		resp.ErrorCreate(c, err, "create app err")
		return
	}
	resp.SuccessWithData(c, app)
}

// ModifyApp
// @Summary	modify app
// @Tags	app
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	appId		path	string			true	"app id"
// @Param	bd			body	request.AppReq	true	"body"
// @Success	200
// @Router	/api/quick/apps/{appId} [put]
func (a Route) ModifyApp(c *gin.Context) {
	var in request.AppReq
	if err := internal.BindUriAndJson(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := admin.ModifyApp(in.AppId, in.ToModel()); err != nil {
		resp.ErrorUpdate(c, err, "modify app err")
		return
	}
	resp.Success(c)
}

// DeleteApp
// @Summary	delete app
// @Tags	app
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	appId		path	string	true	"app id"
// @Success	200
// @Router	/api/quick/apps/{appId} [delete]
func (a Route) DeleteApp(c *gin.Context) {
	if err := admin.DeleteApp(c.Param("appId")); err != nil {
		if errors.Is(err, admin.ErrorDeleteDefaultApp) {
			resp.ErrorUnknown(c, err, err.Error())
		} else {
			resp.ErrorDelete(c, err, "delete app err")
		}
		return
	}
	resp.Success(c)
}
