package middleware

import (
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/internal/global"
	"QuickAuth/pkg/model"
	"github.com/gin-gonic/gin"
)

func TenantHost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tenant model.Tenant
		host := c.Request.Header.Get("vhost")
		if host == "" {
			host = c.Query("vhost")
		}
		if host == "" {
			host = c.Request.Host
		}
		if err := global.DB.Where("host = ?", host).First(&tenant).Error; err != nil {
			resp.ErrorHost(c)
			return
		}
		c.Set("tenant", tenant)
		c.Set("appId", tenant.AppID)
		c.Set("tenantId", tenant.ID)
		c.Set("poolId", tenant.UserPoolID)
	}
}
