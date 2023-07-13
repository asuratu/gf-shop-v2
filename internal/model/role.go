package model

import "github.com/gogf/gf/v2/os/gtime"

// RoleCreateUpdateBase 创建/修改内容基类
type RoleCreateUpdateBase struct {
	Name string
	Desc string
}

type BaseRole struct {
	Id   uint   `orm:"id,primary" json:"id"`    // 自增ID
	Name string `orm:"name,unique" json:"name"` // 角色名称
	Desc string `orm:"desc" json:"desc"`        // 角色描述
}

// RoleCreateInput 创建内容
type RoleCreateInput struct {
	RoleCreateUpdateBase
}

// RoleCreateUpdateOutput  创建和更新内容返回结果
type RoleCreateUpdateOutput struct {
	RoleId uint `json:"role_id"`
}

// RoleUpdateInput 修改内容
type RoleUpdateInput struct {
	RoleCreateUpdateBase
	Id uint
}

// RoleGetListInput 获取内容列表
type RoleGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// RoleGetListOutput 查询列表结果
type RoleGetListOutput struct {
	List  []RoleGetListOutputItem `json:"list" description:"列表"`
	Page  int                     `json:"page" description:"分页码"`
	Size  int                     `json:"size" description:"分页数量"`
	Total int                     `json:"total" description:"数据总数"`
}

type RoleGetListOutputItem struct {
	Id        uint        `json:"id"` // 自增ID
	Name      string      `json:"name"`
	Desc      string      `json:"desc"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type RoleSearchOutputItem struct {
	RoleGetListOutputItem
}

type RoleAddPermissionInput struct {
	RoleId        uint   `json:"role_id"`
	PermissionIds []uint `json:"permission_ids"`
}

type RoleDeletePermissionInput struct {
	RoleId        uint   `json:"role_id"`
	PermissionIds []uint `json:"permission_ids"`
}

// RoleGetPermissionListInput 组合模型
type RoleGetPermissionListInput struct {
	RoleId uint `json:"role_id"`
}

//	type RoleGetPermissionListOutput struct {
//		Role               *BaseRole
//		RolePermissionList BaseRolePermissionList
//	}
type RoleGetPermissionListOutput struct {
	Role               *BaseRole
	RolePermissionList []*RolePermissionEntity
}
