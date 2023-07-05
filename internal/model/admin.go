package model

// AdminCreateUpdateBase 创建/修改管理员基类
type AdminCreateUpdateBase struct {
	Name     string // 用户名
	Password string // 密码
	RoleIds  string // 角色ids
	UserSalt string // 加密盐
	IsAdmin  uint   // 是否超级管理员
}

// AdminCreateInput 创建管理员
type AdminCreateInput struct {
	AdminCreateUpdateBase
}
type AdminCreateOutput struct {
	AdminId uint `json:"admin_id"`
}

// AdminUpdateInput 修改管理员
type AdminUpdateInput struct {
	AdminCreateUpdateBase
	Id uint
}
type AdminUpdateOutput struct{}

type AdminShowOutput struct {
	Id      uint   `json:"id"`       // 管理员id
	Name    string `json:"name"`     // 管理员名称
	RoleIds string `json:"role_ids"` // 管理员角色
	IsAdmin uint   `json:"is_admin"` // 是否是超级管理员 1:是 0:否
}
