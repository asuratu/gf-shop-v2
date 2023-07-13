package backend

import (
	"github.com/gogf/gf/v2/frame/g"

	"shop/internal/model"
)

type RoleCreateUpdateBase struct {
	Name string `json:"name" v:"required#名称必填" dc:"角色名称"`
	Desc string `json:"desc" dc:"角色描述"`
}

type BaseRoleRes struct {
	Id   uint   `json:"id" desc:"角色id"`
	Name string `json:"name" desc:"角色名称"`
	Desc string `json:"desc" desc:"角色描述"`
}

type RoleReq struct {
	g.Meta `path:"/roles" method:"post" desc:"添加角色" tags:"role"`
	RoleCreateUpdateBase
}
type RoleRes struct {
	RoleId uint `json:"role_id"`
}

type RoleUpdateReq struct {
	g.Meta `path:"/roles/{id}" method:"put" desc:"修改角色" tags:"role"`
	Id     uint `json:"id" v:"required#id必填" desc:"id"`
	RoleCreateUpdateBase
}
type RoleUpdateRes struct {
	RoleId uint `json:"role_id"`
}

type RoleDeleteReq struct {
	g.Meta `path:"/roles/{id}" method:"delete" tags:"角色" summary:"删除角色接口"`
	Id     uint `v:"min:1#请选择需要删除的角色" dc:"角色id"`
}
type RoleDeleteRes struct{}

type RoleGetListCommonReq struct {
	g.Meta `path:"/roles" method:"get" tags:"角色" summary:"角色列表接口"`
	CommonPaginationReq
}
type RoleGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type AddRolePermissionReq struct {
	g.Meta        `path:"/roles/{role_id}/permissions" method:"post" tags:"角色" summary:"给角色添加权限"`
	RoleId        uint   `json:"role_id" desc:"角色id"`
	PermissionIds []uint `json:"permission_ids" desc:"权限id数组"`
}
type AddRolePermissionRes struct{}

type DeletePermissionReq struct {
	g.Meta        `path:"/roles/{role_id}/permissions" method:"delete" tags:"角色" summary:"角色删除权限接口"`
	RoleId        uint   `json:"role_id" desc:"角色id"`
	PermissionIds []uint `json:"permission_ids" desc:"权限id数组"`
}
type DeletePermissionRes struct{}

type RoleGetPermissionReq struct {
	g.Meta `path:"/roles/{role_id}/permissions" method:"get" tags:"角色" summary:"获取角色权限接口"`
	RoleId uint `json:"role_id" desc:"角色id"`
}
type RoleGetPermissionRes struct {
	Role               *model.BaseRole               `json:"role" desc:"角色信息"`
	RolePermissionList []*model.RolePermissionEntity `json:"permission_list" desc:"权限列表"`
}
