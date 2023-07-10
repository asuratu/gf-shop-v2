package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"shop/internal/controller"
	"shop/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/backend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				// 不需要登录的路由组绑定
				group.Bind(
				// controller.Admin.Create, // 管理员
				// controller.Login,        // 登录
				)
				// 需要登录的路由组绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					// err := gfAdminToken.Middleware(ctx, group)
					// if err != nil {
					// 	panic(err)
					// }
					group.Bind(
						controller.Rotation, // 轮播图
						controller.Position, // 手工位
						controller.Admin,    // 管理员
						controller.Login,    // 管理员
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
