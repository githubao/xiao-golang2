// rpc 服务端
// author: baoqiang
// time: 2019/1/18 下午3:57
package lang

import (
	"net/rpc"
	"net"
	"fmt"
	"time"
)

type HelloService struct{}

func (p *HelloService) Hello(req string, resp *string) error {
	*resp = fmt.Sprintf("Reply: %s", req)
	return nil
}

func RunRpcServer() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	HandlerError(err)

	conn, err := listener.Accept()
	HandlerError(err)

	rpc.ServeConn(conn)

}

func RunRpc() {
	go RunRpcServer()

	time.Sleep(3 * time.Second)

	RunRpcClient()

}
