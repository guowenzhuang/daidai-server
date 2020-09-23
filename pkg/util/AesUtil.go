package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
)

// 解谜
func Dnctypt(rawData, key, iv string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	key_b, err_1 := base64.StdEncoding.DecodeString(key)
	iv_b, _ := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}
	if err_1 != nil {
		return "", err_1
	}
	dnData, err := aesCBCDncrypt(data, key_b, iv_b)
	if err != nil {
		return "", err
	}
	return string(dnData), nil
}

func aesCBCDncrypt(encrypeData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("微信解谜失败", err)
	}
	blockSize := block.BlockSize()
	if len(encrypeData) < blockSize {
		log.Println("ciphertext too short")
	}
	if len(encrypeData)%blockSize != 0 {
		log.Println("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encrypeData, encrypeData)
	// 解填充
	encrypeData = pKCS7UnPadding(encrypeData)
	return encrypeData, nil
}

// 去除填充
func pKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
