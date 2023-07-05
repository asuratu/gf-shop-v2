package position

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/grand"

	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/service"
	"shop/utility"

	"github.com/gogf/gf/v2/encoding/ghtml"
)

type sAdmin struct{}

func init() {
	service.RegisterAdmin(New())
}

func New() *sAdmin {
	return &sAdmin{}
}

// Create 添加管理员
func (s *sAdmin) Create(ctx context.Context, in model.AdminCreateInput) (out model.AdminCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}

	// TODO name 唯一性校验

	// 处理加密盐和密码的逻辑
	UserSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt

	lastInsertID, err := dao.AdminInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.AdminCreateOutput{AdminId: uint(lastInsertID)}, err
}

// Delete 删除管理员
func (s *sAdmin) Delete(ctx context.Context, id uint) error {
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除管理员
		_, err := dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Id, id).Delete()
		return err
	})
}

// Update 修改管理员
func (s *sAdmin) Update(ctx context.Context, in model.AdminUpdateInput) error {
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}

		// TODO name 唯一性校验

		// 处理加密盐和密码的逻辑
		if in.Password != "" {
			UserSalt := grand.S(10)
			in.Password = utility.EncryptPassword(in.Password, UserSalt)
			in.UserSalt = UserSalt
		}

		_, err := dao.AdminInfo.
			Ctx(ctx).
			Data(in).
			OmitEmpty(). // 忽略空值
			Where(dao.AdminInfo.Columns().Id, in.Id).
			Update()

		return err
	})
}

// Show 获取管理员详情
func (s *sAdmin) Show(ctx context.Context, id uint) (out model.AdminShowOutput, err error) {
	err = dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Id, id).Scan(&out)
	return
}
