package request

type Response struct {
	ResourceId  string `uri:"resourceId"`
	RoleId      string `uri:"roleId"`
	UserId      string `uri:"userId"`
	NodeId      string `uri:"nodeId"`
	OperationId string `uri:"operationId"`
}
