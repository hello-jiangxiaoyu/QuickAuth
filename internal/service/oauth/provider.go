package oauth

import (
	"QuickAuth/internal/endpoint/model"
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/pkg/idp"
	"QuickAuth/pkg/safe"
	"QuickAuth/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// CreateProviderToken 创建ID Token
func (s *ServiceOauth) CreateProviderToken(app model.App, tenant model.Tenant, user *idp.UserInfo, nonce string) (string, error) {
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

	return s.getTokenString(claims, app.Name)
}

// CheckState 校验state，防止CSRF攻击
func (s *ServiceOauth) CheckState(c *gin.Context) error {
	_, err := c.Cookie(resp.CookieState)
	if err != nil {
		return err
	}

	// todo: validate state by db
	return nil
}
