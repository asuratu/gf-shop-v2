package controller

import (
	"context"

	"shop/api/backend"
	"shop/internal/model"
	"shop/internal/service"
)

// Controller 的作用是 承上启下  mvc

// Admin 内容管理
var Admin = cAdmin{}

type cAdmin struct{}

func (a *cAdmin) Create(ctx context.Context, req *backend.AdminReq) (res *backend.AdminRes, err error) {
	out, err := service.Admin().Create(ctx, model.AdminCreateInput{
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			Name:     req.Name,
			Password: req.Password,
			RoleIds:  req.RoleIds,
			IsAdmin:  req.IsAdmin,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.AdminRes{AdminId: out.AdminId}, nil
}

func (a *cAdmin) Delete(ctx context.Context, req *backend.AdminDeleteReq) (res *backend.AdminDeleteRes, err error) {
	err = service.Admin().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &backend.AdminDeleteRes{}, nil
}

func (a *cAdmin) Update(ctx context.Context, req *backend.AdminUpdateReq) (res *backend.AdminUpdateRes, err error) {
	err = service.Admin().Update(ctx, model.AdminUpdateInput{
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			Name:     req.Name,
			Password: req.Password,
			RoleIds:  req.RoleIds,
			IsAdmin:  req.IsAdmin,
		},
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &backend.AdminUpdateRes{}, nil
}

// Show 管理员详情
func (a *cAdmin) Show(ctx context.Context, req *backend.AdminShowReq) (res *backend.AdminShowRes, err error) {
	out, err := service.Admin().Show(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &backend.AdminShowRes{
		Id:      out.Id,
		Name:    out.Name,
		RoleIds: out.RoleIds,
		IsAdmin: out.IsAdmin,
	}, nil
}
