// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrderGoodsCommentsInfoDao is the data access object for table order_goods_comments_info.
type OrderGoodsCommentsInfoDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of current DAO.
	columns OrderGoodsCommentsInfoColumns // columns contains all the column names of Table for convenient usage.
}

// OrderGoodsCommentsInfoColumns defines and stores column names for table order_goods_comments_info.
type OrderGoodsCommentsInfoColumns struct {
	Id             string //
	OrderId        string // 订单id
	GoodsId        string // 商品id
	GoodsOptionsId string // 商品规格id
	ParentId       string // 父级评论id
	Content        string // 评论内容
	CreatedAt      string //
	UpdatedAt      string //
	DeletedAt      string //
}

// orderGoodsCommentsInfoColumns holds the columns for table order_goods_comments_info.
var orderGoodsCommentsInfoColumns = OrderGoodsCommentsInfoColumns{
	Id:             "id",
	OrderId:        "order_id",
	GoodsId:        "goods_id",
	GoodsOptionsId: "goods_options_id",
	ParentId:       "parent_id",
	Content:        "content",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewOrderGoodsCommentsInfoDao creates and returns a new DAO object for table data access.
func NewOrderGoodsCommentsInfoDao() *OrderGoodsCommentsInfoDao {
	return &OrderGoodsCommentsInfoDao{
		group:   "default",
		table:   "order_goods_comments_info",
		columns: orderGoodsCommentsInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrderGoodsCommentsInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrderGoodsCommentsInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrderGoodsCommentsInfoDao) Columns() OrderGoodsCommentsInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrderGoodsCommentsInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrderGoodsCommentsInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrderGoodsCommentsInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
