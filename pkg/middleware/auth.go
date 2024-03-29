package middleware

import (
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/biz/service/oauth"
	"QuickAuth/pkg/global"
	"crypto/rsa"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie(resp.CookieIDToken)
		if err != nil {
			resp.ErrorNoLogin(c, err)
			return
		}

		tenant, err := getTenant(c)
		if err != nil {
			resp.ErrorUnknown(c, err, "get gin tenant err")
			return
		}
		keys, err := oauth.LoadRsaPrivateKeys(tenant.App.Name)
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
				setUserInfo(c, token.Claims)
				return // ok
			}
			global.Log.Warn(fmt.Sprintf("%s token valid err: %s", "default", err))
		}

		resp.ErrorInvalidateToken(c)
	}
}

func setUserInfo(c *gin.Context, claims jwt.Claims) {
	claim, ok := claims.(jwt.MapClaims)
	if !ok {
		return
	}

	c.Set(resp.UserInfo, claim)
}

func M2mAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
