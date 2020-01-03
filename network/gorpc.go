// 指的是golang里面自己的rpc，不是谷歌的gorpc实现
// author: baoqiang
// time: 2019-08-28 14:18
package network

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func RunRpcServer() {
	//runHttpServer()
	//runTcpServer()
	runJsonServer()
}

func RunRpcClient() {
	//runHttpClient()
	//runTcpClient()
	runJsonClient()
}

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B

	return nil
}

// http
func runHttpServer() {
	arith := new(Arith)

	rpc.Register(arith)

	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)

	if err != nil {
		fmt.Println(err.Error())
	}

}

func runHttpClient() {
	client, err := rpc.DialHTTP("tcp", ":1234")
	checkError(err)

	args := Args{17, 8}

	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	checkError(err)
	fmt.Printf("reply: %v\n", reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	checkError(err)
	fmt.Printf("quot: %v\n", quot)
}

// tcp
func runTcpServer() {
	arith := new(Arith)

	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1235")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	//rpc.Accept(listener)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}

}

func runTcpClient() {
	client, err := rpc.Dial("tcp", ":1235")
	checkError(err)

	args := Args{17, 8}

	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	checkError(err)
	fmt.Printf("reply: %v\n", reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	checkError(err)
	fmt.Printf("quot: %v\n", quot)
}

// jsonrpc
func runJsonServer() {
	arith := new(Arith)

	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1235")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	//rpc.Accept(listener)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}

}

func runJsonClient() {
	client, err := jsonrpc.Dial("tcp", ":1235")
	checkError(err)

	args := Args{17, 8}

	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	checkError(err)
	fmt.Printf("reply: %v\n", reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	checkError(err)
	fmt.Printf("quot: %v\n", quot)
}
