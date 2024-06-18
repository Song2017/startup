package pkg

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func Decrypt(encrypted string, key []byte) string {
	ciphertext, _ := DecryptBytes(encrypted, key)
	return string(ciphertext)
}

func DecryptBytes(encrypted string, key []byte) ([]byte, error) {
	ciphertext, _ := base64.RawURLEncoding.DecodeString(encrypted)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}

func DecryptBytesIV(encrypted string, key []byte, iv []byte) ([]byte, error) {
	// java Cipher.getInstance("AES/CBC/NoPadding")
	if encrypted == "" {
		return []byte(""), nil
	}
	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	cfb := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertext))
	cfb.CryptBlocks(decrypted, ciphertext)
	return decrypted, nil
}

func EncryptIV(data string, key []byte, iv []byte) (string, error) {
	// padding character []byte{0}
	if data == "" {
		return "", nil
	}

	// 将数据转换为字节切片
	dataBytes := []byte(data)
	// 创建和初始化Cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	// 计算需要的填充大小
	blockSize := block.BlockSize()
	padding := blockSize - len(dataBytes)%blockSize
	padtext := bytes.Repeat([]byte{byte(0)}, padding)
	paddedData := append(dataBytes, padtext...)
	cipher := cipher.NewCBCEncrypter(block, iv)

	// 创建用于存放加密数据的缓冲区
	ciphertext := make([]byte, len(paddedData))
	// 执行加密
	cipher.CryptBlocks(ciphertext, paddedData)
	// Base64编码加密后的数据
	encoded := base64.StdEncoding.EncodeToString(ciphertext)

	return encoded, nil
}

func Encrypt(idata string, key []byte) string {
	data := []byte(idata)
	encryptData, _ := EncryptBytes(data, key)
	return encryptData
}

func EncryptBytes(data []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)
	return base64.RawURLEncoding.EncodeToString(ciphertext), nil
}

// base64
func B64Encode(data []byte) string {
	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded
}

func B64Decode(data string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
