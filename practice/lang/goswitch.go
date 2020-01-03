// switch impl
// author: baoqiang
// time: 2019/2/28 下午4:24
package lang

import "fmt"

// 多个switch的case不会默认fallthrough,使用多个case在同一行的并列实现
func RunSwitch() {
	s := 6
	switch true {
	case s%2 == 0:
		//fmt.Println("divide by 2")
	case s%3 == 0:
		fmt.Println("divide by 3")
	}

	switch true {
	case s%2 == 0,s%3 == 0:
		fmt.Println("divide by 2 or 3")
	}
}
