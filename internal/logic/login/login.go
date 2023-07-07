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

// Login 添加管理员
func (s *sLogin) Login(ctx context.Context, in model.AdminLoginInput) (err error) {
	// 根据用户名获取管理员信息
	adminInfo := entity.AdminInfo{}
	err = dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, in.Name).Scan(&adminInfo)
	if err != nil {
		return err
	}
	// 验证密码是否正确
	if adminInfo.Password != utility.EncryptPassword(in.Password, adminInfo.UserSalt) {
		return gerror.New("密码错误")
	}

	if err := service.Session().SetUser(ctx, &adminInfo); err != nil {
		return err
	}
	// 自动更新上线 for session
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:      uint(adminInfo.Id),
		Name:    adminInfo.Name,
		IsAdmin: uint8(adminInfo.IsAdmin),
	})
	return nil

}
