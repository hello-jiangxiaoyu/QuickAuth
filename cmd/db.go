package cmd

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/pkg/global"
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

	if err := global.Db().Debug().AutoMigrate(migrateList...); err != nil {
		fmt.Println("[Error] migrate err: ", err)
		os.Exit(1)
		return
	}

	fmt.Println("[OK] Gorm auto migrate ok")
}

func createDbTables() {
	if err := initGlobal(); err != nil {
		os.Exit(1)
	}
	sqlBytes, err := os.ReadFile("./deploy/postgres/create.sql")
	if err != nil {
		fmt.Println("[Error] read file err: ", err)
		os.Exit(1)
	}

	statements := strings.Split(string(sqlBytes), ";")
	for _, sql := range statements {
		if strings.TrimSpace(sql) != "" {
			if err = global.Db().Exec(sql).Error; err != nil {
				fmt.Println("[Error] exec sql err: ", err)
				os.Exit(1)
				return
			}
		}
	}

	fmt.Println("[OK] create database table by sql ok")
}
