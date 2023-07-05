package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AdminReq 添加管理员
type AdminReq struct {
	// tags:"xxx" 中的xxx会作为swagger的tag
	g.Meta   `path:"/admins" tags:"Admin" method:"post" summary:"添加管理员"`
	Name     string `json:"name" v:"required#管理员名称不能为空" dc:"管理员名称"`
	Password string `json:"password" v:"required#管理员密码不能为空" dc:"管理员密码"`
	RoleIds  string `json:"role_ids" v:"required#管理员角色不能为空" dc:"管理员角色"`
	IsAdmin  uint   `json:"is_admin" dc:"是否是超级管理员 1:是 0:否"`
}
type AdminRes struct {
	AdminId uint `json:"admin_id"`
}

// AdminDeleteReq 删除管理员
type AdminDeleteReq struct {
	g.Meta `path:"/admins/{id}" tags:"Admin" method:"delete" summary:"删除管理员"`
	Id     uint `json:"id" v:"required#请选择需要删除的管理员" dc:"管理员id"`
}
type AdminDeleteRes struct{}

// AdminUpdateReq 修改管理员
type AdminUpdateReq struct {
	g.Meta   `path:"/admins/{id}" tags:"Admin" method:"put" summary:"修改管理员"`
	Id       uint   `json:"id" v:"required#请选择需要修改的管理员" dc:"管理员id"`
	Name     string `json:"name" v:"required#管理员名称不能为空" dc:"管理员名称"`
	Password string `json:"password" dc:"管理员密码"`
	RoleIds  string `json:"role_ids" v:"required#管理员角色不能为空" dc:"管理员角色"`
	IsAdmin  uint   `json:"is_admin" dc:"是否是超级管理员 1:是 0:否"`
}
type AdminUpdateRes struct{}

// AdminShowReq 管理员详情
type AdminShowReq struct {
	g.Meta `path:"/admins/{id}" tags:"Admin" method:"get" summary:"管理员详情"`
	Id     uint `json:"id" v:"required#请选择需要查询的管理员" dc:"管理员id"`
}
type AdminShowRes struct {
	Id      uint   `json:"id"`       // 管理员id
	Name    string `json:"name"`     // 管理员名称
	RoleIds string `json:"role_ids"` // 管理员角色
	IsAdmin uint   `json:"is_admin"` // 是否是超级管理员 1:是 0:否
}
