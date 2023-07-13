package main

import (
	"QuickAuth/internal/global"
	"QuickAuth/pkg/model"
	"fmt"
)

func main() {
	migrateList := []any{
		model.User{},
		model.UserPool{},
		model.App{},
		model.AppSecret{},
		model.Tenant{},
	}

	if err := global.DB.AutoMigrate(migrateList...); err != nil {
		return
	}

	fmt.Println("gorm auto migrate ok")
}
