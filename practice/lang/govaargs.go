// go 可变参数
// author: baoqiang
// time: 2018/12/25 下午9:17
package lang

import (
	"fmt"
	"log"
)

func Format(format string, a ...interface{}) string {
	//log.Printf("a type is: %v", reflect.TypeOf(a).Name())

	for _, v := range a {
		log.Print(v)
	}

	// 需要加上三个点...
	return fmt.Sprintf(format, a...)
}

func RunFormat() {
	format := "name:%s, age: %d"
	log.Printf(Format(format, "xiaobao", 18))
}
