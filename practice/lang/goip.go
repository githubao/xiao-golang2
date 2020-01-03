// 获取监听的地址和端口号
// author: baoqiang
// time: 2019/2/28 下午9:39
package lang

import (
	"net"
	"fmt"
)

func RunAddr() {
	l, _ := net.Listen("tcp", ":0")

	port := l.Addr().(*net.TCPAddr).Port
	ip := l.Addr().(*net.TCPAddr).IP
	fmt.Println(ip, port)

	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()

		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.TCPAddr:
				fmt.Println(v.IP)
			}
		}

	}

}
