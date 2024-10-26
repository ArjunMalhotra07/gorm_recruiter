package apigateway

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// ! This should be in an env file in production
func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// ! Encrypt method is to encrypt or hide any classified text
func Encrypt(text, secretKey string) (string, error) {
	fmt.Println(secretKey)
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}
