package middleware

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/biz/endpoint/resp"
	"QuickAuth/pkg/global"
	"errors"
	"github.com/gin-gonic/gin"
)

func TenantHost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tenant model.Tenant
		host := c.Request.Header.Get(resp.VHost)
		if host == "" {
			host = c.Query(resp.VHost)
		}
		if host == "" {
			host = c.Request.Host
		}
		if err := global.Db().Where("host = ?", host).Preload("App").Preload("UserPool").First(&tenant).Error; err != nil {
			resp.ErrorHost(c)
			return
		}
		c.Set(resp.Tenant, tenant)
	}
}

func getTenant(c *gin.Context) (model.Tenant, error) {
	value, ok := c.Get(resp.Tenant)
	if !ok {
		return model.Tenant{}, errors.New("failed to get gin tenant")
	}
	tenant, ok := value.(model.Tenant)
	if !ok {
		return model.Tenant{}, errors.New("failed to convert gin tenant")
	}
	return tenant, nil
}
