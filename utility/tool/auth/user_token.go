package auth

import (
	"github.com/gogf/gf/v2/frame/g"
)

var (
	// 签名字符串
	token_sign = ""
)

/*
设置用户token
*/
func SetUserToken(_data g.Map) (_token string) {
	// -----------  生成jwt token -----------
	token := NewJWToken(token_sign)
	_token, _ = token.GenJWToken(_data)
	return
}

/*
获取用户序token
*/
func GetUserToken(_token string) (_token_map g.Map) {
	// -----------  解析 jwt token -----------
	token := NewJWToken(token_sign)
	_token_map, _ = token.ParseJWToken(_token)
	return
}
