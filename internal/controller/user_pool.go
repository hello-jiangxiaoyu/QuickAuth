package controller

import (
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/pkg/model"
	"github.com/gin-gonic/gin"
)

// @Summary	user pools info
// @Schemes
// @Description	list user pool
// @Tags		user
// @Success		200
// @Router		/api/quick/user-pool [get]
func (o Controller) listUserPool(c *gin.Context) {
	clients, err := o.svc.ListUserPool()
	if err != nil {
		resp.ErrorSelect(c, err, "list user pool err")
		return
	}
	resp.SuccessArray(c, len(clients), clients)
}

// @Summary	user pool info
// @Schemes
// @Description	list user pool
// @Tags		user
// @Param		poolId	path	string	true	"user pool id"
// @Success		200
// @Router		/api/quick/user-pool/{poolId} [get]
func (o Controller) getUserPool(c *gin.Context) {
	client, err := o.svc.GetUserPool(c.Param("poolId"))
	if err != nil {
		resp.ErrorUnknown(c, err, "no such user pool")
		return
	}
	resp.SuccessWithData(c, client)
}

// @Summary	create user pool
// @Schemes
// @Description	create user pool
// @Tags		user
// @Param		bd		body	model.Client	true	"body"
// @Success		200
// @Router		/api/quick/user-pool [post]
func (o Controller) createUserPool(c *gin.Context) {
	var in model.UserPool
	if err := o.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init user pool req err")
		return
	}
	client, err := o.svc.CreateUserPool(in)
	if err != nil {
		resp.ErrorUnknown(c, err, "create user pool err")
		return
	}
	resp.SuccessWithData(c, client)
}

// @Summary	modify user pool
// @Schemes
// @Description	modify user pool
// @Tags		user
// @Param		poolId	path	string	true	"user pool id"
// @Param		bd		body	model.UserPool	true	"body"
// @Success		200
// @Router		/api/quick/user-pool/{poolId} [put]
func (o Controller) modifyUserPool(c *gin.Context) {
	var in model.UserPool
	if err := o.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err, "init user pool req err")
		return
	}
	in.ID = c.Param("poolId")
	if err := o.svc.ModifyUserPool(in); err != nil {
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
// @Router		/api/quick/user-pool/{poolId} [delete]
func (o Controller) deleteUserPool(c *gin.Context) {
	if err := o.svc.DeleteUserPool(c.Param("poolId")); err != nil {
		resp.ErrorUnknown(c, err, "delete user pool err")
		return
	}
	resp.Success(c)
}
