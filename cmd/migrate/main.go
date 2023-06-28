package main

import (
	"QuickAuth/internal/global"
	"QuickAuth/pkg/tools/orm"
	"fmt"
)

func main() {
	if err := orm.MigrateDatabase(global.DB); err != nil {
		fmt.Println("Failed to migrate database: ", err)
		return
	}

	fmt.Println("gorm auto migrate ok")
}
