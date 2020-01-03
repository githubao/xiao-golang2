// 自己写一个服务
// author: baoqiang
// time: 2019-08-23 18:15
package network

import (
	"fmt"
	"net"
	"time"
)

// conn 有个setTimeout的方法可以用于超时
// 一个服务器可以使用select监听多个端口，采用事件监听的机制当io(读或写)准备好的时候进行事件处理

// 2019-08-23 18:26:45.206494 +0800 CST m=+212.289420428
func SimpleTimeServer() {
	service := ":1200"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()

		fmt.Printf("got conn: %v\n", conn.RemoteAddr())

		checkError(err)

		// 单线程一次处理一个请求
		go handleConn(conn)

	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()

	daytime := time.Now().String()
	conn.Write([]byte(daytime))

}
