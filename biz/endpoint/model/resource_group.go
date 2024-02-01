package model

type (
	ResourceGroup struct {
		Id          int64  `gorm:"primaryKey" json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		TenantId    int64  `gorm:"primaryKey" json:"tenantId"`
	}

	ResourceGroupResource struct {
		Id          int64         `gorm:"primaryKey" json:"id"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		GroupId     int64         `json:"groupId" uri:"groupId"`
		Group       ResourceGroup `gorm:"foreignKey:GroupId, TenantId" json:"-" swaggerignore:"true"`
		TenantId    int64         `gorm:"primaryKey" json:"tenantId"`
	}
	ResourceGroupRole struct {
		Id          int64         `gorm:"primaryKey" json:"id"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		GroupId     int64         `json:"groupId" uri:"groupId"`
		Group       ResourceGroup `gorm:"foreignKey:GroupId, TenantId" json:"-" swaggerignore:"true"`
		TenantId    int64         `gorm:"primaryKey" json:"tenantId"`
	}
	ResourceGroupAction struct {
		Id          int64         `gorm:"primaryKey" json:"id"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		GroupId     int64         `json:"groupId" uri:"groupId"`
		Group       ResourceGroup `gorm:"foreignKey:GroupId, TenantId" json:"-" swaggerignore:"true"`
		TenantId    int64         `gorm:"primaryKey" json:"tenantId"`
	}

	ResourceGroupRoleAction struct {
		Id         int64               `gorm:"primaryKey" json:"id"`
		RoleId     int64               `json:"roleId" uri:"roleId"`
		Role       ResourceGroupRole   `gorm:"foreignKey:RoleId, TenantId" json:"-" swaggerignore:"true"`
		ActionId   int64               `json:"actionId" uri:"actionId"`
		Action     ResourceGroupAction `gorm:"foreignKey:ActionId, TenantId" json:"-" swaggerignore:"true"`
		ActionName string              `gorm:"<-:false;-:migration" json:"actionName"`
		RoleName   string              `gorm:"<-:false;-:migration" json:"roleName"`
		TenantId   int64               `gorm:"primaryKey" json:"tenantId"`
	}
	ResourceGroupUser struct {
		Id                int64             `gorm:"primaryKey" json:"id"`
		GroupId           int64             `json:"groupId" uri:"groupId"`
		ResourceGroup     ResourceGroup     `gorm:"foreignKey:GroupId, TenantId" json:"-" swaggerignore:"true"`
		ResourceGroupName string            `json:"resourceGroupName" gorm:"<-:false;-:migration"`
		RoleId            int64             `json:"roleId" uri:"roleId"`
		Role              ResourceGroupRole `gorm:"foreignKey:RoleId, TenantId" json:"-" swaggerignore:"true"`
		RoleName          string            `json:"roleName" gorm:"<-:false;-:migration"`
		UserId            int64             `json:"userId" uri:"userId"`
		TenantId          int64             `gorm:"primaryKey" json:"tenantId"`
		Sub               string            `json:"sub" gorm:"<-:false;-:migration"`
		DisplayName       string            `json:"displayName" gorm:"<-:false;-:migration"`
	}
)

type RequestResourceGroup struct {
	Tenant      Tenant  `json:"-" swaggerignore:"true"`
	GroupId     int64   `json:"groupId" uri:"groupId" swaggerignore:"true"`
	ResourceId  int64   `json:"resourceId" uri:"resourceId" swaggerignore:"true"`
	ActionId    int64   `json:"actionId" uri:"actionId" swaggerignore:"true"`
	RoleId      int64   `json:"roleId" uri:"roleId"`
	UserId      int64   `json:"userId" uri:"userId" swaggerignore:"true"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Uid         int64   `json:"uid"`
	ActionIds   []int64 `json:"actionIds"`
	UserIds     []int64 `json:"userIds"`
}
