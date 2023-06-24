package middleware

import (
	"QuickAuth/internal/endpoint/response"
	"QuickAuth/internal/global"
	"QuickAuth/internal/model"
	"github.com/gin-gonic/gin"
)

func TenantHost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tenant model.Tenant
		if err := global.DB.Where("host = ?", c.Request.Host).
			First(&tenant).Error; err != nil {
			response.ErrorHost(c)
			return
		}
		c.Set("tenant", tenant)
	}
}
