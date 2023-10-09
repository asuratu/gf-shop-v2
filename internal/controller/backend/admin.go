package backend

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"shop/api/backend"
	"shop/internal/consts"
	"shop/internal/model"
	"shop/internal/service"
)

// Controller 的作用是 承上启下  mvc

// Admin 内容管理
var Admin = cAdmin{}

type cAdmin struct{}

// Create
//
//	@Description:
//	@receiver a
//	@param ctx
//	@param req
//	@return res
//	@return err
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
	g.Dump("req.Id", req.AdminId)
	out, err := service.Admin().Show(ctx, req.AdminId)
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

// Me 管理员详情 for jwt
// func (a *cAdmin) Me(ctx context.Context, req *backend.AdminInfoReq) (res *backend.AdminInfoRes, err error) {
// 	return &backend.AdminInfoRes{
// 		Id:          gconv.Int(service.Auth().GetIdentity(ctx)),
// 		IdentityKey: service.Auth().IdentityKey,
// 		Payload:     service.Auth().GetPayload(ctx),
// 	}, nil
// }

// Me 管理员详情 for gtoken
func (a *cAdmin) Me(ctx context.Context, req *backend.AdminInfoReq) (res *backend.AdminShowRes, err error) {
	return &backend.AdminShowRes{
		Id:      gconv.Uint(ctx.Value(consts.CtxAdminId)),
		Name:    gconv.String(ctx.Value(consts.CtxAdminName)),
		RoleIds: gconv.String(ctx.Value(consts.CtxAdminRoleIds)),
		IsAdmin: gconv.Uint(ctx.Value(consts.CtxAdminIsAdmin)),
	}, nil
}
