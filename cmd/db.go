package cmd

import (
	"QuickAuth/internal/global"
	"QuickAuth/pkg/model"
	"fmt"
	"os"
	"strings"
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
		os.Exit(ExitMigrate)
		return
	}

	fmt.Println("Gorm auto migrate ok")
}

func createDbTables() {
	if err := initGlobal(); err != nil {
		return
	}
	sqlBytes, err := os.ReadFile("./deploy/postgres/create.sql")
	if err != nil {
		fmt.Println(err)
		os.Exit(ExitReadFile)
	}

	statements := strings.Split(string(sqlBytes), ";")
	for _, sql := range statements {
		if strings.TrimSpace(sql) != "" {
			if err = global.DB.Exec(sql).Error; err != nil {
				fmt.Println(err)
				os.Exit(ExitExecSql)
			}
		}
	}
}
