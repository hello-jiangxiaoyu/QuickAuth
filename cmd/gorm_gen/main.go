// Copyright 2023 jiang. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"QuickAuth/pkg/conf"
	"QuickAuth/pkg/gorm"
	"QuickAuth/pkg/utils"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gen"
	"regexp"
	"strings"
)

const (
	genDir   = "../pkg"
	modelDir = genDir + "/models"
	queryDir = genDir + "/query"
)

func main() {
	g, err := getG("quick_auth")
	if err != nil {
		fmt.Println("failed to get generator", err)
		return
	}

	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) { return columnName })
	opt := []gen.ModelOpt{jsonField}
	opt = append(opt, gen.FieldType("grant_types", "pq.StringArray"))
	g.GenerateAllTable(opt...)
	g.Execute()

	_ = utils.AmendFile(modelDir, convertToCamelCase)
}

func getG(dbName string) (*gen.Generator, error) {
	db, err := gorm.NewGormDB(conf.DBPostgres, fmt.Sprintf("host=127.0.0.1 user=admin password=admin dbname=%s port=5432 %s", dbName, ""))
	if err != nil {
		return nil, err
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           queryDir,                                      // 相对执行`go run`时的路径, 会自动创建目录
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface, // 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		FieldNullable:     true,                                          // generate pointer when field is nullable
		FieldCoverable:    false,                                         // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values
		FieldSignable:     false,                                         // detect integer field's unsigned type, adjust generated data type
		FieldWithIndexTag: false,                                         // generate with gorm index tag
		FieldWithTypeTag:  true,                                          // generate with gorm column type tag
	})

	g.UseDB(db)
	return g, err
}

func convertToCamelCase(str []byte) []byte {
	re := regexp.MustCompile(`json:"(\w+)"`)
	res := re.ReplaceAllStringFunc(string(str), func(match string) string {
		word := strings.TrimPrefix(match, `json:"`)
		titleSpace := cases.Title(language.English).String(strings.Replace(word[:len(word)-1], "_", " ", -1))
		camelCase := strings.Replace(titleSpace, " ", "", -1)
		return `json:"` + strings.ToLower(string(camelCase[0])) + camelCase[1:] + `"`
	})
	return []byte(res)
}
