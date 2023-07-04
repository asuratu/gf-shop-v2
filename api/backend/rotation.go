package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RotationReq 添加轮播图
type RotationReq struct {
	g.Meta `path:"/rotations" tags:"Rotation" method:"post" summary:"添加轮播图"`
	PicUrl string `json:"pic_url" v:"required#图片不能为空" dc:"轮播图地址"`
	Link   string `json:"link" v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   uint   `json:"sort" dc:"排序"`
}
type RotationRes struct {
	RotationId uint `json:"rotation_id"`
}

// RotationDeleteReq 删除轮播图
type RotationDeleteReq struct {
	g.Meta `path:"/rotations/{id}" tags:"Rotation" method:"delete" summary:"删除轮播图"`
	Id     uint `json:"id" v:"required#请选择需要删除的轮播图" dc:"轮播图id"`
}
type RotationDeleteRes struct{}

// RotationUpdateReq 修改轮播图
type RotationUpdateReq struct {
	g.Meta `path:"/rotations/{id}" tags:"Rotation" method:"put" summary:"修改轮播图"`
	Id     uint   `json:"id" v:"required#请选择需要修改的轮播图" dc:"轮播图id"`
	PicUrl string `json:"pic_url" v:"required#图片不能为空" dc:"轮播图地址"`
	Link   string `json:"link" v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   uint   `json:"sort" dc:"排序"`
}
type RotationUpdateRes struct{}

// RotationPageListReq 轮播图分页列表
type RotationPageListReq struct {
	g.Meta `path:"/rotations" tags:"Rotation" method:"get" summary:"轮播图分页列表"`
	Sort   uint `json:"sort" in:"query" dc:"排序类型 1:升序 2:降序"`
	CommonPaginationReq
}

type RotationPageListRes struct {
	CommonPaginationRes
}
