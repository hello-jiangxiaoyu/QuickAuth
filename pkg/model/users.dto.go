package model

type (
	UserDto struct {
		ID          string  `json:"id"`
		Username    string  `json:"username"`
		DisplayName string  `json:"displayName"`
		Email       *string `json:"email"`
		Phone       *string `json:"phone"`
		Avatar      string  `json:"avatar"`
	}
)

func (u *User) Dto() *UserDto {
	return &UserDto{
		ID:          u.ID,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		Email:       u.Email,
		Phone:       u.Phone,
		Avatar:      u.Avatar,
	}
}
