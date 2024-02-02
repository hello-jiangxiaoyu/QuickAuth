package cmd

import (
	"QuickAuth/biz"
	"QuickAuth/pkg/global"
	"fmt"
	"os"
)

func startServer() {
	if err := initGlobal(); err != nil {
		os.Exit(1)
	}

	if _, err := biz.GetValidator(); err != nil {
		fmt.Println("init validator err: ", err)
		os.Exit(1)
	}

	svc := biz.GetServer()
	if err := svc.Run(global.Config.Svc.Listen); err != nil {
		fmt.Println("server run err: ", err)
		os.Exit(1)
	}

	fmt.Println("[Error] server turned off")
}

func initDefault() {
	if err := initGlobal(); err != nil {
		return
	}

	if err := InitDefaultTenant(); err != nil {
		fmt.Println("init tenant err: ", err)
		os.Exit(1)
	}
	fmt.Println("[OK] init default successfully")
}
func initGlobal() error {
	var err error
	if err = InitLogger(); err != nil {
		fmt.Println("[Error] init logger err: ", err)
		return err
	}

	if err = InitGorm(); err != nil {
		fmt.Println("[Error] init gorm err: ", err)
		return err
	}

	return nil
}
