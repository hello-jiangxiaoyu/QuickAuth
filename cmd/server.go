package cmd

import (
	"QuickAuth/internal"
	"QuickAuth/internal/global"
	"fmt"
	"os"
)

func initSystem() error {
	var err error
	if err = internal.InitConfig(); err != nil {
		fmt.Println("init config err: ", err)
		return err
	}
	if err = internal.InitLogger(); err != nil {
		fmt.Println("init logger err: ", err)
		return err
	}
	if _, err = internal.GetValidator(); err != nil {
		fmt.Println("init validator err: ", err)
		return err
	}
	if err = internal.InitGorm(); err != nil {
		fmt.Println("init gorm err: ", err)
		return err
	}
	if err = internal.InitDefaultTenant(); err != nil {
		fmt.Println("init tenant err: ", err)
		return err
	}

	return nil
}

func startServer() {
	if err := initSystem(); err != nil {
		return
	}

	svc := internal.GetServer()
	if err := svc.Run(global.Config.Svc.Listen); err != nil {
		fmt.Println("server run err: ", err)
		os.Exit(2)
	}

	fmt.Println("server turned off")
}
