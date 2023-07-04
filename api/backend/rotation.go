package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RotationReq struct {
	g.Meta `path:"/rotations" tags:"Rotation" method:"post" summary:"添加轮播图"`
	PicUrl string `json:"pic_url" v:"required#图片不能为空" dc:"轮播图地址"`
	Link   string `json:"link" v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   uint   `json:"sort" dc:"排序"`
}
type RotationRes struct {
	RotationId uint `json:"rotation_id"`
}
