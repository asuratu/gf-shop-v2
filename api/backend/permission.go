package backend

type BasePermission struct {
	Id   uint   `json:"id" desc:"权限id"`
	Name string `json:"name" desc:"权限名称"`
	Path string `json:"path" desc:"权限路径"`
}
