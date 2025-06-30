package engine

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kshdb/ks_gf_utility/utility/tool"
	"github.com/kshdb/ks_gf_utility/utility/tool/auth"
)

var (
	//全局上下文
	Ctx g.Ctx = gctx.New()
	//接口端口
	ServerAddress = g.Cfg().MustGet(Ctx, "server.address").String()
	//接口基础地址 前后不带斜杠
	BaseApiPath = g.Cfg().MustGet(Ctx, "systemRun.baseApiPath").String()
)

func init() {
	if BaseApiPath == "" {
		BaseApiPath = "prod-api"
	}
}

// 跨域中间件
func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// 上下文
func CtxInfo(r *ghttp.Request) {
	_ctx := r.Context()
	// 初始化，务必最开始执行
	customCtx := &tool.UserContext{
		OtherInfo: g.Map{"run_type": g.Cfg().MustGet(_ctx, "systemRun.runType").String()},
	}
	tool.Context().Init(r, customCtx)
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

func Demo(r *ghttp.Request) {
	_ctx := r.Context()
	//演示例外url地址集合
	list_url_exception := g.Map{
		//登录
		fmt.Sprintf("/%s/auth/login", BaseApiPath): struct{}{},
		//退出
		fmt.Sprintf("/%s/auth/logout", BaseApiPath): struct{}{},
	}
	_res := tool.DoRes[any]{}
	_run_type := gconv.Map(tool.Context().Get(_ctx).OtherInfo)
	if _run_type != nil && gconv.String(_run_type["run_type"]) == "demo" {
		if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodDelete {
			_, ok := list_url_exception[r.URL.Path]
			if ok {
				r.Middleware.Next()
			} else {
				_res.Code = -1
				_res.Msg = "演示模式，不允许操作"
				_res.RtJs(_ctx, _res)
				return
			}
		} else {
			r.Middleware.Next()
		}
	}
	r.Middleware.Next()
}

// 鉴权
func Auth(r *ghttp.Request) {
	_ctx := r.Context()
	//获取token，如果token有时效，可以做刷新令牌
	authHeader := r.GetHeader("Authorization")
	_res := tool.DoRes[any]{}
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			_res.Code = -1
			_res.Msg = "未登录或非法访问!"
			_res.RtJs(_ctx, _res)
		} else if parts[1] == "" {
			_res.Code = -1
			_res.Msg = "未登录或非法访问!"
			_res.RtJs(_ctx, _res)
		} else {
			_token := parts[1]
			tool.Context().SetToken(_ctx, _token)
			_user_map := auth.GetUserToken(_token)
			var _user_info tool.UserInfoModel
			gconv.Struct(_user_map, &_user_info)
			tool.Context().SetUser(_ctx, &_user_info)
			// 执行下一步请求逻辑
			r.Middleware.Next()
		}
	}
}

// 自由鉴权
func AuthToken(r *ghttp.Request) {
	_ctx := r.Context()
	//获取token，如果token有时效，可以做刷新令牌
	authHeader := r.GetHeader("token")
	_res := tool.DoRes[any]{}
	if authHeader == "" {
		_res.Code = -1
		_res.Msg = "未登录或非法访问!"
		_res.RtJs(_ctx, _res)
	} else {
		r.Middleware.Next()
	}
}
