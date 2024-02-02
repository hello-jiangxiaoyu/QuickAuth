package admin

import (
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/pkg/safe"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ListUser
// @Summary	list users
// @Tags	user
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	poolId		path	integer	true	"user pool id"
// @Success	200
// @Router	/api/quick/user-pools/{poolId}/users [get]
func (a Route) ListUser(c *gin.Context) {
	var in request.UserReq
	if err := a.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	users, err := a.svc.ListUser(in.UserPoolID)
	if err != nil {
		resp.ErrorSelect(c, err, "list user err")
		return
	}
	resp.SuccessArrayData(c, len(users), users)
}

// GetUser
// @Summary	get user info
// @Tags	user
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	poolId		path	integer	true	"user pool id"
// @Param	userId		path	string	true	"user id"
// @Success	200
// @Router	/api/quick/user-pools/{poolId}/users/{userId} [get]
func (a Route) GetUser(c *gin.Context) {
	var in request.UserReq
	if err := a.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	user, err := a.svc.GetUserById(in.UserPoolID, in.UserID)
	if err != nil {
		resp.ErrorSelect(c, err, "no such user")
		return
	}
	resp.SuccessWithData(c, user)
}

// CreateUser
// @Summary	create user
// @Tags	user
// @Param	X-User-ID	header	string			false	"user id"
// @Param	X-Pool-ID	header	string			false	"user pool id"
// @Param	poolId		path	integer			true	"user pool id"
// @Param	bd			body	request.UserReq	true	"body"
// @Success	200
// @Router	/api/quick/user-pools/{poolId}/users [post]
func (a Route) CreateUser(c *gin.Context) {
	var in request.UserReq
	if err := a.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	var err error
	in.Password, err = safe.HashPassword(in.Password)
	if err != nil {
		resp.ErrorUnknown(c, err, "hash password err")
		return
	}
	in.OpenId = uuid.NewString()
	user, err := a.svc.CreateUser(in.ToModel())
	if err != nil {
		resp.ErrorCreate(c, err, "create user err")
		return
	}

	resp.SuccessWithData(c, user.Dto())
}

// ModifyUser
// @Summary	modify user
// @Tags	user
// @Param	X-User-ID	header	string			false	"user id"
// @Param	X-Pool-ID	header	string			false	"user pool id"
// @Param	poolId		path	integer			true	"user pool id"
// @Param	userId		path	string			true	"user id"
// @Param	bd			body	request.UserReq	true	"body"
// @Success	200
// @Router	/api/quick/user-pools/{poolId}/users/{userId} [put]
func (a Route) ModifyUser(c *gin.Context) {
	var in request.UserReq
	if err := a.SetCtx(c).BindUriAndJson(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := a.svc.ModifyUser(in.UserID, in.ToModel()); err != nil {
		resp.ErrorUpdate(c, err, "modify user err")
		return
	}
	resp.Success(c)
}

// DeleteUser
// @Summary	delete user
// @Tags	user
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	poolId		path	integer	true	"user pool id"
// @Param	userId		path	string	true	"user id"
// @Success	200
// @Router	/api/quick/user-pools/{poolId}/users/{userId} [delete]
func (a Route) DeleteUser(c *gin.Context) {
	var in request.UserReq
	if err := a.SetCtx(c).BindUri(&in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := a.svc.DeleteUser(in.UserPoolID, in.UserID); err != nil {
		resp.ErrorDelete(c, err, "delete user err")
		return
	}
	resp.Success(c)
}
