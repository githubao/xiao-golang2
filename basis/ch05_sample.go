// Created by BaoQiang at 2017/4/29 10:07

package main

import (
	"fmt"
	"reflect"
)

type (
	str string
)

func main() {
	//VarDemo()
	//OperateDemo()
	//PtrDemo()
	//ArrayDemo()
	//SliceDemo()
	//MapDemo()
	//FuncDemo()
	ClosureDemo()
}

//for range 的闭包特性，下例中三次都输出c

func ClosureDemo() {
	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func() {
			fmt.Println(v)
		}()
	}

	//纠正
	for _, v := range s {
		go func(v string) {
			fmt.Println(v)
		}(v)
	}

	select {}
}

func FuncDemo() {
	//fmt.Println(FuncRun(1, 2))
	//f := closure(1)
	//fmt.Println(f(3))
	//DeferDemo()
	PanicDemo()
}

func PanicDemo() {
	A()
	B()
	C()
}

func A() {
	fmt.Println("Run in a")
}
func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from b")
		}
	}()
	panic("panic in b")
	fmt.Println("run in b")

}
func C() {
	fmt.Println("Run in c")
}

func closure(x int) (func(int) (int)) {
	return func(y int) int {
		return x + y
	}
}

func DeferDemo() {
	// main函数运行完了再执行，后进先出的执行
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	//闭包
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func FuncRun(a int, b int) (int) {
	return a + b
}

func MapDemo() {
	//m := make(map[string]int)
	m := map[string]int{}
	m["2"] = 3
	fmt.Println(m)

	for k, _ := range m {
		m[k] = 2
	}
	fmt.Println(m)
}

func SliceDemo() {
	//左闭右开

	//数组和切片
	arr := [...]int{1, 2}
	arr2 := []int{1, 2}
	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(reflect.TypeOf(arr2))

	sli := make([]int, 5)
	fmt.Println(sli)

	//copy(dst,src)

}

func ArrayDemo() {
	//arr := []int{1, 2}
	arr := [...]int{1, 2}
	fmt.Println(arr)

	//指针类型 的数组
	x, y := 1, 2
	arr2 := []*int{&x, &y}
	fmt.Println(arr2)

	//数组类型 的指针
	fmt.Println(&arr)

	//new返回指针
	p := new([10]int)
	fmt.Println(p)
	fmt.Println(p[1])
}

func PtrDemo() {
	a := 1
	p := &a
	println(p) //0xc042031f20
	println(*p)
}

var (
	aa = "aa"
	bb = "bb"
)

var a byte = 'a'
//a := 'a'

const (
	AA = 'A'
	BB
	CC = iota
	DD
)

func OperateDemo() {
	a := 1
	println(^a)

	//&^
	//后面的运算数，是1的时候，把上面的运算数强制改为0
	//^
	//按位取反
}

func VarDemo() {
	//var s str
	//s = "eee"
	//var s = "222"
	//定义并赋值
	s := "333"
	println(s)

	//全局变量必须使用var关键字
	var b byte = 3
	println(b)

	//并行初始化
	//cc,dd := "cc","dd"

	println(AA)
	println(BB)
	println(CC)
	println(DD)

}

func HeadIn() {
	a := 1
	println(^a)
}
