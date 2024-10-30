package tool

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

/*
基础响应对象
*/
type DoRes[T g.Map | g.List | any] struct {
	Code int    `json:"code" dc:"状态码"`
	Data T      `json:"data,omitempty" dc:"返回数据"`
	Msg  string `json:"msg" dc:"返回消息"`
}

/*
输出json
*/
func (res *DoRes[T]) RtJs(ctx context.Context, _json any) {
	g.RequestFromCtx(ctx).Response.WriteJson(_json)
}
