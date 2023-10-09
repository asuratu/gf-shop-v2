package cmd

import (
	"context"
	"os"
	"strconv"

	apiBackend "shop/api/backend"
	"shop/boot/job"
	"shop/internal/consts"
	"shop/internal/controller/backend"
	"shop/internal/controller/frontend"
	"shop/internal/dao"
	"shop/internal/model/entity"
	"shop/internal/service"
	"shop/utility"
	"shop/utility/response"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 初始化 AsynqServer
			job.AsynqServer = job.NewAsynqServer(ctx)

			// 初始化 AsynqClient
			job.AsynqClient = job.NewAsynqClient(ctx)

			cronJob := job.NewCronJob(ctx)
			mux := cronJob.Register()

			go func() {
				if err := job.AsynqServer.Run(mux); err != nil {
					os.Exit(1)
				}
			}()

			// 优雅重启、关闭
			// 访问: http://127.0.0.1:8199/debug/admin
			s.EnableAdmin()

			// 启动gtoken
			gfAdminToken := &gtoken.GfToken{
				CacheMode:       2,                          // 缓存模式 1:内存 gcache 2:gredis 3:fileCache 默认1
				EncryptKey:      []byte("1234567890123456"), // 加密key
				ServerName:      "shop",                     // 服务名称
				LoginPath:       "/login",                   // 登录路径
				LoginBeforeFunc: LoginBeforeFunc,            // 登录验证方法
				LoginAfterFunc:  loginAfterFunc,             // 登录返回方法
				LogoutPath:      "/logout",                  // 退出路径
				AuthPaths:       g.SliceStr{"/admins/info"}, //
				// AuthExcludePaths: g.SliceStr{"/admin/user/info", "/admin/system/user/info"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
				AuthAfterFunc: authAfterFunc, // 拦截认证前后调用
				MultiLogin:    true,          // 是否支持多点登录
			}

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
					backend.Data,     // 数据大屏
					// jwt 登录
					// backend.Login,    // 管理员
				)
				// 需要登录的路由组绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					// 使用 jwt 中间件
					// group.Middleware(service.Middleware().Auth)
					// 使用 gtoken 中间件
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Bind(
						backend.Admin, // 管理员
						backend.Role,  // 角色
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

// TODO 迁移到合适的位置
func LoginBeforeFunc(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()
	ctx := context.TODO()

	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}

	// 验证账号密码是否正确
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", name).Scan(&adminInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}
	if utility.EncryptPassword(password, adminInfo.UserSalt) != adminInfo.Password {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}
	// 唯一标识，扩展参数user data
	return consts.GTokenAdminPrefix + strconv.Itoa(adminInfo.Id), adminInfo
}

// TODO 迁移到合适的位置
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		// 获得登录用户id
		userKey := respData.GetString("userKey")
		adminId := gstr.StrEx(userKey, consts.GTokenAdminPrefix)
		g.Dump("adminId:", adminId)
		// 根据id获得登录用户其他信息
		adminInfo := entity.AdminInfo{}
		err := dao.AdminInfo.Ctx(context.TODO()).WherePri(adminId).Scan(&adminInfo)
		if err != nil {
			return
		}
		// 通过角色查询权限
		// 先通过角色查询权限id
		var rolePermissionInfos []entity.RolePermissionInfo
		err = dao.RolePermissionInfo.Ctx(context.TODO()).WhereIn(dao.RolePermissionInfo.Columns().RoleId, g.Slice{adminInfo.RoleIds}).Scan(&rolePermissionInfos)
		if err != nil {
			return
		}
		permissionIds := g.Slice{}
		for _, info := range rolePermissionInfos {
			permissionIds = append(permissionIds, info.PermissionId)
		}

		var permissions []entity.PermissionInfo
		err = dao.PermissionInfo.Ctx(context.TODO()).WhereIn(dao.PermissionInfo.Columns().Id, permissionIds).Scan(&permissions)
		if err != nil {
			return
		}
		data := &apiBackend.LoginRes{
			Type:        "Bearer",
			Token:       respData.GetString("token"),
			ExpireIn:    10 * 24 * 60 * 60, // 单位秒,
			IsAdmin:     adminInfo.IsAdmin,
			RoleIds:     adminInfo.RoleIds,
			Permissions: permissions,
		}
		response.JsonExit(r, 0, "", data)
	}
}

// TODO 迁移到合适的位置
func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	var adminInfo entity.AdminInfo
	err := gconv.Struct(respData.GetString("data"), &adminInfo)
	// 认证失败
	if err != nil {
		response.Auth(r)
		return
	}
	// 账号被冻结拉黑
	if adminInfo.DeletedAt != nil {
		response.AuthBlack(r)
		return
	}

	// 存入上下文, 以便后续使用
	r.SetCtxVar(consts.CtxAdminId, adminInfo.Id)
	r.SetCtxVar(consts.CtxAdminName, adminInfo.Name)
	r.SetCtxVar(consts.CtxAdminIsAdmin, adminInfo.IsAdmin)
	r.SetCtxVar(consts.CtxAdminRoleIds, adminInfo.RoleIds)
	r.Middleware.Next()
}
