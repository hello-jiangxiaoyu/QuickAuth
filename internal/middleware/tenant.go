package middleware

import (
	"QuickAuth/internal/endpoint/resp"
	"QuickAuth/internal/global"
	"QuickAuth/pkg/models"
	"github.com/gin-gonic/gin"
)

func TenantHost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tenant models.Tenant
		if err := global.DB.Where("host = ?", c.Request.Host).
			First(&tenant).Error; err != nil {
			resp.ErrorHost(c)
			return
		}
		c.Set("tenant", tenant)
	}
}
