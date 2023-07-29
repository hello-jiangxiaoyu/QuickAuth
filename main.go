// Copyright 2023 jiang. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"QuickAuth/internal"
	"QuickAuth/internal/global"
	"fmt"
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
	if err = internal.InitGorm(); err != nil {
		fmt.Println("init gorm err: ", err)
		return err
	}
	if _, err = internal.GetValidator(); err != nil {
		fmt.Println("init validator err: ", err)
		return err
	}

	return nil
}

// @title Quick Auth 接口文档
// @version 1.0
// @description
// @license.name MIT
// @license.url https://github.com/hello-jiangxiaoyu/QuickAuth/blob/main/LICENSE
// @securityDefinitions.apikey  Login
// @in                          header
// @name                        token
func main() {
	if err := initSystem(); err != nil {
		return
	}

	svc := internal.GetServer()
	if err := svc.Run(global.Config.Svc.Listen); err != nil {
		fmt.Println("server run err: ", err)
	}

	fmt.Println("server turned off")
}
