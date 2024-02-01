package admin

import (
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// ListUserPool
// @Summary	list user pool
// @Tags	user
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Success	200
// @Router	/api/quick/user-pools [get]
func (a Route) ListUserPool(c *gin.Context) {
	pools, err := a.svc.ListUserPool()
	if err != nil {
		resp.ErrorSelect(c, err, "list user pool err", true)
		return
	}
	resp.SuccessArrayData(c, len(pools), pools)
}

// GetUserPool
// @Summary	get user pool info
// @Tags	user
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	poolId		path	integer	true	"user pool id"
// @Success	200
// @Router	/api/quick/user-pools/{poolId} [get]
func (a Route) GetUserPool(c *gin.Context) {
	var in request.UserPoolReq
	if err := a.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	pool, err := a.svc.GetUserPool(in.PoolId)
	if err != nil {
		resp.ErrorSelect(c, err, "get user pool err")
		return
	}
	resp.SuccessWithData(c, pool)
}

// CreateUserPool
// @Summary	create user pool
// @Tags	user
// @Param	X-User-ID	header	string				false	"user id"
// @Param	X-Pool-ID	header	string				false	"user pool id"
// @Param	bd			body	request.UserPoolReq	true	"body"
// @Success	200
// @Router	/api/quick/user-pools [post]
func (a Route) CreateUserPool(c *gin.Context) {
	var in request.UserPoolReq
	if err := a.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	pool, err := a.svc.CreateUserPool(in.ToModel())
	if err != nil {
		resp.ErrorCreate(c, err, "create user pool err")
		return
	}
	resp.SuccessWithData(c, pool)
}

// ModifyUserPool
// @Summary	modify user pool
// @Tags	user
// @Param	X-User-ID	header	string				false	"user id"
// @Param	X-Pool-ID	header	string				false	"user pool id"
// @Param	poolId		path	integer				true	"user pool id"
// @Param	bd			body	request.UserPoolReq	true	"body"
// @Success	200
// @Router	/api/quick/user-pools/{poolId} [put]
func (a Route) ModifyUserPool(c *gin.Context) {
	var in request.UserPoolReq
	if err := a.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := a.svc.ModifyUserPool(in.PoolId, in.ToModel()); err != nil {
		resp.ErrorUpdate(c, err, "modify user pool err")
		return
	}
	resp.Success(c)
}

// DeleteUserPool
// @Summary	delete user pool
// @Tags	user
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	poolId		path	integer	true	"user pool id"
// @Success	200
// @Router	/api/quick/user-pools/{poolId} [delete]
func (a Route) DeleteUserPool(c *gin.Context) {
	var in request.UserPoolReq
	if err := a.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := a.svc.DeleteUserPool(in.PoolId); err != nil {
		resp.ErrorDelete(c, err, "delete user pool err")
		return
	}
	resp.Success(c)
}
