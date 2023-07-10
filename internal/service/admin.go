// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"shop/internal/model"
)

type (
	IAdmin interface {
		Create(ctx context.Context, in model.AdminCreateInput) (out model.AdminCreateOutput, err error)
		Delete(ctx context.Context, id uint) error
		Update(ctx context.Context, in model.AdminUpdateInput) error
		Show(ctx context.Context, id uint) (out model.AdminShowOutput, err error)
		GetAdminByNamePassword(ctx context.Context, in model.AdminLoginInput) map[string]interface{}
	}
)

var (
	localAdmin IAdmin
)

func Admin() IAdmin {
	if localAdmin == nil {
		panic("implement not found for interface IAdmin, forgot register?")
	}
	return localAdmin
}

func RegisterAdmin(i IAdmin) {
	localAdmin = i
}
