// 切片转化
// author: baoqiang
// time: 2019/3/3 下午4:26
package impl

import (
	"reflect"
	"unsafe"
)

func ByteSlice(slice interface{}) (data []byte) {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic("no slice")
	}

	h := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	h.Cap = sv.Cap() * int(sv.Type().Elem().Size())
	h.Len = sv.Len() * int(sv.Type().Elem().Size())
	h.Data = sv.Pointer()

	return
}
