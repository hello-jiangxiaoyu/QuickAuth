package service

import (
	"QuickAuth/pkg/models"
	"QuickAuth/pkg/utils"
	"QuickAuth/pkg/utils/safe"
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	TokenType string `json:"tokenType,omitempty"`
	Nonce     string `json:"nonce,omitempty"`
	Scope     string `json:"scope,omitempty"`
	jwt.RegisteredClaims
}

func (s *Service) CreateAccessToken(client models.Client, tenantId, host, userId, nonce, scope string) (string, error) {
	var token *jwt.Token
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(client.TokenExpire) * time.Hour)
	claims := Claims{
		TokenType: "access-token",
		Nonce:     nonce,
		Scope:     scope,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    utils.GetUrlByHost(host),
			Subject:   userId,
			Audience:  []string{client.ID},
			ExpiresAt: jwt.NewNumericDate(expireTime),
			NotBefore: jwt.NewNumericDate(nowTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			ID:        client.Name + "-" + safe.Rand62(31),
		},
	}
	token = jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	keys, err := s.LoadRsaPrivateKeys(tenantId)
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
