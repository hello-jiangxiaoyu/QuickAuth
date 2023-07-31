// Copyright 2023 jiang. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"QuickAuth/cmd"
)

// @title Quick Auth 接口文档
// @version 1.0
// @description
// @license.name MIT
// @license.url https://github.com/hello-jiangxiaoyu/QuickAuth/blob/main/LICENSE
// @securityDefinitions.apikey  Login
// @in                          header
// @name                        token
func main() {
	cmd.Execute()
}
