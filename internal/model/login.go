package model

// AdminLoginInput 管理员账号密码登录
type AdminLoginInput struct {
	Name     string // 账号
	Password string // 密码(明文)
}
