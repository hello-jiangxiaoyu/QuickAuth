package controller

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"github.com/gin-gonic/gin"
)

// @Summary	user pools info
// @Schemes
// @Description	list user pool
// @Tags		user
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Success		200
// @Router		/api/quick/user-pools [get]
func (o Controller) listUserPool(c *gin.Context) {
	pools, err := o.svc.ListUserPool()
	if err != nil {
		resp.ErrorSelect(c, err, "list user pool err")
		return
	}
	resp.SuccessArray(c, len(pools), pools)
}

// @Summary	user pool info
// @Schemes
// @Description	list user pool
// @Tags		user
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		poolId		path	integer	true	"user pool id"
// @Success		200
// @Router		/api/quick/user-pools/{poolId} [get]
func (o Controller) getUserPool(c *gin.Context) {
	var in request.UserPoolReq
	if err := o.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid user-pools request param")
		return
	}

	pool, err := o.svc.GetUserPool(in.PoolId)
	if err != nil {
		resp.ErrorUnknown(c, err, "no such user pool")
		return
	}
	resp.SuccessWithData(c, pool)
}

// @Summary	create user pool
// @Schemes
// @Description	create user pool
// @Tags		user
// @Param		X-User-ID	header	string				false	"user id"
// @Param		X-Pool-ID	header	string				false	"user pool id"
// @Param		bd			body	request.UserPoolReq	true	"body"
// @Success		200
// @Router		/api/quick/user-pools [post]
func (o Controller) createUserPool(c *gin.Context) {
	var in request.UserPoolReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid user-pools request param")
		return
	}
	pool, err := o.svc.CreateUserPool(in.ToModel())
	if err != nil {
		resp.ErrorUnknown(c, err, "create user pool err")
		return
	}
	resp.SuccessWithData(c, pool)
}

// @Summary	modify user pool
// @Schemes
// @Description	modify user pool
// @Tags		user
// @Param		X-User-ID	header	string				false	"user id"
// @Param		X-Pool-ID	header	string				false	"user pool id"
// @Param		poolId		path	integer				true	"user pool id"
// @Param		bd			body	request.UserPoolReq	true	"body"
// @Success		200
// @Router		/api/quick/user-pools/{poolId} [put]
func (o Controller) modifyUserPool(c *gin.Context) {
	var in request.UserPoolReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid user-pools request param")
		return
	}

	if err := o.svc.ModifyUserPool(in.PoolId, in.ToModel()); err != nil {
		resp.ErrorUnknown(c, err, "modify user pool err")
		return
	}
	resp.Success(c)
}

// @Summary	delete user pool
// @Schemes
// @Description	delete user pool
// @Tags		user
// @Param		X-User-ID	header	string	false	"user id"
// @Param		X-Pool-ID	header	string	false	"user pool id"
// @Param		poolId		path	integer	true	"user pool id"
// @Success		200
// @Router		/api/quick/user-pools/{poolId} [delete]
func (o Controller) deleteUserPool(c *gin.Context) {
	var in request.UserPoolReq
	if err := o.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid user-pools request param")
		return
	}

	if err := o.svc.DeleteUserPool(in.PoolId); err != nil {
		resp.ErrorUnknown(c, err, "delete user pool err")
		return
	}
	resp.Success(c)
}
