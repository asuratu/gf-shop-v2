package model

type BaseRolePermission struct {
	Id           uint `orm:"id,primary" json:"id"`               // 自增ID
	RoleId       uint `orm:"role_id" json:"role_id"`             // 角色ID
	PermissionId uint `orm:"permission_id" json:"permission_id"` // 权限ID
}

// RolePermissionEntity 适用于查询单个角色的权限列表
type RolePermissionEntity struct {
	RolePermission *BaseRolePermission `json:"role_permission"`
	Permission     *BasePermission     `json:"permission"`
}

type BaseRolePermissionEntity struct {
	Id             uint   `orm:"id,primary" json:"id"`               // 自增ID
	RoleId         uint   `orm:"role_id" json:"role_id"`             // 角色ID
	PermissionId   uint   `orm:"permission_id" json:"permission_id"` // 权限ID
	PermissionName string `json:"permission_name"`                   // 权限名称
	PermissionPath string `json:"permission_path"`                   // 权限路径
}

// RolePermissionAllEntity 适用于查询多个角色的权限列表
type RolePermissionAllEntity struct {
	Role           *BaseRole                   `json:"role"`
	RolePermission []*BaseRolePermissionEntity `orm:"with:role_id:id" json:"role_permission"`
}
