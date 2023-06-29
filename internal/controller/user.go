package controller

import (
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/pkg/model"
	"github.com/gin-gonic/gin"
)

// @Summary	user info
// @Schemes
// @Description	list users
// @Tags		user
// @Success		200
// @Router		/api/quick/user-pool/{poolId}/users [get]
func (o Controller) listUser(c *gin.Context) {
	clients, err := o.svc.ListUser(c.Param("poolId"))
	if err != nil {
		resp.ErrorSelect(c, err, "list user pool err")
		return
	}
	resp.SuccessArray(c, len(clients), clients)
}

// @Summary	user info
// @Schemes
// @Description	list user
// @Tags		user
// @Param		poolId	path	string	true	"user pool id"
// @Param		userId	path	string	true	"user id"
// @Success		200
// @Router		/api/quick/user-pool/{poolId}/users/{userId} [get]
func (o Controller) getUser(c *gin.Context) {
	poolId := c.Param("poolId")
	userId := c.Param("userId")
	client, err := o.svc.GetUserById(poolId, userId)
	if err != nil {
		resp.ErrorUnknown(c, err, "no such user pool")
		return
	}
	resp.SuccessWithData(c, client)
}

// @Summary	create user
// @Schemes
// @Description	create user
// @Tags		user
// @Param		bd		body	model.Client	true	"body"
// @Success		200
// @Router		/api/quick/user-pool/{poolId}/users [post]
func (o Controller) createUser(c *gin.Context) {
	var in model.User
	if err := o.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init user pool req err")
		return
	}
	in.UserPoolID = c.Param("poolId")
	client, err := o.svc.CreateUser(in)
	if err != nil {
		resp.ErrorUnknown(c, err, "create user pool err")
		return
	}
	resp.SuccessWithData(c, client)
}

// @Summary	modify user
// @Schemes
// @Description	modify user
// @Tags		user
// @Param		poolId	path	string			true	"user pool id"
// @Param		userId	path	string			true	"user id"
// @Param		bd		body	model.UserPool	true	"body"
// @Success		200
// @Router		/api/quick/user-pool/{poolId}/users/{userId} [put]
func (o Controller) modifyUser(c *gin.Context) {
	var in model.User
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init user pool req err")
		return
	}
	in.UserPoolID = c.Param("poolId")
	in.ID = c.Param("userId")
	if err := o.svc.ModifyUser(in); err != nil {
		resp.ErrorUnknown(c, err, "modify user pool err")
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
// @Router		/api/quick/user-pool/{poolId}/users/{user} [delete]
func (o Controller) deleteUser(c *gin.Context) {
	poolId := c.Param("poolId")
	userId := c.Param("userId")
	if err := o.svc.DeleteUser(poolId, userId); err != nil {
		resp.ErrorUnknown(c, err, "delete user pool err")
		return
	}
	resp.Success(c)
}
