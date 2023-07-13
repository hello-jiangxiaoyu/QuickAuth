package controller

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/pkg/tools/safe"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary	user info
// @Schemes
// @Description	list users
// @Tags		user
// @Param		poolId	path	string	true	"user pool id"
// @Success		200
// @Router		/api/quick/user-pools/{poolId}/users [get]
func (o Controller) listUser(c *gin.Context) {
	users, err := o.svc.ListUser(c.Param("poolId"))
	if err != nil {
		resp.ErrorSelect(c, err, "list user pool err")
		return
	}
	resp.SuccessArray(c, len(users), users)
}

// @Summary	user info
// @Schemes
// @Description	list user
// @Tags		user
// @Param		poolId	path	string	true	"user pool id"
// @Param		userId	path	string	true	"user id"
// @Success		200
// @Router		/api/quick/user-pools/{poolId}/users/{userId} [get]
func (o Controller) getUser(c *gin.Context) {
	poolId := c.Param("poolId")
	userId := c.Param("userId")
	user, err := o.svc.GetUserById(poolId, userId)
	if err != nil {
		resp.ErrorUnknown(c, err, "no such user")
		return
	}
	resp.SuccessWithData(c, user)
}

// @Summary	create user
// @Schemes
// @Description	create user
// @Tags		user
// @Param		poolId	path	string			true	"user pool id"
// @Param		bd		body	request.UserReq	true	"body"
// @Success		200
// @Router		/api/quick/user-pools/{poolId}/users [post]
func (o Controller) createUser(c *gin.Context) {
	var in request.UserReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init user req err")
		return
	}

	var err error
	in.Password, err = safe.HashPassword(in.Password)
	if err != nil {
		resp.ErrorUnknown(c, err, "hash password err")
		return
	}
	in.OpenId = uuid.NewString()
	user, err := o.svc.CreateUser(in.ToModel())
	if err != nil {
		resp.ErrorUnknown(c, err, "create user err")
		return
	}
	resp.SuccessWithData(c, user)
}

// @Summary	modify user
// @Schemes
// @Description	modify user
// @Tags		user
// @Param		poolId	path	string			true	"user pool id"
// @Param		userId	path	string			true	"user id"
// @Param		bd		body	request.UserReq	true	"body"
// @Success		200
// @Router		/api/quick/user-pools/{poolId}/users/{userId} [put]
func (o Controller) modifyUser(c *gin.Context) {
	var in request.UserReq
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init user req err")
		return
	}

	if err := o.svc.ModifyUser(in.ToModel()); err != nil {
		resp.ErrorUnknown(c, err, "modify user err")
		return
	}
	resp.Success(c)
}

// @Summary	delete user
// @Schemes
// @Description	delete user
// @Tags		user
// @Param		poolId	path	string	true	"user pool id"
// @Param		userId	path	string	true	"user id"
// @Success		200
// @Router		/api/quick/user-pools/{poolId}/users/{userId} [delete]
func (o Controller) deleteUser(c *gin.Context) {
	poolId := c.Param("poolId")
	userId := c.Param("userId")
	if err := o.svc.DeleteUser(poolId, userId); err != nil {
		resp.ErrorUnknown(c, err, "delete user err")
		return
	}
	resp.Success(c)
}
