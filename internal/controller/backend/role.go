package backend

import (
	"shop/api/backend"
	"shop/internal/model"
	"shop/internal/service"

	"golang.org/x/net/context"
)

// Role 角色管理
var Role = cRole{}

type cRole struct{}

func (a *cRole) Create(ctx context.Context, req *backend.RoleReq) (res *backend.RoleRes, err error) {
	out, err := service.Role().Create(ctx, model.RoleCreateInput{
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Name: req.Name,
			Desc: req.Desc,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RoleRes{RoleId: out.RoleId}, nil
}
