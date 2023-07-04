package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RotationCreateUpdateBase 创建/修改轮播图基类
type RotationCreateUpdateBase struct {
	PicUrl string // 图片地址
	Link   string // 跳转链接
	Sort   uint   // 排序
}

// RotationCreateInput 创建轮播图
type RotationCreateInput struct {
	RotationCreateUpdateBase
}

// RotationCreateOutput 创建轮播图返回结果
type RotationCreateOutput struct {
	RotationId uint `json:"rotation_id"`
}

// RotationUpdateInput 修改轮播图
type RotationUpdateInput struct {
	RotationCreateUpdateBase
	Id uint
}

// RotationPageListInput 轮播图分页列表
type RotationPageListInput struct {
	Sort uint
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// RotationPageListOutput 轮播图分页列表结果
type RotationPageListOutput struct {
	List  []RotationPageListOutputItem `json:"list" description:"列表"`
	Page  int                          `json:"page" description:"分页码"`
	Size  int                          `json:"size" description:"分页数量"`
	Total int                          `json:"total" description:"数据总数"`
}

// RotationPageListOutputItem 轮播图分页列表结果单条
type RotationPageListOutputItem struct {
	Id        uint        `json:"id"` // 自增ID
	PicUrl    string      `json:"pic_url"`
	Link      string      `json:"link"`
	Sort      uint        `json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
