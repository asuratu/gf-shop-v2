package rotation

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"

	"shop/internal/consts"
	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"

	"github.com/gogf/gf/v2/encoding/ghtml"
)

type sRotation struct{}

func init() {
	service.RegisterRotation(New())
}

func New() *sRotation {
	return &sRotation{}
}

// Create 添加轮播图
func (s *sRotation) Create(ctx context.Context, in model.RotationCreateInput) (out model.RotationCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.RotationInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RotationCreateOutput{RotationId: uint(lastInsertID)}, err
}

// Delete 删除轮播图
func (s *sRotation) Delete(ctx context.Context, id uint) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除轮播图
		// Unscoped() 直接删除，不走软删除
		// _, err := dao.RotationInfo.Ctx(ctx).Where(dao.RotationInfo.Columns().Id, id).Unscoped().Delete()
		// 走软删除
		_, err := dao.RotationInfo.Ctx(ctx).Where(dao.RotationInfo.Columns().Id, id).Delete()
		return err
	})
}

// Update 修改轮播图
func (s *sRotation) Update(ctx context.Context, in model.RotationUpdateInput) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.RotationInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.RotationInfo.Columns().Id).
			Where(dao.RotationInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 轮播图列表
func (s *sRotation) GetList(ctx context.Context, in model.RotationPageListInput) (out *model.RotationPageListOutput, err error) {
	// 1. 获得*gdb.Model对象，方面后续调用
	m := dao.RotationInfo.Ctx(ctx)

	// 2. 实例化响应结构体
	out = &model.RotationPageListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	switch in.Sort {
	case consts.OrderByIdAsc:
		m = m.OrderAsc(dao.RotationInfo.Columns().Id)
	case consts.OrderByIdDesc:
		m = m.OrderDesc(dao.RotationInfo.Columns().Id)
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
	out.List = make([]model.RotationPageListOutputItem, 0, in.Size)

	// 6.把查询到的结果赋值到响应结构体中
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	return

}
