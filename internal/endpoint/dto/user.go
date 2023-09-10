package dto

import (
	"QuickAuth/internal/endpoint/model"
	"time"
)

type UserDtoModel struct {
	ID          string    `json:"id"`
	UserPoolID  int64     `json:"userPoolId"`
	Username    string    `json:"username"`
	DisplayName string    `json:"displayName"`
	Email       *string   `json:"email"`
	Phone       *string   `json:"phone"`
	Type        int32     `json:"type"`
	IsDisabled  bool      `json:"isDisabled"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
}

func UserDto(u model.User) *UserDtoModel {
	return &UserDtoModel{
		ID:          u.ID,
		UserPoolID:  u.UserPoolID,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		Email:       u.Email,
		Phone:       u.Phone,
		Type:        u.Type,
		IsDisabled:  u.IsDisabled,
		CreateTime:  u.CreatedAt,
		UpdateTime:  u.UpdatedAt,
	}
}
