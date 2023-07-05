package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PositionCreateUpdateBase 创建/修改手工位基类
type PositionCreateUpdateBase struct {
	PicUrl    string // 图片地址
	Link      string // 跳转链接
	GoodsName string // 商品名称
	GoodsId   string // 商品ID
	Sort      uint   // 排序
}

// PositionCreateInput 创建手工位
type PositionCreateInput struct {
	PositionCreateUpdateBase
}

// PositionCreateOutput 创建手工位返回结果
type PositionCreateOutput struct {
	PositionId uint `json:"position_id"`
}

// PositionUpdateInput 修改手工位
type PositionUpdateInput struct {
	PositionCreateUpdateBase
	Id uint
}

// PositionPageListInput 手工位分页列表
type PositionPageListInput struct {
	Sort uint
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// PositionPageListOutput 手工位分页列表结果
type PositionPageListOutput struct {
	List  []PositionPageListOutputItem `json:"list" description:"列表"`
	Page  int                          `json:"page" description:"分页码"`
	Size  int                          `json:"size" description:"分页数量"`
	Total int                          `json:"total" description:"数据总数"`
}

// PositionPageListOutputItem 手工位分页列表结果单条
type PositionPageListOutputItem struct {
	Id        uint        `json:"id"` // 自增ID
	PicUrl    string      `json:"pic_url"`
	Link      string      `json:"link"`
	GoodsName string      `json:"goods_name"`
	GoodsId   string      `json:"goods_id"`
	Sort      uint        `json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
