package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"shop/internal/controller/backend"
	"shop/internal/controller/frontend"
	"shop/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 优雅重启、关闭
			// 访问: http://127.0.0.1:8199/debug/admin
			s.EnableAdmin()

			s.Group("/backend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				// 不需要登录的路由组绑定
				group.Bind(
					backend.Rotation, // 轮播图
					backend.Position, // 手工位
					backend.Login,    // 管理员
				)
				// 需要登录的路由组绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Bind(
						backend.Admin, // 管理员
					)
				})
			})

			s.Group("/frontend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				// 不需要登录的路由组绑定
				group.Bind(
					frontend.Rotation, // 轮播图
				)
			})
			s.Run()
			return nil
		},
	}
)
