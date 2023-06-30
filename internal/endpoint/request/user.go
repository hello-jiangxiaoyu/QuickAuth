package request

import "QuickAuth/pkg/model"

type UserReq struct {
	UserID      int64  `json:"-" uri:"userId"`
	UserPoolID  int64  `json:"-" uri:"poolId"`
	OpenId      string `json:"-"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Type        int32  `json:"type"`
}

func (u *UserReq) ToModel() model.User {
	return model.User{
		ID:          u.UserID,
		OpenID:      u.OpenId,
		UserPoolID:  u.UserPoolID,
		Username:    u.Username,
		Password:    u.Password,
		DisplayName: &u.DisplayName,
		Email:       &u.Email,
		Phone:       &u.Phone,
		Type:        u.Type,
	}
}

type UserPoolReq struct {
	PoolId   int64  `json:"-" uri:"poolId"`
	Name     string `json:"name"`
	Describe string `json:"describe"`
}

func (u *UserPoolReq) ToModel() model.UserPool {
	return model.UserPool{
		ID:       u.PoolId,
		Name:     u.Name,
		Describe: &u.Describe,
	}
}
