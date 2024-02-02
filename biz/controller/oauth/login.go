package oauth

import (
	"QuickAuth/biz/controller/internal"
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/admin"
	"QuickAuth/biz/service/oauth"
	"QuickAuth/pkg/safe"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login
// @Summary	login using username and password
// @Tags	login
// @Param	username	formData	string	true	"username"
// @Param	password	formData	string	true	"password"
// @Param	next		query		string	false	"next"
// @Success	200
// @Router	/api/quick/login [post]
func Login(c *gin.Context) {
	var in request.Login
	if cookie, err := c.Cookie(resp.CookieIDToken); err == nil && cookie != "" {
		resp.DoNothing(c, "user is already logged in, nothing to do")
		return
	}
	if err := internal.New(c).BindForm(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	// 用户名和密码检查
	user, err := admin.GetUserByName(in.Tenant.UserPoolID, in.UserName)
	if err != nil {
		resp.ErrorNotFound(c, err, "no such user")
		return
	}
	if !safe.CheckPasswordHash(in.Password, user.Password) {
		resp.ErrorForbidden(c, "user name or password is incorrect")
		return
	}

	// 生成包含用户信息的id_token
	token, err := oauth.CreateIdToken(in.Tenant.App, in.Tenant, *user, "")
	if err != nil {
		resp.ErrorUnknown(c, err, "create id token err")
		return
	}
	c.SetCookie(resp.CookieIDToken, token, int(in.Tenant.IDExpire), "/api/quick", "", false, true)
	if next := c.Query("next"); next != "" {
		c.Redirect(http.StatusFound, next)
		return
	}
	resp.Success(c)
}

// Logout
// @Summary	logout current user
// @Tags	login
// @Success	200
// @Router	/api/quick/logout [get]
func Logout(c *gin.Context) {
	c.SetCookie(resp.CookieIDToken, "", -1, "/api/quick", "", false, true)
	resp.Success(c)
}

// Register
// @Summary	login using username and password
// @Tags	login
// @Param	username	formData	string	true	"username"
// @Param	password	formData	string	true	"password"
// @Param	next		query		string	false	"next"
// @Success	200
// @Router	/api/quick/register [post]
func Register(c *gin.Context) {
	var in request.Login
	if err := internal.New(c).BindForm(&in).SetTenant(&in.Tenant).Error; err != nil {
		resp.ErrorRequest(c, err)
		return
	}

	var err error
	in.Password, err = safe.HashPassword(in.Password)
	if err != nil {
		resp.ErrorUnknown(c, err, "hash password err")
		return
	}

	user, err := admin.CreateUser(&model.User{
		UserPoolID: in.Tenant.UserPoolID,
		Username:   in.UserName,
		Password:   in.Password,
	})
	if err != nil {
		resp.ErrorCreate(c, err, "create user err")
		return
	}

	resp.SuccessWithData(c, user.Dto())
}
