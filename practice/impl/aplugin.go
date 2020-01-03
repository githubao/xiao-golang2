// 一个插件
// author: baoqiang
// time: 2018/11/28 下午8:28
package impl

import "fmt"

func Add(x, y int) int {
	fmt.Println("add load from aplugin")
	return x + y
}

func Sub(x, y int) int {
	fmt.Println("sub load from aplugin")
	return x - y
}
