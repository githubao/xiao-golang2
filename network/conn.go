// 自己写一个连接程序
// author: baoqiang
// time: 2019-08-23 18:04
package network

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func HttpGet() {
	// got 302
	//_ = httpGet("www.google.com:80")
	_ = httpGet("127.0.0.1:1200")
}

func httpGet(hostWithPort string) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", hostWithPort)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	// 发送数据
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Printf("Got data: %v\n", string(result))

	return nil
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(-1)
	}
}
