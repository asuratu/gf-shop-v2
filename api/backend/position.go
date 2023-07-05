package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PositionCreateUpdateBase struct {
	PicUrl    string `json:"pic_url" v:"required#图片不能为空" dc:"手工位地址"`
	Link      string `json:"link" v:"required#跳转链接不能为空" dc:"跳转链接"`
	GoodsName string `json:"goods_name" v:"required#商品名称不能为空" dc:"商品名称"`
	GoodsId   string `json:"goods_id" v:"required#商品ID不能为空" dc:"商品ID"`
	Sort      uint   `json:"sort" dc:"排序"`
}

// PositionReq 添加手工位
type PositionReq struct {
	g.Meta `path:"/positions" tags:"Position" method:"post" summary:"添加手工位"`
	PositionCreateUpdateBase
}
type PositionRes struct {
	PositionId uint `json:"position_id"`
}

// PositionDeleteReq 删除手工位
type PositionDeleteReq struct {
	g.Meta `path:"/positions/{id}" tags:"Position" method:"delete" summary:"删除手工位"`
	Id     uint `json:"id" v:"required#请选择需要删除的手工位" dc:"手工位id"`
}
type PositionDeleteRes struct{}

// PositionUpdateReq 修改手工位
type PositionUpdateReq struct {
	g.Meta `path:"/positions/{id}" tags:"Position" method:"put" summary:"修改手工位"`
	Id     uint `json:"id" v:"required#请选择需要修改的手工位" dc:"手工位id"`
	PositionCreateUpdateBase
}
type PositionUpdateRes struct{}

// PositionPageListReq 手工位分页列表
type PositionPageListReq struct {
	g.Meta `path:"/positions" tags:"Position" method:"get" summary:"手工位分页列表"`
	Sort   uint `json:"sort" in:"query" dc:"排序类型 1:升序 2:降序"`
	CommonPaginationReq
}

type PositionPageListRes struct {
	CommonPaginationRes
}
