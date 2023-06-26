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
		return err
	}
	if err = internal.InitLogger(); err != nil {
		return err
	}
	if err = internal.InitGorm(); err != nil {
		return err
	}

	return nil
}

// @title swagger 接口文档
// @version 1.0
// @description
// @license.name MIT
// @license.url https://github.com/hello-jiangxiaoyu/QuickAuth/blob/main/LICENSE
// @securityDefinitions.apikey  Login
// @in                          header
// @name                        token
func main() {
	if err := initSystem(); err != nil {
		fmt.Println("init system err: ", err)
		return
	}

	svc := internal.GetServer()
	if err := svc.Run(global.Config.Svc.Listen); err != nil {
		fmt.Println("server run err: ", err)
	}

	fmt.Println("server turned off")
}
