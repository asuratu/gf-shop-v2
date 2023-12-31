// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RolePermissionInfo is the golang structure for table role_permission_info.
type RolePermissionInfo struct {
	Id           int         `json:"id"           description:""`
	RoleId       int         `json:"roleId"       description:"角色id"`
	PermissionId int         `json:"permissionId" description:"权限id"`
	CreatedAt    *gtime.Time `json:"createdAt"    description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    description:""`
}
