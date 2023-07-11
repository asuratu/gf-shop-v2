package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RotationPageListReq 轮播图分页列表
type RotationPageListReq struct {
	g.Meta `path:"/rotations" tags:"Rotation" method:"get" summary:"轮播图分页列表"`
	Sort   uint `json:"sort" in:"query" dc:"排序类型 1:升序 2:降序"`
	CommonPaginationReq
}
type RotationPageListRes struct {
	CommonPaginationRes
}
