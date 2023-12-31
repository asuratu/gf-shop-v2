// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ArticleInfo is the golang structure for table article_info.
type ArticleInfo struct {
	Id        int         `json:"id"        description:""`
	UserId    int         `json:"userId"    description:"作者id"`
	Title     string      `json:"title"     description:"标题"`
	Desc      string      `json:"desc"      description:"摘要"`
	PicUrl    string      `json:"picUrl"    description:"封面图"`
	IsAdmin   int         `json:"isAdmin"   description:"1后台管理员发布 2前台用户发布"`
	Praise    int         `json:"praise"    description:"点赞数"`
	Detail    string      `json:"detail"    description:"文章详情"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
}
