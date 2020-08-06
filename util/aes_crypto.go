// Copyright 2014 chanxuehong(chanxuehong@gmail.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// 把整数 n 格式化成 4 字节的网络字节序
func encodeNetworkByteOrder(b []byte, n uint32) {
	b[0] = byte(n >> 24)
	b[1] = byte(n >> 16)
	b[2] = byte(n >> 8)
	b[3] = byte(n)
}

// 从 4 字节的网络字节序里解析出整数
func decodeNetworkByteOrder(b []byte) (n uint32) {
	return uint32(b[0])<<24 |
		uint32(b[1])<<16 |
		uint32(b[2])<<8 |
		uint32(b[3])
}

// AESEncryptMsg 消息加密
// ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
func AESEncryptMsg(random, rawXMLMsg []byte, appId string, encodingAESKey string) (ciphertext string) {
	aesKey, _ := base64.StdEncoding.DecodeString(encodingAESKey + "=")
	const (
		BLOCK_SIZE = 32             // PKCS#7
		BLOCK_MASK = BLOCK_SIZE - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	appIdOffset := 20 + len(rawXMLMsg)
	contentLen := appIdOffset + len(appId)
	amountToPad := BLOCK_SIZE - contentLen&BLOCK_MASK
	plaintextLen := contentLen + amountToPad

	plaintext := make([]byte, plaintextLen)

	// 拼接
	copy(plaintext[:16], random)
	encodeNetworkByteOrder(plaintext[16:20], uint32(len(rawXMLMsg)))
	copy(plaintext[20:], rawXMLMsg)
	copy(plaintext[appIdOffset:], appId)

	// PKCS#7 补位
	for i := contentLen; i < plaintextLen; i++ {
		plaintext[i] = byte(amountToPad)
	}

	// 加密
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCEncrypter(block, aesKey[:16])
	mode.CryptBlocks(plaintext, plaintext)

	return base64.StdEncoding.EncodeToString(plaintext)
}

// AESDecryptMsg 消息解密
// ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
func AESDecryptMsg(base64CipherText string, encodingAESKey string) (random, rawXMLMsg, appId []byte, err error) {
	ciphertext, err := base64.StdEncoding.DecodeString(base64CipherText)
	if err != nil {
		return
	}

	aesKey, err := base64.StdEncoding.DecodeString(encodingAESKey + "=")
	if err != nil {
		return
	}

	const (
		BLOCK_SIZE = 32             // PKCS#7
		BLOCK_MASK = BLOCK_SIZE - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	if len(ciphertext) < BLOCK_SIZE {
		err = fmt.Errorf("the length of ciphertext too short: %d", len(ciphertext))
		return
	}
	if len(ciphertext)&BLOCK_MASK != 0 {
		err = fmt.Errorf("ciphertext is not a multiple of the block size, the length is %d", len(ciphertext))
		return
	}

	plaintext := make([]byte, len(ciphertext)) // len(plaintext) >= BLOCK_SIZE

	// 解密
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return
	}
	mode := cipher.NewCBCDecrypter(block, aesKey[:16])
	mode.CryptBlocks(plaintext, ciphertext)

	// PKCS#7 去除补位
	amountToPad := int(plaintext[len(plaintext)-1])
	if amountToPad < 1 || amountToPad > BLOCK_SIZE {
		err = fmt.Errorf("the amount to pad is incorrect: %d", amountToPad)
		return
	}
	plaintext = plaintext[:len(plaintext)-amountToPad]

	// 反拼接
	// len(plaintext) == 16+4+len(rawXMLMsg)+len(appId)
	if len(plaintext) <= 20 {
		err = fmt.Errorf("plaintext too short, the length is %d", len(plaintext))
		return
	}
	rawXMLMsgLen := int(decodeNetworkByteOrder(plaintext[16:20]))
	if rawXMLMsgLen < 0 {
		err = fmt.Errorf("incorrect msg length: %d", rawXMLMsgLen)
		return
	}
	appIdOffset := 20 + rawXMLMsgLen
	if len(plaintext) <= appIdOffset {
		err = fmt.Errorf("msg length too large: %d", rawXMLMsgLen)
		return
	}

	random = plaintext[:16:20]
	rawXMLMsg = plaintext[20:appIdOffset:appIdOffset]
	appId = plaintext[appIdOffset:]
	return
}

// AESDecryptData 数据解密
func AESDecryptData(cipherText []byte, aesKey []byte, iv []byte) (rawData []byte, err error) {

	const (
		BLOCK_SIZE = 32             // PKCS#7
		BLOCK_MASK = BLOCK_SIZE - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	if len(cipherText) < BLOCK_SIZE {
		err = fmt.Errorf("the length of ciphertext too short: %d", len(cipherText))
		return
	}

	plaintext := make([]byte, len(cipherText)) // len(plaintext) >= BLOCK_SIZE

	// 解密
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, cipherText)

	// PKCS#7 去除补位
	amountToPad := int(plaintext[len(plaintext)-1])
	if amountToPad < 1 || amountToPad > BLOCK_SIZE {
		err = fmt.Errorf("the amount to pad is incorrect: %d", amountToPad)
		return
	}
	plaintext = plaintext[:len(plaintext)-amountToPad]

	// 反拼接
	// len(plaintext) == 16+4+len(rawXMLMsg)+len(appId)
	if len(plaintext) <= 20 {
		err = fmt.Errorf("plaintext too short, the length is %d", len(plaintext))
		return
	}

	rawData = plaintext

	return

}
