package data

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
	"shop/utility"

	"github.com/gogf/gf/v2/os/gtime"
)

type sData struct {
}

// 注册服务
func init() {
	service.RegisterData(New())
}

func New() *sData {
	return &sData{}
}

func (s *sData) DataHead(ctx context.Context) (out *model.DataHeadOutput, err error) {
	return &model.DataHeadOutput{
		TodayOrderCount: todayOrderCount(ctx),
		DAU:             utility.RandInt(1000),
		ConversionRate:  utility.RandInt(50),
	}, nil
}

// 查询今天的订单总数
func todayOrderCount(ctx context.Context) (count int) {
	g.Dump(gtime.Now().StartOfDay(), gtime.Now().EndOfDay())
	count, err := dao.OrderInfo.Ctx(ctx).
		WhereBetween(dao.OrderInfo.Columns().CreatedAt, gtime.Now().StartOfDay(), gtime.Now().EndOfDay()).
		Count(dao.OrderInfo.Columns().Id)
	if err != nil {
		count = -1
	}
	return
}
