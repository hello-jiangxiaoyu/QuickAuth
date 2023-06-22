// Copyright 2023 jiang. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"QuickAuth/global"
	"QuickAuth/initial"
	"fmt"
	"go.uber.org/zap"
	"time"
)

func initSystem() error {
	var err error
	if err = initial.InitConfig(); err != nil {
		return err
	}
	if err = initial.InitLogger(); err != nil {
		return err
	}
	if err = initial.InitGorm(); err != nil {
		return err
	}
	if err = initial.MigrateDatabase(); err != nil {
		return err
	}

	return nil
}

func main() {
	var err error
	for i := 0; i < 3; i++ {
		err = initSystem() // 系统初始化，重试3次失败则退出
		if err == nil {
			break
		}
		fmt.Println("init system err: ", err, "\n\nwaiting for starting...")
		time.Sleep(time.Second * 10)
	}

	if err != nil {
		fmt.Println("failed to init system, down")
		return
	}

	svc := initial.GetServer()
	runAddr := global.Config.Svc.Listen
	if err = svc.Run(runAddr); err != nil {
		global.Log.Error("server run err: ", zap.Error(err))
	}

	global.Log.Info("server turned off")
}
