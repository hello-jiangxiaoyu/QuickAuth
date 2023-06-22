package main

import (
	"QuickAuth/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestGin(t *testing.T) {
	if err := initSystem(); err != nil {
		return
	}
	r := gin.New()
	r.Use(middleware.Recovery())
	var err error
	r.GET("/", func(c *gin.Context) {
		fmt.Println(err.Error())
		c.String(http.StatusOK, "ok")
	})
	_ = r.Run(":1000")
}
