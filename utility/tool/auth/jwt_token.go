package auth

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v4"
)

// JWToken jwt token
type JWToken struct {
	SignString string
}

// NewJWToken 创建JWToken对象
func NewJWToken(signString string) *JWToken {
	return &JWToken{SignString: signString}
}

// GenJWToken 生成一个jwt token
func (t *JWToken) GenJWToken(rawContent g.Map) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(rawContent))
	tokenString, err := token.SignedString([]byte(t.SignString))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseJWToken 解析 JWToken
func (t *JWToken) ParseJWToken(tokenString string) (g.Map, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.SignString), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		return claims, nil
	}
	//if ok && token.Valid {
	//	return claims, nil
	//}
	return nil, err
}
