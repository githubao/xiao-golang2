// go获取指针
// author: baoqiang
// time: 2018/12/26 下午2:32
package lang

import (
	"log"
	"runtime"
	"time"
	"unsafe"
)

func GetPtr() {
	var x = 42
	var p = uintptr(unsafe.Pointer(&x))

	runtime.GC()

	time.Sleep(time.Second * 2)

	// gc之后可能拿不到刚才的地址
	var px = (*int)(unsafe.Pointer(p))
	log.Printf("origin p: %v, got px: %v", p, *px)

}
