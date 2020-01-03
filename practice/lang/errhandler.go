// 错误处理
// author: baoqiang
// time: 2018/12/21 下午12:49
package lang

import "log"

func HandlerError(err error) {
	if err != nil {
		log.Fatalf("%v\n", err)
	}
}
