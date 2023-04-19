package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

// GetPages 获取分页参数
func GetPages(page, size int64) (offset int64) {
	offset = (page - 1) * size
	return
}

// CeilPages 计算总页数
func CeilPages(num, pageSize int64) int64 {
	if num < pageSize {
		return 1
	}
	var d int64 = 0
	if num%pageSize > 0 {
		d = 1
	}
	return num/pageSize + d
}

func MD5(params string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(params))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

// Base64 base64...
func Base64(params string) string {
	return base64.StdEncoding.EncodeToString([]byte(params))
}

// Shuffle 打乱数组原因顺序
func Shuffle(slice []interface{}) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

// GetFileExt 获取文件后缀
func GetFileExt(fp multipart.File) string {
	buffer := make([]byte, 32)
	if _, err := fp.Read(buffer); err != nil {
		// 获取失败
		return ""
	}
	return http.DetectContentType(buffer)
}

func Sha1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// GenerateSecret 生成密码加密串
func GenerateSecret(n int) string {
	if n == 0 {
		rand.Seed(time.Now().UnixNano())
		n = rand.Intn(15)
		if n < 3 {
			n = 8
		}
	}
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	bytes := []byte(str)
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = bytes[rand.Intn(len(bytes))]
	}
	return string(result)
}

// Password 登陆密码
func Password(pass, secret string) string {
	return strings.ToUpper(MD5(base64.StdEncoding.EncodeToString([]byte(secret + pass + secret))))
}

// StringToFloat string to float64
func StringToFloat(d string) float64 {
	f, err := strconv.ParseFloat(d, 64)
	if err != nil {
		return 0
	}
	return f
}

// Round 保留小数位
func Round(f float64, n int) float64 {
	d, _ := decimal.NewFromFloat(f).Round(int32(n)).Float64()
	return d
}

// BufferConcat 字符串拼接
func BufferConcat(s []string, seq string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < len(s); i++ {
		if i > 0 && seq != "" {
			buf.WriteString(seq)
		}
		buf.WriteString(s[i])
	}
	return buf.String()
}

// GenValidateCode 生成短信验证码
func GenValidateCode(width int) (code string, err error) {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		if _, err = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)]); err != nil {
			return "", err
		}
	}
	code = sb.String()
	return
}

//AesEncrypt 加密
func AesEncrypt(data []byte, key []byte) ([]byte, error) {
	//创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//判断加密快的大小
	blockSize := block.BlockSize()
	//填充
	encryptBytes := pkcs7Padding(data, blockSize)
	//初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	//执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)
	return crypted, nil
}

//pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

//AesDecrypt 解密
func AesDecrypt(data []byte, key []byte) ([]byte, error) {
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	//初始化解密数据接收切片
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去除填充
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

//pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}
