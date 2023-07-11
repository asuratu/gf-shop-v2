package backend

import (
	"context"

	"shop/api/backend"
	"shop/internal/model"
	"shop/internal/service"
)

// Controller 的作用是 承上启下  mvc

// Rotation 内容管理
var Rotation = cRotation{}

type cRotation struct{}

func (a *cRotation) Create(ctx context.Context, req *backend.RotationReq) (res *backend.RotationRes, err error) {
	out, err := service.Rotation().Create(ctx, model.RotationCreateInput{
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RotationRes{RotationId: out.RotationId}, nil
}

func (a *cRotation) Delete(ctx context.Context, req *backend.RotationDeleteReq) (res *backend.RotationDeleteRes, err error) {
	err = service.Rotation().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &backend.RotationDeleteRes{}, nil
}

func (a *cRotation) Update(ctx context.Context, req *backend.RotationUpdateReq) (res *backend.RotationUpdateRes, err error) {
	err = service.Rotation().Update(ctx, model.RotationUpdateInput{
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &backend.RotationUpdateRes{}, nil
}

// Index 轮播图列表
func (a *cRotation) Index(ctx context.Context, req *backend.RotationPageListReq) (res *backend.RotationPageListRes, err error) {
	list, err := service.Rotation().GetList(ctx, model.RotationPageListInput{
		Sort: req.Sort,
		Page: req.Page,
		Size: req.Size,
	})

	if err != nil {
		return nil, err
	}
	return &backend.RotationPageListRes{
		CommonPaginationRes: backend.CommonPaginationRes{
			List:  list.List,
			Total: list.Total,
			Page:  list.Page,
			Size:  list.Size,
		},
	}, nil

}
