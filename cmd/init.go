package cmd

import (
	"QuickAuth/biz"
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

	if err = internal.InitLogger(); err != nil {
		fmt.Println("[Error] init logger err: ", err)
		return err
	}

	if err = internal.InitGorm(); err != nil {
		fmt.Println("[Error] init gorm err: ", err)
		return err
	}

	return nil
}
