package service

import (
	"QuickAuth/pkg/model"
	"QuickAuth/pkg/safe"
	"QuickAuth/pkg/utils"
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	TokenType string   `json:"tokenType,omitempty"`
	Nonce     string   `json:"nonce,omitempty"`
	Scope     []string `json:"scope,omitempty"`
	jwt.RegisteredClaims
}

func (s *Service) CreateAccessToken(app model.App, tenantName, host, userId, nonce string, scope []string) (string, error) {
	var token *jwt.Token
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(24) * time.Hour)
	claims := Claims{
		TokenType: "access-token",
		Nonce:     nonce,
		Scope:     scope,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    utils.GetUrlByHost(host),
			Subject:   userId,
			Audience:  []string{app.ID},
			ExpiresAt: jwt.NewNumericDate(expireTime),
			NotBefore: jwt.NewNumericDate(nowTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			ID:        app.Name + "-" + safe.Rand62(31),
		},
	}
	token = jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	keys, err := s.LoadRsaPrivateKeys(tenantName)
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
