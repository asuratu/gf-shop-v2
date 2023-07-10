package position

import (
	"context"

	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/entity"
	"shop/internal/service"
	"shop/utility"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
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

func (s *sAdmin) GetAdminByNamePassword(ctx context.Context, in model.AdminLoginInput) map[string]interface{} {
	// 验证账号密码是否正确
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return nil
	}
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return nil
	} else {
		return g.Map{
			"id":       adminInfo.Id,
			"username": adminInfo.Name,
		}
	}
}
