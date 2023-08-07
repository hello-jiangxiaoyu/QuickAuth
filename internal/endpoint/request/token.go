package request

import "github.com/golang-jwt/jwt/v5"

type (
	AccessClaims struct {
		Nonce string   `json:"nonce,omitempty"`
		Scope []string `json:"scope,omitempty"`
		jwt.RegisteredClaims
	}

	IDClaims struct {
		Nonce         string  `json:"nonce,omitempty"`
		PoolId        int64   `json:"pool_id,omitempty"`
		Name          string  `json:"name,omitempty"`
		NickName      string  `json:"nick_name,omitempty"`
		Gender        string  `json:"gender,omitempty"`                // 性别
		Birthdate     string  `json:"birthdate,omitempty"`             // 出生日期
		Picture       string  `json:"picture,omitempty"`               // URL of the End-User's profile picture
		Email         *string `json:"email,omitempty"`                 // End-User's preferred e-mail address.
		EmailVerified bool    `json:"email_verified,omitempty"`        // True if the End-User's e-mail address has been verified; otherwise false.
		Phone         *string `json:"phone_number,omitempty"`          // End-User's preferred telephone number.
		PhoneVerified bool    `json:"phone_number_verified,omitempty"` // True if the End-User's phone number has been verified; otherwise false.
		Country       string  `json:"country,omitempty"`               // 国家
		Region        string  `json:"region,omitempty"`                // 省
		Locality      string  `json:"locality,omitempty"`              // 城市
		Addr          string  `json:"addr,omitempty"`
		jwt.RegisteredClaims
	}
)
