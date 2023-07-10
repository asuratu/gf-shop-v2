package model

// AdminLoginInput 管理员账号密码登录
type AdminLoginInput struct {
	Name     string // 账号
	Password string // 密码(明文)
}
type AdminLoginOutput struct {
	Id      int    `json:"id"`       // ID
	Name    string `json:"name"`     // 账号
	RoleIds string `json:"role_ids"` // 角色ID列表
	IsAdmin uint8  `json:"is_admin"` // 是否是管理员
}
