// 指针内存操作
// author: baoqiang
// time: 2019-07-31 14:34
package uptr

import (
	"fmt"
	"unsafe"
)

type User struct {
	name string
	age  int
}

func ExampleUptr() {
	u := new(User)
	fmt.Println(*u)

	pName := (*string)(unsafe.Pointer(u))
	*pName = "小包"

	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)))
	*pAge = 18

	fmt.Println(*u)
}
