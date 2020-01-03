// rpc 客户端
// author: baoqiang
// time: 2019/1/18 下午3:57
package lang

import (
	"net/rpc"
	"fmt"
)

func RunRpcClient() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	HandlerError(err)

	var reply string
	err = client.Call("HelloService.Hello", "aaa", &reply)
	HandlerError(err)

	fmt.Println(reply)

}
