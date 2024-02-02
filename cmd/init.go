package cmd

import (
	"QuickAuth/biz"
	"fmt"
)

func initGlobal() error {
	var err error
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
