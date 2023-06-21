package controller

import (
	"QuickAuth/global"
	"QuickAuth/response"
	"QuickAuth/server/model"
	"QuickAuth/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewOauth2Router(r *gin.RouterGroup) {
	r.GET("/v1/:tenant/.well-known/openid-configuration", OIDC)
	r.GET("/v1/:tenant/.well-known/jwks.json", GetJwks)
}

func OIDC(c *gin.Context) {
	tenantName := "default"
	prefix := utils.GetHostWithScheme(c)
	conf := model.OpenidConfigurationDto{
		Issuer:                            fmt.Sprintf("%s/%s", prefix, tenantName),
		AuthorizationEndpoint:             fmt.Sprintf("%s/%s/oauth2/auth", prefix, tenantName),
		TokenEndpoint:                     fmt.Sprintf("%s/%s/oauth2/token", prefix, tenantName),
		UserinfoEndpoint:                  fmt.Sprintf("%s/%s/me/profile", prefix, tenantName),
		JwksUri:                           fmt.Sprintf("%s/%s/.well-known/jwks.json", prefix, tenantName),
		ScopesSupported:                   []string{"openid", "profile", "email", "offline_access"},
		ResponseTypesSupported:            []string{"code", "id_token", "code id_token", "id_token token"},
		SubjectTypesSupported:             []string{"pairwise"},
		IdTokenSigningAlgValuesSupported:  []string{"RS256"},
		TokenEndpointAuthMethodsSupported: []string{"client_secret_basic", "client_secret_post"},
		ClaimsSupported:                   []string{"sub", "iss", "aud", "exp", "iat", "nonce", "name", "email"},
		RequestUriParameterSupported:      false,
	}
	c.JSON(http.StatusOK, conf)
}

func GetJwks(c *gin.Context) {
	tenantName := "default"
	jwks, err := utils.LoadRsaPublicKeys(tenantName)
	if err != nil {
		response.ErrorUnknown(c, "failed to get pub keys")
		global.Log.Error("get jwks err: " + err.Error())
		return
	}

	response.SuccessWithData(c, jwks)
}
