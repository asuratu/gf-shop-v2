// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PositionInfo is the golang structure for table position_info.
type PositionInfo struct {
	Id        int         `json:"id"        description:""`
	PicUrl    string      `json:"picUrl"    description:"图片链接"`
	GoodsName string      `json:"goodsName" description:"商品名称"`
	Link      string      `json:"link"      description:"跳转链接"`
	Sort      int         `json:"sort"      description:"排序"`
	GoodsId   int         `json:"goodsId"   description:"商品id"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
}
