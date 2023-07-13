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
// @Param		poolId	path	string	true	"user pool id"
// @Success		200
// @Router		/api/quick/user-pools/{poolId} [get]
func (o Controller) getUserPool(c *gin.Context) {
	pool, err := o.svc.GetUserPool(c.Param("poolId"))
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
// @Param		bd		body	request.UserPoolReq	true	"body"
// @Success		200
// @Router		/api/quick/user-pools [post]
func (o Controller) createUserPool(c *gin.Context) {
	var in request.UserPoolReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init user pool req err")
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
// @Param		poolId	path	string				true	"user pool id"
// @Param		bd		body	request.UserPoolReq	true	"body"
// @Success		200
// @Router		/api/quick/user-pools/{poolId} [put]
func (o Controller) modifyUserPool(c *gin.Context) {
	var in request.UserPoolReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init user pool req err")
		return
	}

	if err := o.svc.ModifyUserPool(in.ToModel()); err != nil {
		resp.ErrorUnknown(c, err, "modify user pool err")
		return
	}
	resp.Success(c)
}

// @Summary	delete user pool
// @Schemes
// @Description	delete user pool
// @Tags		user
// @Param		poolId	path	string	true	"user pool id"
// @Success		200
// @Router		/api/quick/user-pools/{poolId} [delete]
func (o Controller) deleteUserPool(c *gin.Context) {
	if err := o.svc.DeleteUserPool(c.Param("poolId")); err != nil {
		resp.ErrorUnknown(c, err, "delete user pool err")
		return
	}
	resp.Success(c)
}
