// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserCouponInfoDao is the data access object for table user_coupon_info.
type UserCouponInfoDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns UserCouponInfoColumns // columns contains all the column names of Table for convenient usage.
}

// UserCouponInfoColumns defines and stores column names for table user_coupon_info.
type UserCouponInfoColumns struct {
	Id        string // 用户优惠券表
	UserId    string //
	CouponId  string //
	Status    string // 状态：1可用 2已用 3过期
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// userCouponInfoColumns holds the columns for table user_coupon_info.
var userCouponInfoColumns = UserCouponInfoColumns{
	Id:        "id",
	UserId:    "user_id",
	CouponId:  "coupon_id",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewUserCouponInfoDao creates and returns a new DAO object for table data access.
func NewUserCouponInfoDao() *UserCouponInfoDao {
	return &UserCouponInfoDao{
		group:   "default",
		table:   "user_coupon_info",
		columns: userCouponInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserCouponInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserCouponInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserCouponInfoDao) Columns() UserCouponInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserCouponInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserCouponInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserCouponInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
