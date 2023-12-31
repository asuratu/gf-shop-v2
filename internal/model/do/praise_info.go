// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PraiseInfo is the golang structure of table praise_info for DAO operations like Where/Data.
type PraiseInfo struct {
	g.Meta    `orm:"table:praise_info, do:true"`
	Id        interface{} // 点赞表
	UserId    interface{} //
	Type      interface{} // 点赞类型 1商品 2文章
	ObjectId  interface{} // 点赞对象id 方便后期扩展
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
