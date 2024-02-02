package oauth

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/idp"
	"QuickAuth/pkg/safe"
	"QuickAuth/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// CreateProviderToken 创建ID Token
func CreateProviderToken(app model.App, tenant model.Tenant, user *idp.UserInfo, nonce string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(tenant.IDExpire) * time.Hour)
	claims := request.IDClaims{
		Nonce:   nonce,
		Name:    user.DisplayName,
		Picture: user.AvatarUrl,
		Email:   &user.Email,
		Phone:   &user.Phone,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    utils.GetUrlByHost(tenant.Host) + resp.ApiPrefix,
			Subject:   user.UnionId,
			Audience:  []string{app.ID},
			ExpiresAt: jwt.NewNumericDate(expireTime),
			NotBefore: jwt.NewNumericDate(nowTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			ID:        app.Name + "-" + safe.Rand62(31),
		},
	}

	return getTokenString(claims, app.Name)
}

// CheckState 校验state，防止CSRF攻击
func CheckState(c *gin.Context) error {
	_, err := c.Cookie(resp.CookieState)
	if err != nil {
		return err
	}

	// todo: validate state by db
	return nil
}
