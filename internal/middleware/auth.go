package middleware

import (
	"QuickAuth/internal/endpoint/request"
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/internal/global"
	"QuickAuth/internal/service"
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie(resp.IDToken)
		if err != nil {
			resp.ErrorNoLogin(c)
			return
		}

		tenant, err := getTenant(c)
		if err != nil {
			resp.ErrorUnknown(c, err, "get gin tenant err")
			return
		}
		keys, err := service.LoadRsaPrivateKeys(tenant.App.Name)
		if err != nil {
			resp.ErrorUnknown(c, err, "load rsa private key err")
			return
		}

		var key *rsa.PrivateKey
		for _, key = range keys {
			claim := jwt.New(jwt.SigningMethodRS256)
			token, err := jwt.ParseWithClaims(cookie, claim.Claims, func(token *jwt.Token) (interface{}, error) {
				return key.Public(), nil
			})

			if err == nil && token.Valid {
				setUserInfo(c, token)
				return // ok
			}
			global.Log.Warn(fmt.Sprintf("%s token valid err: %s", "default", err))
		}

		resp.ErrorForbidden(c, "invalidated token")
	}
}

func setUserInfo(c *gin.Context, token *jwt.Token) {
	claim, _ := token.Claims.(request.IDClaims)
	c.Set(resp.Claim, claim)
}

func getUserInfo(c *gin.Context) (request.IDClaims, error) {
	value, ok := c.Get(resp.Claim)
	if !ok {
		return request.IDClaims{}, errors.New("failed to get gin IDClaims")
	}
	claim, ok := value.(request.IDClaims)
	if !ok {
		return request.IDClaims{}, errors.New("failed to convert gin IDClaims")
	}
	return claim, nil
}
