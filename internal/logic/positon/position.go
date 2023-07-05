package position

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"

	"shop/internal/consts"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"

	"github.com/gogf/gf/v2/encoding/ghtml"
)

type sPosition struct{}

func init() {
	service.RegisterPosition(New())
}

func New() *sPosition {
	return &sPosition{}
}

// Create 添加手工位
func (s *sPosition) Create(ctx context.Context, in model.PositionCreateInput) (out model.PositionCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.PositionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.PositionCreateOutput{PositionId: uint(lastInsertID)}, err
}

// Delete 删除手工位
func (s *sPosition) Delete(ctx context.Context, id uint) error {
	return dao.PositionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除手工位
		// Unscoped() 直接删除，不走软删除
		// _, err := dao.PositionInfo.Ctx(ctx).Where(dao.PositionInfo.Columns().Id, id).Unscoped().Delete()
		// 走软删除
		_, err := dao.PositionInfo.Ctx(ctx).Where(dao.PositionInfo.Columns().Id, id).Delete()
		return err
	})
}

// Update 修改手工位
func (s *sPosition) Update(ctx context.Context, in model.PositionUpdateInput) error {
	return dao.PositionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.PositionInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.PositionInfo.Columns().Id).
			Where(dao.PositionInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 手工位列表
func (s *sPosition) GetList(ctx context.Context, in model.PositionPageListInput) (out *model.PositionPageListOutput, err error) {
	// 1. 获得*gdb.Model对象，方面后续调用
	m := dao.PositionInfo.Ctx(ctx)

	// 2. 实例化响应结构体
	out = &model.PositionPageListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	switch in.Sort {
	case consts.OrderByIdAsc:
		m = m.OrderAsc(dao.PositionInfo.Columns().Id)
	case consts.OrderByIdDesc:
		m = m.OrderDesc(dao.PositionInfo.Columns().Id)
	}

	// 3. 分页查询
	listModel := m.Page(in.Page, in.Size)

	// 4. 再查询count，判断有无数据
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}

	// 查询总条数
	if out.Total, err = m.Count(); err != nil || out.Total == 0 {
		return out, err
	}

	// 5. 延迟初始化list切片 确定有数据，再按期望大小初始化切片容量
	out.List = make([]model.PositionPageListOutputItem, 0, in.Size)

	// 6.把查询到的结果赋值到响应结构体中
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	return

}
