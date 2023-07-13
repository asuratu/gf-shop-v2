package model

type BasePermission struct {
	Id   uint   `orm:"id,primary" json:"id"`    // 自增ID
	Name string `orm:"name,unique" json:"name"` // 权限名称
	Path string `orm:"path" json:"path"`        // 权限路径
}
