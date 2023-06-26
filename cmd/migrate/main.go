package main

import (
	"QuickAuth/internal/global"
	"QuickAuth/pkg/gorm"
	"fmt"
)

func main() {
	if err := gorm.MigrateDatabase(global.DB); err != nil {
		fmt.Println("Failed to migrate database: ", err)
		return
	}

	fmt.Println("gorm auto migrate ok")
}
