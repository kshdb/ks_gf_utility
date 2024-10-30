package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// 生成加密密码
func CreatePwdHash(_pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(_pwd), bcrypt.DefaultCost)
	return string(bytes), err
}

// 验证密码
func CheckPwdHash(_pwd, _hash_pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(_hash_pwd), []byte(_pwd))
	return err == nil
}
