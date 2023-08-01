package cmd

import (
	"QuickAuth/internal"
	"QuickAuth/internal/global"
	"fmt"
	"os"
)

func startServer() {
	if err := initGlobal(); err != nil {
		return
	}

	if _, err := internal.GetValidator(); err != nil {
		fmt.Println("init validator err: ", err)
		os.Exit(ExitServer)
		return
	}

	if err := internal.InitDefaultTenant(); err != nil {
		fmt.Println("init tenant err: ", err)
		os.Exit(ExitServer)
		return
	}

	svc := internal.GetServer()
	if err := svc.Run(global.Config.Svc.Listen); err != nil {
		fmt.Println("server run err: ", err)
		os.Exit(ExitServer)
		return
	}

	fmt.Println("server turned off")
}
