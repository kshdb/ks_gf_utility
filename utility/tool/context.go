package tool

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

const (
	ContextKey = "ContextKey" // 上下文变量存储键名，前后端系统共享
)

type sContext struct{}

// Context 上下文管理服务
func Context() *sContext {
	return &sContext{}
}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *sContext) Init(r *ghttp.Request, customCtx *UserContext) {
	r.SetCtxVar(ContextKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func (s *sContext) Get(ctx context.Context) *UserContext {
	value := ctx.Value(ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*UserContext); ok {
		return localCtx
	}
	return nil
}

// SetToken 上下文中写入token信息
func (s *sContext) SetToken(ctx context.Context, _token string) {
	s.Get(ctx).Token = _token
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *sContext) SetUser(ctx context.Context, ctxUser *UserInfoModel) {
	s.Get(ctx).UserInfo = ctxUser
}

// SetData 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *sContext) SetOther(ctx context.Context, _map g.Map) {
	s.Get(ctx).OtherInfo = _map
}

// 用户上下文
type UserContext struct {
	UserInfo  *UserInfoModel `json:"userInfo" dc:"用户信息"`
	Token     string         `json:"token" dc:"用户token"`
	OtherInfo g.Map          `json:"otherInfo" dc:"其它kv"` // 自定KV变量，业务模块根据需要设置，不固定
}

/*
用户信息
*/
type UserInfoModel struct {
	LoginType string `json:"loginType" dc:"登录方式"`
	LoginId   string `json:"loginId" dc:"登录id"`
	RnStr     string `json:"rnStr" dc:"随机字符串"`
	Clientid  string `json:"clientid" dc:"客户端id"`
	TenantId  string `json:"tenantId" dc:"租户id"`
	UserId    int64  `json:"userId" dc:"用户id"`
	UserName  string `json:"userName" dc:"用户名"`
	DeptId    int64  `json:"deptId" dc:"所属部门id"`
	DeptName  string `json:"deptName" dc:"所属部门名称"`
	//RoleIds   []uint   // 角色id列表
	//RoleNames []string // 角色名称
}
