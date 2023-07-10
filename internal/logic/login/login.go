package login

import (
	"context"

	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/entity"
	"shop/internal/service"
	"shop/utility"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(New())
}

func New() *sLogin {
	return &sLogin{}
}

// LoginByPassword Login 执行登录
func (s *sLogin) LoginByPassword(ctx context.Context, in model.AdminLoginInput) (out *model.AdminLoginOutput, err error) {
	// 根据用户名获取管理员信息
	adminInfo := entity.AdminInfo{}
	err = dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, in.Name).Scan(&adminInfo)
	if err != nil {
		return nil, err
	}
	// 验证密码是否正确
	if adminInfo.Password != utility.EncryptPassword(in.Password, adminInfo.UserSalt) {
		return nil, gerror.New("密码错误")
	}

	user := &model.AdminLoginOutput{
		Id:      adminInfo.Id,
		Name:    adminInfo.Name,
		RoleIds: adminInfo.RoleIds,
		IsAdmin: uint8(adminInfo.IsAdmin),
	}

	// 设置session
	if err := service.Session().SetUser(ctx, user); err != nil {
		return nil, err
	}

	// 自动更新上下文
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:      uint(adminInfo.Id),
		Name:    adminInfo.Name,
		IsAdmin: uint8(adminInfo.IsAdmin),
	})

	return user, nil
}
