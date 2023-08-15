package resp

import (
	"context"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

// ErrorUpdate SQL更新失败
func ErrorUpdate(ctx context.Context, err error, respMsg string, isArray ...bool) {
	if err != nil && strings.HasPrefix(err.Error(), "ERROR: duplicate key value violates unique constraint") {
		errorResponse(ctx, http.StatusInternalServerError, CodeSqlModifyDuplicate, err, respMsg, isArray)
	} else {
		errorResponse(ctx, http.StatusInternalServerError, CodeSqlModify, err, respMsg, isArray)
	}
}

func ErrorCreate(ctx context.Context, err error, respMsg string, isArray ...bool) {
	if err != nil && strings.HasPrefix(err.Error(), "ERROR: duplicate key value violates unique constraint") {
		errorResponse(ctx, http.StatusConflict, CodeSqlCreateDuplicate, err, "Duplicate field name", isArray)
	} else {
		errorResponse(ctx, http.StatusInternalServerError, CodeSqlCreate, err, respMsg, isArray)
	}
}

// ErrorSelect 数据库查询错误
func ErrorSelect(ctx context.Context, err error, respMsg string, isArray ...bool) {
	if err == gorm.ErrRecordNotFound { // gorm find操作record not found
		errorResponse(ctx, http.StatusNotFound, CodeSqlSelectNotFound, err, respMsg, isArray)
	} else {
		errorResponse(ctx, http.StatusInternalServerError, CodeSqlSelect, err, respMsg, isArray)
	}
}

// ErrorDelete 数据库删除错误
func ErrorDelete(ctx context.Context, err error, respMsg string, isArray ...bool) {
	if err == gorm.ErrForeignKeyViolated { // 外键依赖导致无法删除
		errorResponse(ctx, http.StatusConflict, CodeSqlDeleteForKey, err, respMsg, isArray)
	} else {
		errorResponse(ctx, http.StatusInternalServerError, CodeSqlDelete, err, respMsg, isArray)
	}
}
