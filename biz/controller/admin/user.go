package admin

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/admin"
	"QuickAuth/pkg/safe"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ListUser
// @Summary	list users
// @Tags	user
// @Param	X-User-ID	header	string	false	"user id"
// @Param	X-Pool-ID	header	string	false	"user pool id"
// @Param	vhost		query	string	false	"virtual host"
// @Param	poolId		path	integer	true	"user pool id"
// @Success	200
// @Router	/api/quick/user-pools/{poolId}/users [get]
func ListUser(c *gin.Context) {
	var in request.UserReq
	if err := internal.BindUri(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	users, err := admin.ListUser(in.UserPoolID)
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
// @Param	vhost		query	string	false	"virtual host"
// @Param	poolId		path	integer	true	"user pool id"
// @Param	userId		path	string	true	"user id"
// @Success	200
// @Router	/api/quick/user-pools/{poolId}/users/{userId} [get]
func GetUser(c *gin.Context) {
	var in request.UserReq
	if err := internal.BindUri(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	user, err := admin.GetUserById(in.UserPoolID, in.UserID)
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
// @Param	vhost		query	string			false	"virtual host"
// @Param	poolId		path	integer			true	"user pool id"
// @Param	bd			body	request.UserReq	true	"body"
// @Success	200
// @Router	/api/quick/user-pools/{poolId}/users [post]
func CreateUser(c *gin.Context) {
	var in request.UserReq
	if err := internal.BindUriAndJson(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	var err error
	in.Password = safe.HashPassword(in.Password)
	if err != nil {
		resp.ErrorUnknown(c, err, "hash password err")
		return
	}
	in.OpenId = uuid.NewString()
	user, err := admin.CreateUser(in.ToModel())
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
// @Param	vhost		query	string			false	"virtual host"
// @Param	poolId		path	integer			true	"user pool id"
// @Param	userId		path	string			true	"user id"
// @Param	bd			body	request.UserReq	true	"body"
// @Success	200
// @Router	/api/quick/user-pools/{poolId}/users/{userId} [put]
func ModifyUser(c *gin.Context) {
	var in request.UserReq
	if err := internal.BindUriAndJson(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := admin.ModifyUser(in.UserID, in.ToModel()); err != nil {
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
// @Param	vhost		query	string	false	"virtual host"
// @Param	poolId		path	integer	true	"user pool id"
// @Param	userId		path	string	true	"user id"
// @Success	200
// @Router	/api/quick/user-pools/{poolId}/users/{userId} [delete]
func DeleteUser(c *gin.Context) {
	var in request.UserReq
	if err := internal.BindUri(c, &in).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	if err := admin.DeleteUser(in.UserPoolID, in.UserID); err != nil {
		resp.ErrorDelete(c, err, "delete user err")
		return
	}
	resp.Success(c)
}
