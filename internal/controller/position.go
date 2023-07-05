package controller

import (
	"context"

	"shop/api/backend"
	"shop/internal/model"
	"shop/internal/service"
)

// Controller 的作用是 承上启下  mvc

// Position 内容管理
var Position = cPosition{}

type cPosition struct{}

func (a *cPosition) Create(ctx context.Context, req *backend.PositionReq) (res *backend.PositionRes, err error) {
	out, err := service.Position().Create(ctx, model.PositionCreateInput{
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.PicUrl,
			Link:      req.Link,
			GoodsName: req.GoodsName,
			GoodsId:   req.GoodsId,
			Sort:      req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.PositionRes{PositionId: out.PositionId}, nil
}

func (a *cPosition) Delete(ctx context.Context, req *backend.PositionDeleteReq) (res *backend.PositionDeleteRes, err error) {
	err = service.Position().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &backend.PositionDeleteRes{}, nil
}

func (a *cPosition) Update(ctx context.Context, req *backend.PositionUpdateReq) (res *backend.PositionUpdateRes, err error) {
	err = service.Position().Update(ctx, model.PositionUpdateInput{
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.PicUrl,
			Link:      req.Link,
			GoodsName: req.GoodsName,
			GoodsId:   req.GoodsId,
			Sort:      req.Sort,
		},
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &backend.PositionUpdateRes{}, nil
}

// Index 手工位列表
func (a *cPosition) Index(ctx context.Context, req *backend.PositionPageListReq) (res *backend.PositionPageListRes, err error) {
	list, err := service.Position().GetList(ctx, model.PositionPageListInput{
		Sort: req.Sort,
		Page: req.Page,
		Size: req.Size,
	})

	if err != nil {
		return nil, err
	}
	return &backend.PositionPageListRes{
		CommonPaginationRes: backend.CommonPaginationRes{
			List:  list.List,
			Total: list.Total,
			Page:  list.Page,
			Size:  list.Size,
		},
	}, nil

}
