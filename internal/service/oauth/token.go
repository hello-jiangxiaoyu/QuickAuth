package oauth

import (
	"QuickAuth/internal/endpoint/model"
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/pkg/safe"
	"QuickAuth/pkg/utils"
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// CreateAccessToken 创建Access Token
func (s *ServiceOauth) CreateAccessToken(app model.App, tenant model.Tenant, userId string, nonce string, scope []string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(tenant.AccessExpire) * time.Second)
	claims := request.AccessClaims{
		Nonce: nonce,
		Scope: scope,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    utils.GetUrlByHost(tenant.Host),
			Subject:   userId,
			Audience:  []string{app.ID},
			ExpiresAt: jwt.NewNumericDate(expireTime),
			NotBefore: jwt.NewNumericDate(nowTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			ID:        app.Name + "-" + safe.Rand62(31),
		},
	}

	return s.getTokenString(claims, app.Name)
}

// CreateIdToken 创建ID Token
func (s *ServiceOauth) CreateIdToken(app model.App, tenant model.Tenant, user model.User, nonce string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(tenant.IDExpire) * time.Hour)
	birthdate := ""
	if user.Birthdate != nil {
		birthdate = user.Birthdate.Format("2006-01-02")
	}
	claims := request.IDClaims{
		Nonce:     nonce,
		Name:      user.DisplayName,
		NickName:  user.NickName,
		Gender:    user.Gender,
		Birthdate: birthdate,
		Picture:   user.Avatar,
		Email:     user.Email,
		Addr:      user.Addr,
		Phone:     user.Phone,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    utils.GetUrlByHost(tenant.Host) + resp.ApiPrefix,
			Subject:   user.ID,
			Audience:  []string{app.ID},
			ExpiresAt: jwt.NewNumericDate(expireTime),
			NotBefore: jwt.NewNumericDate(nowTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			ID:        app.Name + "-" + safe.Rand62(31),
		},
	}

	return s.getTokenString(claims, app.Name)
}

func (s *ServiceOauth) getTokenString(claims jwt.Claims, appName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	keys, err := LoadRsaPrivateKeys(appName)
	if err != nil {
		return "", err
	}

	var kid string
	var key *rsa.PrivateKey
	for kid, key = range keys {
		break
	}

	token.Header["kid"] = kid
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
