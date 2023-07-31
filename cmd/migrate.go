package cmd

import (
	"QuickAuth/internal/global"
	"QuickAuth/pkg/model"
	"fmt"
	"os"
)

func autoMigrateDB() {
	fmt.Println("Start Gorm auto migrate database.")
	migrateList := []any{
		model.User{},
		model.UserPool{},
		model.App{},
		model.AppSecret{},
		model.Tenant{},
	}

	if err := global.DB.Debug().AutoMigrate(migrateList...); err != nil {
		os.Exit(3)
		return
	}

	fmt.Println("Gorm auto migrate ok")
}
