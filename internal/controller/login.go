package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"shop/api/backend"
	"shop/internal/model"
	"shop/internal/service"
)

// Controller 的作用是 承上启下  mvc

var Login = cLogin{}

type cLogin struct{}

func (l *cLogin) Login(ctx context.Context, req *backend.LoginDoReq) (res *backend.LoginDoRes, err error) {
	out, err := service.Login().LoginByPassword(ctx, model.AdminLoginInput{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	g.Dump("登录获取用户信息: ", out)

	res = &backend.LoginDoRes{
		Info: out,
	}

	return
}
