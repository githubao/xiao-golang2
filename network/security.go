// 安全性
// author: baoqiang
// time: 2019-08-26 17:02
package network

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"
	//"golang.org/x/crypto/blowfish"
)

// 数据安全涉及的主要问题： 认证，数据完整性，保密，签名，访问控制，可用性
// 常见的hash算法：md5，sha1，sha256

func Security() {
	//RunMd5()
	//runBlowFish()
	//genKey()
	//genCer()
	//readCer()
	tlsServer()
}

func Security2() {
	tlsClient()
}

func RunMd5() {
	hash := md5.New()

	bytes := []byte("hello\n")
	hash.Write(bytes)

	hashValue := hash.Sum(nil)
	hashSize := hash.Size()

	for n := 0; n < hashSize; n += 4 {
		var val uint32
		val = uint32(hashValue[n])<<24 +
			uint32(hashValue[n+1])<<16 +
			uint32(hashValue[n+2])<<8 +
			uint32(hashValue[n+3])

		fmt.Printf("%x ", val)
	}
	fmt.Println()

	fmt.Printf("%x\n", hashValue)
}

//aes对称加密
// blowfish 对称加密
func runBlowFish() {
	//key := []byte("the power")
	//cipher, err := blowfish.NewClipher(key)
	//checkError(err)
	//
	//src := []byte("hello\n\n\n")
	//var enc [512]byte
	//
	////encode
	//cipher.Encrypt(enc[0:], src)
	//
	////decode
	//var decrypt []byte
	//cipher.Decrypt(decrypt[0:], enc[0:])
	//
	//// get results
	//result := bytes.NewBuffer(nil)
	//result.Write(decrypt[0:8])
	//fmt.Println(string(result.Bytes()))
}

// 生成rsa公钥和私钥
func genKey() {
	reader := rand.Reader
	bitSize := 512

	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	fmt.Println("Private key primes", key.Primes[0].String(), key.Primes[1].String())
	fmt.Println("Private key exponent", key.D.String())

	publicKey := key.PublicKey
	fmt.Println("Public key modules", publicKey.N.String())
	fmt.Println("Public key exponent", publicKey.E)

	saveGobKey("ssh/private.key", key)
	saveGobKey("ssh/public.key", publicKey)

	savePEMKey("ssh/private.pem", key)
}

// 为自己的站点生成一个证书
func genCer() {
	random := rand.Reader

	// load
	var key rsa.PrivateKey
	loadKey("ssh/private.key", &key)

	now := time.Now()
	then := now.Add(time.Hour * 24 * 365)

	// template
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:   "jan.newmarch.name",
			Organization: []string{"Jan Newmarch"},
		},
		NotBefore:             now,
		NotAfter:              then,
		SubjectKeyId:          []byte{1, 2, 3, 4},
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageKeyEncipherment | x509.KeyUsageKeyEncipherment,
		BasicConstraintsValid: true,
		IsCA:                  true,
		DNSNames:              []string{"jan.newmarch.name", "localhost"},
	}

	// create
	derBytes, err := x509.CreateCertificate(random, &template,
		&template, &key.PublicKey, &key)
	checkError(err)

	// save cer
	//certCerFile, err := os.Create("jan.newmarch.name.cer")
	//checkError(err)
	//count, err := certCerFile.Write(derBytes)
	ioutil.WriteFile("jan.newmarch.name.cer", derBytes, 0644)
	//checkError(err)
	//fmt.Printf("write len: %v\n", count)
	//certCerFile.Close()

	// write pem
	certPEMFile, err := os.Create("jan.newmarch.name.pem")
	checkError(err)
	pem.Encode(certPEMFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certPEMFile.Close()

	// private key
	//keyPEMFile, err := os.Create("newmarch_private.pem")
	//checkError(err)
	//pem.Encode(keyPEMFile, &pem.Block{Type: "RAS PRIVATE KEY",
	//	Bytes: x509.MarshalPKCS1PrivateKey(&key)})
	//keyPEMFile.Close()

}

// 读取证书
func readCer() {
	certCerFile, err := os.Open("jan.newmarch.name.cer")
	checkError(err)

	derBytes := make([]byte, 1000)
	count, err := certCerFile.Read(derBytes)
	checkError(err)
	certCerFile.Close()

	cert, err := x509.ParseCertificate(derBytes[0:count])
	checkError(err)

	fmt.Printf("Name: %s\n", cert.Subject.CommonName)
	fmt.Printf("Not Before: %s\n", cert.NotBefore.String())
	fmt.Printf("Not After: %s\n", cert.NotAfter.String())

}

// tls server
func tlsServer() {
	cerFile := "ssh/jan.newmarch.name.pem"
	keyFile := "ssh/private.pem"

	cert, err := tls.LoadX509KeyPair(cerFile, keyFile)
	checkError(err)

	config := tls.Config{Certificates: []tls.Certificate{cert}}

	now := time.Now()
	config.Time = func() time.Time {
		return now
	}
	config.Rand = rand.Reader

	service := "0.0.0.0:1200"
	listener, err := tls.Listen("tcp", service, &config)

	// do listener logic
	for {
		conn, err := listener.Accept()
		checkError(err)
		fmt.Printf("Got conn: %v\n", conn.RemoteAddr())

	}
}

func tlsClient() {
	service := "0.0.0.0:1200"

	conn, err := tls.Dial("tcp", service, nil)
	checkError(err)
	fmt.Println(conn)

	// do conn logic

}

func saveGobKey(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)

	outFile.Close()
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)

	var privateKey = &pem.Block{Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key)}

	pem.Encode(outFile, privateKey)

	outFile.Close()
}

func loadKey(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkError(err)

	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)

	inFile.Close()
}
