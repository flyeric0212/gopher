package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

const (
	AES_KEY = "liangboo1464345177----1577808000"
)

var (
	AESBlock  cipher.Block
	AesKeyArr []byte
)

func init() {
	var err error

	AesKeyArr = []byte(AES_KEY)
	AESBlock, err = aes.NewCipher(AesKeyArr)
	if err != nil {
		panic(err)
	}
}

func testAes() {
	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	src := "{\"address\":\"b0cab88a2d0ddb9efbf46facac2e4491677a3787\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"40cd36724af2877347ff88551e95c6b09c0f831ab5cfcb3612bb4546430c0fc5\",\"cipherparams\":{\"iv\":\"2dd7d8a43442d1ad7b6439daaf6290ee\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"f3785e4a1413638329d7229c149f142a3a32b6d6d67e68629142647dd389adfa\"},\"mac\":\"480a539a7c18019e0f1bc252d7a7759d243496bda581d26e0dbfb281be6d9126\"},\"id\":\"ed8638e2-44d4-48d2-8f43-52d4d0be2b22\",\"version\":3}"
	result := AesEncrypt(string(src))
	fmt.Println(string(result))
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	fmt.Println(len(base64.StdEncoding.EncodeToString(result)))

	ret, _ := base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString(result))
	//fmt.Println(string(ret))
	origData := AesDecrypt(string(ret))
	fmt.Println(string(origData))
	fmt.Println(len(string(origData)))
}

func AesEncrypt(origData string) []byte {
	origin := []byte(origData)

	blockSize := AESBlock.BlockSize()
	origin = PKCS5Padding(origin, blockSize)
	blockMode := cipher.NewCBCEncrypter(AESBlock, AesKeyArr[:blockSize])
	crypted := make([]byte, len(origin))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origin)
	return crypted
}

func AesDecrypt(cryptedData string) []byte {
	crypted := []byte(cryptedData)

	blockSize := AESBlock.BlockSize()
	blockMode := cipher.NewCBCDecrypter(AESBlock, AesKeyArr[:blockSize])

	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
