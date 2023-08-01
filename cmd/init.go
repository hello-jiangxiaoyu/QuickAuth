package cmd

import (
	"QuickAuth/internal"
	"fmt"
	"os"
)

func initGlobal() error {
	var err error
	defer func() {
		if err != nil {
			os.Exit(ExitInitGlobal)
		}
	}()

	if err = internal.InitConfig(); err != nil {
		fmt.Println("init config err: ", err)
		return err
	}

	if err = internal.InitLogger(); err != nil {
		fmt.Println("init logger err: ", err)
		return err
	}

	if err = internal.InitGorm(); err != nil {
		fmt.Println("init gorm err: ", err)
		return err
	}

	return nil
}
