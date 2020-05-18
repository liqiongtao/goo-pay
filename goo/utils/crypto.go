package gooUtils

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

func MD5(buf []byte) string {
	h := md5.New()
	h.Write(buf)
	return hex.EncodeToString(h.Sum(nil))
}

func SHA1(buf []byte) string {
	h := sha1.New()
	h.Write(buf)
	return hex.EncodeToString(h.Sum(nil))
}

func SHA256(buf, key []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(buf)
	return hex.EncodeToString(h.Sum(nil))
}

func HMacMd5(buf, key []byte) string {
	h := hmac.New(md5.New, key)
	h.Write(buf)
	return hex.EncodeToString(h.Sum(nil))
}

func HMacSha1(buf, key []byte) string {
	h := hmac.New(sha1.New, key)
	h.Write(buf)
	return hex.EncodeToString(h.Sum(nil))
}

func HMacSha256(buf, key []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(buf)
	return hex.EncodeToString(h.Sum(nil))
}

func Base64Encode(buf []byte) string {
	return base64.StdEncoding.EncodeToString(buf)
}

func Base64Decode(str string) []byte {
	var count = (4 - len(str)%4) % 4
	str += strings.Repeat("=", count)
	buf, _ := base64.StdEncoding.DecodeString(str)
	return buf
}

func AES256ECBEncrypt(data, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := cipher.BlockSize()
	paddingSize := blockSize - len(data)%blockSize
	if paddingSize != 0 {
		data = append(data, bytes.Repeat([]byte{byte(0)}, paddingSize)...)
	}
	encrypted := make([]byte, len(data))
	for bs, be := 0, blockSize; bs < len(data); bs, be = bs+blockSize, be+blockSize {
		cipher.Encrypt(encrypted[bs:be], data[bs:be])
	}
	return encrypted, nil
}

func AES256ECBDecrypt(buf, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := cipher.BlockSize()
	fmt.Println("=======1===", len(buf), blockSize)
	decrypted := make([]byte, len(buf))
	for bs, be := 0, blockSize; bs < len(buf); bs, be = bs+blockSize, be+blockSize {
		cipher.Decrypt(decrypted[bs:be], buf[bs:be])
	}
	paddingSize := int(decrypted[len(decrypted)-1])
	return decrypted[0 : len(decrypted)-paddingSize], nil
}

// func pkcs7unpadding(plantText []byte, blockSize int) []byte {
// 	length := len(plantText)
// 	unpadding := int(plantText[length-1])
// 	if length < unpadding {
// 		return nil
// 	}
// 	return plantText[:(length - unpadding)]
// }

func Encrypt(origData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if iv == nil {
		iv = key[:blockSize]
	}
	origData = pkcs7padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func Decrypt(crypted, key, iv []byte) (origData []byte, err error) {
	var block cipher.Block
	block, err = aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if iv == nil {
		iv = key[:blockSize]
	}
	defer func() {
		if r := recover(); r != nil {
			origData = nil
			err = errors.New(r.(string))
		}
	}()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData = make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pkcs7unpadding(origData, blockSize)
	return
}

func pkcs7padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs7unpadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	if length < unpadding {
		return nil
	}
	return plantText[:(length - unpadding)]
}

func SHAwithRSA(key, data []byte) (string, error) {
	pkey, err := x509.ParsePKCS8PrivateKey(key)
	if err != nil {
		return "", err
	}

	h := crypto.Hash.New(crypto.SHA1)
	h.Write(data)
	hashed := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, pkey.(*rsa.PrivateKey), crypto.SHA1, hashed)
	if err != nil {
		return "", err
	}

	sign := base64.StdEncoding.EncodeToString(signature)
	return sign, nil
}
