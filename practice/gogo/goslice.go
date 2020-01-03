// slice测试
// author: baoqiang
// time: 2019/3/4 下午2:40
package gogo

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Slice []int

// 只改变了内部的引用，并没有改变外部的caller的值
func (A Slice)Append(value int) {
	A = append(A, value)
}

func (A *Slice)Append2(value int) {
	*A = append(*A, value)
}

// 添加
func (A Slice)Append3(value int) {
	A1 := append(A, value)

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&A))
	sh1 := (*reflect.SliceHeader)(unsafe.Pointer(&A1))

	fmt.Printf("A Data:%d,Len:%d,Cap:%d\n",sh.Data,sh.Len,sh.Cap)
	fmt.Printf("A1 Data:%d,Len:%d,Cap:%d\n",sh1.Data,sh1.Len,sh1.Cap)
}

func RunSlice() {
	mSlice := make(Slice, 10, 20)
	//mSlice.Append2(5)
	mSlice.Append3(5)
	fmt.Println(mSlice)
}
