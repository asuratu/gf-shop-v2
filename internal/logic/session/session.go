package session

import (
	"context"

	"shop/internal/model"
	"shop/internal/service"
)

type sSession struct{}

const (
	sessionKeyUser = "SessionKeyUser" // 用户信息存放在Session中的Key
)

func init() {
	service.RegisterSession(New())
}

func New() *sSession {
	return &sSession{}
}

// SetUser 设置用户Session.
func (s *sSession) SetUser(ctx context.Context, user *model.AdminLoginOutput) error {
	return service.BizCtx().Get(ctx).Session.Set(sessionKeyUser, user)
}

// GetUser 获取当前登录的用户信息对象，如果用户未登录返回nil。
func (s *sSession) GetUser(ctx context.Context) (user *model.AdminLoginOutput) {
	user = &model.AdminLoginOutput{}
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		v, _ := customCtx.Session.Get(sessionKeyUser)
		if !v.IsNil() {
			_ = v.Struct(&user)
		}
	}
	return
}

// RemoveUser 删除用户Session。
func (s *sSession) RemoveUser(ctx context.Context) error {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKeyUser)
	}
	return nil
}
