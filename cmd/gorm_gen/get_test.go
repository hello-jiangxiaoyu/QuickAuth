// Copyright 2023 jiang. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"QuickAuth/pkg/tools/utils"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/gen"
	"testing"
)

func TestGen(*testing.T) {
	generator, err := getG("dev_quick_auth")
	if err != nil {
		fmt.Println("get generator err: ", err)
		return
	}

	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) { return columnName })
	opt := []gen.ModelOpt{jsonField}
	opt = append(opt, gen.FieldType("grant_types", "pq.StringArray"))
	opt = append(opt, gen.FieldType("redirect_uris", "pq.StringArray"))
	generator.GenerateAllTable(opt...)
	generator.Execute()

	if err = utils.AmendFile(modelDir, convertToCamelCase); err != nil {
		panic(err)
	}
}
