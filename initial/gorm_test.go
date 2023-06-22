package initial

import (
	"QuickAuth/conf"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/gen"
	"testing"
)

func TestGenDao(t *testing.T) {
	g, err := getG("quick_auth")
	if err != nil {
		fmt.Println("failed to get generator", err)
		return
	}

	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) { return columnName })
	opt := []gen.ModelOpt{jsonField}
	g.GenerateAllTable(opt...)
	g.Execute()
}

func getG(dbName string) (*gen.Generator, error) {
	db, err := NewGormDB(conf.DBPostgres, fmt.Sprintf("host=127.0.0.1 user=root password=123456 dbname=%s port=5432 %s", dbName, ""))
	if err != nil {
		return nil, err
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           "./query",                                     // 相对执行`go run`时的路径, 会自动创建目录
		OutFile:           "go",                                          // 文件名后缀
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
