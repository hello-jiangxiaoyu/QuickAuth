package resp

import "net/http"

const (
	IDToken   = "id_token"
	Tenant    = "tenant"
	VHost     = "vhost"
	UserInfo  = "claim"
	ApiPrefix = "/api/quick"

	MsgSuccess = ""
)

const (
	CodeSuccess = http.StatusOK // 成功
	CodeAccept  = http.StatusAccepted

	// ========== 系统相关错误码 ==========

	CodeNoSuchRoute = 1000
	CodeRequestPara = 1001
	CodeForbidden   = 1002

	// ========== 业务相关错误码 ==========

	CodeNoSuchHost   = 2000
	CodeNotLogin     = 2001
	CodeInvalidToken = 2002

	// ========== Sql相关错误码 ==========

	CodeSqlSelect          = 3000 // 查询失败
	CodeSqlSelectNotFound  = 3001 // 不存在该数据
	CodeSqlModify          = 3002 // 修改失败
	CodeSqlModifyDuplicate = 3003 // 数据冲突导致修改失败
	CodeSqlCreate          = 3004 // 创建失败
	CodeSqlCreateDuplicate = 3005 // 数据重复导致创建失败
	CodeSqlDelete          = 3006 // 删除失败
	CodeSqlDeleteForKey    = 3007 // 外键依赖导致删除失败

	// ========== 请求相关错误码 ==========

	// ========== error相关错误码 ==========

	CodeServerPanic = 5000 // panic异常
	CodeUnknown     = 5001 // 未知错误
	CodeNotFound    = 5002 // 资源不存在
	CodeSaveSession = 5003 // 保存session失败
)
