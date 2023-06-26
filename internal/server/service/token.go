package service

import (
	"QuickAuth/pkg/model"
	"QuickAuth/pkg/utils"
	"QuickAuth/pkg/utils/safe"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"github.com/go-jose/go-jose/v3"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	TokenType string `json:"tokenType,omitempty"`
	Nonce     string `json:"nonce,omitempty"`
	Scope     string `json:"scope,omitempty"`
	jwt.RegisteredClaims
}

func LoadRsaPublicKeys(tenant string) (*jose.JSONWebKeySet, error) {
	if tenant == "" {
		return nil, errors.New("tenant name should not be null")
	}
	var err error
	res := map[string][]byte{}
	if res, err = utils.GetJWKs(tenant); err != nil || len(res) == 0 {
		if res, err = utils.GenerateKey(tenant); err != nil {
			return nil, err
		}
	}

	var jwkSet jose.JSONWebKeySet
	var key *rsa.PrivateKey
	for k, v := range res {
		key, err = jwt.ParseRSAPrivateKeyFromPEM(v)
		if err != nil {
			return nil, err
		}

		jwk := jose.JSONWebKey{
			Key:                       key.Public(),
			KeyID:                     k,
			Algorithm:                 "RS256",
			Use:                       "sig",
			Certificates:              []*x509.Certificate{},
			CertificateThumbprintSHA1: []uint8{},
		}
		jwkSet.Keys = append(jwkSet.Keys, jwk)
	}

	return &jwkSet, nil
}

func LoadRsaPrivateKeys(tenantId string) (map[string]*rsa.PrivateKey, error) {
	res, err := utils.GetJWKs(tenantId)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, errors.New("jwk is nil")
	}

	keys := make(map[string]*rsa.PrivateKey)
	for k, v := range res {
		key, err := jwt.ParseRSAPrivateKeyFromPEM(v)
		if err != nil {
			return nil, err
		}
		keys[k] = key
	}

	return keys, nil
}

func CreateAccessToken(client model.Client, tenantId, host, userId, nonce, scope string) (string, error) {
	var token *jwt.Token
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(client.TokenExpire) * time.Hour)
	//refreshExpireTime := nowTime.Add(time.Duration(client.RefreshExpire) * time.Hour)
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
	keys, err := LoadRsaPrivateKeys(tenantId)
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
