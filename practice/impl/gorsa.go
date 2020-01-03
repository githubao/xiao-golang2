// go rsa，签名算法
// author: baoqiang
// time: 2019/3/1 下午8:32
package impl

import (
	"encoding/pem"
	"crypto/x509"
	"errors"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/rand"
	"crypto"
	"encoding/base64"
	"fmt"
)

var errGlobal = errors.New("global err")

func RunSign() {

	// 待签名的数据
	toSign := "secret data"

	// 签名类
	signer, _ := loadPrivateKey("private.pem")

	// 签名后的数据
	signed, _ := signer.Sign([]byte(toSign))

	// 签名后的字符串
	sig := base64.StdEncoding.EncodeToString(signed)

	fmt.Println(sig)

	// 签名类
	parser, _ := loadPublicKey("public.pem")

	//原始数据有没有被篡改
	faked := "changed data" //crypto/rsa: verification error
	//err := parser.UnSign([]byte(toSign), signed)
	err := parser.UnSign([]byte(faked), signed)
	fmt.Println(err)
}

// 加解密接口
type Signer interface {
	Sign(data []byte) ([]byte, error)
}

type UnSigner interface {
	UnSign(data []byte, sig []byte) (error)
}

//使用rsa实现上面的两个接口
type rsaPrivateKey struct {
	*rsa.PrivateKey
}

type rsaPublicKey struct {
	*rsa.PublicKey
}

func (r *rsaPrivateKey) Sign(data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	d := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, r.PrivateKey, crypto.SHA256, d)
}

func (r *rsaPublicKey) UnSign(data []byte, sig []byte) (error) {
	h := sha256.New()
	h.Write(data)
	d := h.Sum(nil)
	return rsa.VerifyPKCS1v15(r.PublicKey, crypto.SHA256, d, sig)
}

// 加密
func parsePrivateKey(pemBytes []byte) (Signer, error) {
	block, _ := pem.Decode(pemBytes)

	var rawKey interface{}
	switch block.Type {
	case "RSA PRIVATE KEY":
		rsa, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
		rawKey = rsa
	default:
		return nil, errGlobal
	}

	return newSignerFromKey(rawKey)
}

func newSignerFromKey(k interface{}) (Signer, error) {
	var sshKey Signer

	switch t := k.(type) {
	case *rsa.PrivateKey:
		sshKey = &rsaPrivateKey{t}
	default:
		return nil, errGlobal
	}

	return sshKey, nil
}

// 解密
func parsePublicKey(pemBytes []byte) (UnSigner, error) {
	block, _ := pem.Decode(pemBytes)

	var rawKey interface{}
	switch block.Type {
	case "PUBLIC KEY":
		rsa, _ := x509.ParsePKIXPublicKey(block.Bytes)
		rawKey = rsa
	default:
		return nil, errGlobal
	}

	return newUnSignerFromKey(rawKey)
}

func newUnSignerFromKey(k interface{}) (UnSigner, error) {
	var sshKey UnSigner

	switch t := k.(type) {
	case *rsa.PublicKey:
		sshKey = &rsaPublicKey{t}
	default:
		return nil, errGlobal
	}

	return sshKey, nil
}

// 密钥

func loadPrivateKey(path string) (Signer, error) {
	return parsePrivateKey([]byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDCFENGw33yGihy92pDjZQhl0C36rPJj+CvfSC8+q28hxA161QF
NUd13wuCTUcq0Qd2qsBe/2hFyc2DCJJg0h1L78+6Z4UMR7EOcpfdUE9Hf3m/hs+F
UR45uBJeDK1HSFHD8bHKD6kv8FPGfJTotc+2xjJwoYi+1hqp1fIekaxsyQIDAQAB
AoGBAJR8ZkCUvx5kzv+utdl7T5MnordT1TvoXXJGXK7ZZ+UuvMNUCdN2QPc4sBiA
QWvLw1cSKt5DsKZ8UETpYPy8pPYnnDEz2dDYiaew9+xEpubyeW2oH4Zx71wqBtOK
kqwrXa/pzdpiucRRjk6vE6YY7EBBs/g7uanVpGibOVAEsqH1AkEA7DkjVH28WDUg
f1nqvfn2Kj6CT7nIcE3jGJsZZ7zlZmBmHFDONMLUrXR/Zm3pR5m0tCmBqa5RK95u
412jt1dPIwJBANJT3v8pnkth48bQo/fKel6uEYyboRtA5/uHuHkZ6FQF7OUkGogc
mSJluOdc5t6hI1VsLn0QZEjQZMEOWr+wKSMCQQCC4kXJEsHAve77oP6HtG/IiEn7
kpyUXRNvFsDE0czpJJBvL/aRFUJxuRK91jhjC68sA7NsKMGg5OXb5I5Jj36xAkEA
gIT7aFOYBFwGgQAQkWNKLvySgKbAZRTeLBacpHMuQdl1DfdntvAyqpAZ0lY0RKmW
G6aFKaqQfOXKCyWoUiVknQJAXrlgySFci/2ueKlIE1QqIiLSZ8V8OlpFLRnb1pzI
7U1yQXnTAEFYM560yJlzUpOb1V4cScGd365tiSMvxLOvTA==
-----END RSA PRIVATE KEY-----`))
}

func loadPublicKey(path string) (UnSigner, error) {

	return parsePublicKey([]byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDCFENGw33yGihy92pDjZQhl0C3
6rPJj+CvfSC8+q28hxA161QFNUd13wuCTUcq0Qd2qsBe/2hFyc2DCJJg0h1L78+6
Z4UMR7EOcpfdUE9Hf3m/hs+FUR45uBJeDK1HSFHD8bHKD6kv8FPGfJTotc+2xjJw
oYi+1hqp1fIekaxsyQIDAQAB
-----END PUBLIC KEY-----`))
}
