// 在app层面实现连接和通信
// author: baoqiang
// time: 2019-08-26 15:40
package network

import (
	"net"
	"os"
)

const (
	DIR = "DIR"
	CD  = "CD"
	PWD = "PWD"
)

func RunServer() {
	service := "0.0.0.0:1202"

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			conn.Close()
			return
		}

		s := string(buf[0:n])
		if s[0:2] == CD {
			chdir(conn, s[3:])
		} else if s[0:3] == DIR {
			dirList(conn)
		} else if s[0:3] == PWD {
			pwd(conn)
		}

	}

}

func chdir(conn net.Conn, s string) {
	if os.Chdir(s) == nil {
		_, _ = conn.Write([]byte("OK"))
	} else {
		_, _ = conn.Write([]byte("ERROR"))
	}
}

func dirList(conn net.Conn) {
	defer conn.Write([]byte("\r\n"))

	dir, err := os.Open(".")
	if err != nil {
		return
	}

	names, err := dir.Readdirnames(-1)
	if err != nil {
		return
	}

	for _, nm := range names {
		_, _ = conn.Write([]byte(nm + "\r\n"))
	}

}

func pwd(conn net.Conn) {
	s, err := os.Getwd()
	if err != nil {
		_, _ = conn.Write([]byte(""))
		return
	}

	_, _ = conn.Write([]byte(s))

}
