// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderGoodsCommentsInfo is the golang structure for table order_goods_comments_info.
type OrderGoodsCommentsInfo struct {
	Id             int         `json:"id"             description:""`
	OrderId        int         `json:"orderId"        description:"订单id"`
	GoodsId        int         `json:"goodsId"        description:"商品id"`
	GoodsOptionsId int         `json:"goodsOptionsId" description:"商品规格id"`
	ParentId       int         `json:"parentId"       description:"父级评论id"`
	Content        string      `json:"content"        description:"评论内容"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:""`
	DeletedAt      *gtime.Time `json:"deletedAt"      description:""`
}
