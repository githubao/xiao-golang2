// 一些获取ip的接口
// author: baoqiang
// time: 2019-08-23 17:35
package network

import (
	"fmt"
	"net"
)

func NetworkSample() {
	//resolveIp()
	//lookupPort()
	resolveTcp()
}

func resolveIp() {
	// ip 就是 []byte类型

	ipStr := "172.217.160.36"
	sampleIPAddr := net.ParseIP(ipStr)
	fmt.Printf("%s -> %v\n", ipStr, sampleIPAddr)

	google := "www.google.com"
	googleIPAddr, _ := net.ResolveIPAddr("ip", google)
	fmt.Printf("%s -> %v\n", google, googleIPAddr)
}

func lookupPort() {
	service := "domain"
	port, _ := net.LookupPort("tcp", service)
	fmt.Printf("%s -> %v\n", service, port)
}

func resolveTcp() {
	// tcpAddr 有两个字段，一个是ip类型，一个是int类型的port

	ipPortStr := "www.google.com:80"
	sampleTCPAddr, _ := net.ResolveTCPAddr("tcp", ipPortStr)
	fmt.Printf("%s -> %v\n", ipPortStr, sampleTCPAddr)
}
