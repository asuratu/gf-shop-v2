package frontend

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

// Index
//
//	@Description: 轮播图列表
//	@receiver a
//	@param ctx
//	@param req
//	@return res
//	@return err
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
