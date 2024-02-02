package admin

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/admin"
	"github.com/gin-gonic/gin"
)

// ListUserPool
// @Summary	list user pool
// @Tags	user
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Success	200
// @Router	/api/quick/user-pools [get]
func ListUserPool(c *gin.Context) {
	pools, err := admin.ListUserPool()
	if err != nil {
		resp.ErrorSelect(c, err, "list user pool err")
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
func GetUserPool(c *gin.Context) {
	var in request.UserPoolReq
	if err := internal.BindUri(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	pool, err := admin.GetUserPool(in.PoolId)
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
func CreateUserPool(c *gin.Context) {
	var in request.UserPoolReq
	if err := internal.BindUriAndJson(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}
	pool, err := admin.CreateUserPool(in.ToModel())
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
func ModifyUserPool(c *gin.Context) {
	var in request.UserPoolReq
	if err := internal.BindUriAndJson(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := admin.ModifyUserPool(in.PoolId, in.ToModel()); err != nil {
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
func DeleteUserPool(c *gin.Context) {
	var in request.UserPoolReq
	if err := internal.BindUri(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := admin.DeleteUserPool(in.PoolId); err != nil {
		resp.ErrorDelete(c, err, "delete user pool err")
		return
	}
	resp.Success(c)
}
