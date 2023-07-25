package main

import (
	"QuickAuth/pkg/conf"
	"QuickAuth/pkg/orm"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gen"
	"regexp"
	"strings"
)

const (
	genDir   = "../../pkg"
	modelDir = genDir + "/model"
	queryDir = genDir + "/query"
)

func getG(dbName string) (*gen.Generator, error) {
	db, err := orm.NewGormDB(conf.DBPostgres, fmt.Sprintf("host=127.0.0.1 user=admin password=admin dbname=%s port=5432 %s", dbName, ""))
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
		if !strings.Contains(match, "_") {
			return match
		}
		word := strings.TrimPrefix(match, `json:"`)
		titleSpace := cases.Title(language.English).String(strings.Replace(word[:len(word)-1], "_", " ", -1))
		camelCase := strings.Replace(titleSpace, " ", "", -1)
		return `json:"` + strings.ToLower(string(camelCase[0])) + camelCase[1:] + `"`
	})
	return []byte(res)
}
