package controller

import (
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/pkg/model"
	"fmt"
	"github.com/gin-contrib/sessions"
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
	var tenant model.Tenant
	if err := o.SetCtx(c).SetTenant(&tenant).Error; err != nil {
		resp.ErrorRequest(c, err, "invalid oidc request param")
		return
	}
	conf := OpenidConfigurationDto{
		Issuer:                            fmt.Sprintf("%s", tenant.Host),
		AuthorizationEndpoint:             fmt.Sprintf("%s/api/quick/oauth2/auth", tenant.Host),
		TokenEndpoint:                     fmt.Sprintf("%s/api/quick/oauth2/token", tenant.Host),
		UserinfoEndpoint:                  fmt.Sprintf("%s/api/quick/me/profile", tenant.Host),
		JwksUri:                           fmt.Sprintf("%s/api/quick/.well-known/jwks.json", tenant.Host),
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

// @Summary	get jwks
// @Schemes
// @Description	get jwks
// @Tags		oidc
// @Success		200
// @Router		/api/quick/me/profile [get]
func (o Controller) getProfile(c *gin.Context) {
	session := sessions.Default(c)
	userId, ok := session.Get("userId").(int64)
	if !ok || userId == 0 {
		resp.ErrorNoLogin(c)
		return
	}

	resp.SuccessWithData(c, userId)
}
