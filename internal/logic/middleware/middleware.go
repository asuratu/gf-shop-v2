package middleware

import (
	"shop/internal/model"
	"shop/internal/service"
	"shop/utility/response"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sMiddleware struct {
	LoginUrl string // 登录路由地址
}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{
		LoginUrl: "/backend/login",
	}
}

type TokenInfo struct {
	Id   int
	Name string
}

func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()
	// 如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		return
	}
	var (
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)

	if err != nil {
		if bizCode, ok := code.(response.BizCode); ok {
			response.JsonExit(r, bizCode.BizDetail().HttpCode, bizCode.Message(), res)
		}
		response.JsonExit(r, 500, err.Error(), res)
	}

	response.JsonExit(r, 0, "OK", res)
}

// Ctx 自定义上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: r.Session,
		Data:    make(g.Map),
	}
	service.BizCtx().Init(r, customCtx)
	if userEntity := service.Session().GetUser(r.Context()); userEntity.Id > 0 {
		customCtx.User = &model.ContextUser{
			Id:      uint(userEntity.Id),
			Name:    userEntity.Name,
			IsAdmin: userEntity.IsAdmin,
		}
	}
	// 将自定义的上下文对象传递到模板变量中使用
	r.Assigns(g.Map{
		"Context": customCtx,
	})
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *sMiddleware) Auth(r *ghttp.Request) {
	service.Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}

// var GToken *gtoken.GfToken

// GTokenSetCtx Gtoken鉴权
/*func (s *sMiddleware) GTokenSetCtx(r *ghttp.Request) {
	var tokenInfo TokenInfo
	token := GToken.GetTokenData(r)
	err := gconv.Struct(token.GetString("data"), &tokenInfo)
	if err != nil {
		response.Auth(r)
		return
	}

	// 账号被冻结拉黑
	// if tokenInfo.Status == 2 {
	//	response.AuthBlack(r)
	//	return
	// }

	r.SetCtxVar(CtxAccountId, tokenInfo.Id)
	r.SetCtxVar(CtxAccountName, tokenInfo.Name)
	r.Middleware.Next()
}*/
