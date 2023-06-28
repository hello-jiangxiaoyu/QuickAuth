package controller

import (
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/internal/server/controller/internal"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OpenidConfigurationDto struct {
	Issuer                            string   `json:"issuer"`
	AuthorizationEndpoint             string   `json:"authorization_endpoint"`
	TokenEndpoint                     string   `json:"token_endpoint"`
	UserinfoEndpoint                  string   `json:"userinfo_endpoint"`
	JwksUri                           string   `json:"jwks_uri"`
	ScopesSupported                   []string `json:"scopes_supported"`
	ResponseTypesSupported            []string `json:"response_types_supported"`
	SubjectTypesSupported             []string `json:"subject_types_supported"`
	IdTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
	ClaimsSupported                   []string `json:"claims_supported"`
	RequestUriParameterSupported      bool     `json:"request_uri_parameter_supported"`
}

// @Summary	get OIDC
// @Schemes
// @Description	get open id configuration
// @Tags		oidc
// @Success		200		{object}	OpenidConfigurationDto
// @Router		/api/quick/.well-known/openid-configuration [get]
func (o Controller) getOIDC(c *gin.Context) {
	tenantName := "default"
	prefix := internal.GetHostWithScheme(c)
	conf := OpenidConfigurationDto{
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

// @Summary	get jwks
// @Schemes
// @Description	get jwks
// @Tags		oidc
// @Success		200
// @Router		/api/quick/.well-known/jwks.json [get]
func (o Controller) getJwks(c *gin.Context) {
	tenantName := "default"
	jwks, err := o.svc.LoadRsaPublicKeys(tenantName)
	if err != nil {
		resp.ErrorUnknown(c, err, "failed to get pub keys")
		return
	}

	resp.SuccessWithData(c, jwks)
}
