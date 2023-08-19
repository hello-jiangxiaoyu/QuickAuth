// Copyright 2023 jiang. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"QuickAuth/pkg/utils"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"os"
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

	// model定义
	app := generator.GenerateModel("apps", opt...)
	tenant := generator.GenerateModel("tenants", opt...)

	appSecret := generator.GenerateModel("app_secrets", append(opt,
		gen.FieldRelate(field.HasOne, "App", app, &field.RelateConfig{}),
		gen.FieldType("scope", "pq.StringArray"),
	)...)

	code := generator.GenerateModel("codes", append(opt,
		gen.FieldRelate(field.HasOne, "App", app, &field.RelateConfig{}),
		gen.FieldType("scope", "pq.StringArray"),
	)...)

	provider := generator.GenerateModel("providers", append(opt,
		gen.FieldRelate(field.HasOne, "App", app, &field.RelateConfig{}),
		gen.FieldRelate(field.HasOne, "Tenant", tenant, &field.RelateConfig{}),
	)...)
	userPool := generator.GenerateModel("user_pools", append(opt,
		gen.FieldRelate(field.HasMany, "Tenant", tenant, &field.RelateConfig{}),
	)...)
	user := generator.GenerateModel("users", append(opt,
		gen.FieldRelate(field.HasOne, "UserPool", userPool, &field.RelateConfig{}),
	)...)

	// app 依赖修正
	app = generator.GenerateModel("apps", append(opt,
		gen.FieldRelate(field.HasMany, "Tenant", tenant, &field.RelateConfig{}),
		gen.FieldRelate(field.HasMany, "Secret", appSecret, &field.RelateConfig{}),
	)...)

	// tenant 依赖修正
	tenant = generator.GenerateModel("tenants", append(opt,
		gen.FieldRelate(field.HasOne, "App", app, &field.RelateConfig{}),
		gen.FieldRelate(field.HasOne, "UserPool", userPool, &field.RelateConfig{}),
		gen.FieldType("grant_types", "pq.StringArray"),
		gen.FieldType("redirect_uris", "pq.StringArray"),
	)...)

	// ========== 权限表 ==========
	resource := generator.GenerateModel("resources", append(opt,
		gen.FieldRelate(field.HasOne, "Tenant", tenant, &field.RelateConfig{}),
	)...)
	resourceNode := generator.GenerateModel("resource_nodes", append(opt,
		gen.FieldRelate(field.HasOne, "Tenant", tenant, &field.RelateConfig{}),
	)...)
	resourceRole := generator.GenerateModel("resource_roles", append(opt,
		gen.FieldRelate(field.HasOne, "Tenant", tenant, &field.RelateConfig{}),
	)...)
	resourceOperation := generator.GenerateModel("resource_operations", append(opt,
		gen.FieldRelate(field.HasOne, "Tenant", tenant, &field.RelateConfig{}),
	)...)
	resourceRoleOperation := generator.GenerateModel("resource_role_operations", append(opt,
		gen.FieldRelate(field.HasOne, "Tenant", tenant, &field.RelateConfig{}),
	)...)
	resourceUserRole := generator.GenerateModel("resource_user_roles", append(opt,
		gen.FieldRelate(field.HasOne, "Tenant", tenant, &field.RelateConfig{}),
	)...)
	resourceJsonUserRole := generator.GenerateModel("resource_json_user_roles", append(opt,
		gen.FieldRelate(field.HasOne, "Tenant", tenant, &field.RelateConfig{}),
	)...)

	// 生成model
	generator.ApplyBasic(app, appSecret, code, provider, tenant, userPool, user,
		resource, resourceNode, resourceRole, resourceOperation,
		resourceRoleOperation, resourceUserRole, resourceJsonUserRole)
	generator.Execute()

	// 删除query目录
	if _, err = os.Stat(queryDir); err == nil {
		if err = os.RemoveAll(queryDir); err != nil {
			panic(err)
		}
	}

	// 修正model文件json tag，将下划线转为小驼峰
	if err = utils.AmendFile(modelDir, convertToCamelCase); err != nil {
		panic(err)
	}
}
