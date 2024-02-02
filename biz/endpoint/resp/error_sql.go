package resp

import (
	"context"
	"net/http"
	"strings"
)

// ErrorUpdate SQL更新失败
func ErrorUpdate(ctx context.Context, err error, respMsg string) {
	if err != nil && strings.Contains(err.Error(), "ERROR: duplicate key value violates unique constraint") {
		errorResponse(ctx, http.StatusInternalServerError, CodeSqlModifyDuplicate, err, respMsg)
	} else {
		errorResponse(ctx, http.StatusInternalServerError, CodeSqlModify, err, respMsg)
	}
}

func ErrorCreate(ctx context.Context, err error, respMsg string) {
	if err != nil && strings.Contains(err.Error(), "ERROR: duplicate key value violates unique constraint") {
		errorResponse(ctx, http.StatusConflict, CodeSqlCreateDuplicate, err, "Duplicate field name")
	} else {
		errorResponse(ctx, http.StatusInternalServerError, CodeSqlCreate, err, respMsg)
	}
}

// ErrorSelect 数据库查询错误
func ErrorSelect(ctx context.Context, err error, respMsg string) {
	if err != nil && strings.Contains(err.Error(), "record not found") { // gorm find操作record not found
		errorResponse(ctx, http.StatusNotFound, CodeSqlSelectNotFound, err, respMsg)
	} else {
		errorResponse(ctx, http.StatusInternalServerError, CodeSqlSelect, err, respMsg)
	}
}

// ErrorDelete 数据库删除错误
func ErrorDelete(ctx context.Context, err error, respMsg string) {
	if err != nil && strings.Contains(err.Error(), "violates foreign key constraint") { // 外键依赖导致无法删除
		errorResponse(ctx, http.StatusConflict, CodeSqlDeleteForKey, err, respMsg)
	} else {
		errorResponse(ctx, http.StatusInternalServerError, CodeSqlDelete, err, respMsg)
	}
}
