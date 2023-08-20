package request

import "QuickAuth/internal/endpoint/model"

type Iam struct {
	ResourceId    int64                       `uri:"resourceId"`
	RoleId        int64                       `uri:"roleId"`
	OperationId   int64                       `uri:"operationId"`
	UserId        string                      `uri:"userId"`
	NodeId        int64                       `uri:"nodeId"`
	ParentId      int64                       `uri:"nodeId"`
	Path          string                      `query:"path" form:"path"`
	Tenant        model.Tenant                `json:"-"`
	Resource      model.Resource              `json:"-"`
	Node          model.ResourceNode          `json:"-"`
	Role          model.ResourceRole          `json:"-"`
	Operation     model.ResourceOperation     `json:"-"`
	RoleOperation model.ResourceRoleOperation `json:"-"`
	UserRole      model.ResourceUserRole      `json:"-"`
	JsonUserRole  model.ResourceJSONUserRole  `json:"-"`
}
