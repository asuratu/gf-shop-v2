package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// LoginDoReq 账号密码登录
type LoginDoReq struct {
	g.Meta   `path:"/login" tags:"Login" method:"post" summary:"执行登录请求"`
	Name     string `json:"name" v:"required#请输入账号" dc:"账号"`
	Password string `json:"password" v:"required#请输入密码" dc:"密码(明文)"`
}
type LoginDoRes struct {
	Info interface{} `json:"info"`
}
