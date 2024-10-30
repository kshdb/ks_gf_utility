package auth

import (
	"encoding/base64"
	"github.com/wumansgy/goEncrypt/aes"
)

func GetAesDe(ciphertext, _key string) (plaintext []byte, err error) {
	// 解码 Base64 字符串
	key_str, _err := Base64Decode(_key)
	if err != nil {
		err = _err
	}
	plaintext, err = aes.AesEcbDecryptByBase64(ciphertext, key_str)
	return
}
func Base64Decode(encoded string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
