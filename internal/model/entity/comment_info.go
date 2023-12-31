// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CommentInfo is the golang structure for table comment_info.
type CommentInfo struct {
	Id        int         `json:"id"        description:""`
	ParentId  int         `json:"parentId"  description:"父级评论id"`
	UserId    int         `json:"userId"    description:""`
	ObjectId  int         `json:"objectId"  description:""`
	Type      int         `json:"type"      description:"评论类型：1商品 2文章"`
	Content   string      `json:"content"   description:"评论内容"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
}
