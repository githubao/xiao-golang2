// 加载使用插件
// author: baoqiang
// time: 2018/11/28 下午8:31
package impl

import (
	"plugin"
	"fmt"
)

func main() {
	RunPlugin()
}

func RunPlugin() {
	p, err := plugin.Open("./aplugin.so")
	fmt.Println(err)

	add, _ := p.Lookup("Add")
	sub, _ := p.Lookup("Sub")

	addFunc := add.(func(int, int) int)
	subFunc := sub.(func(int, int) int)

	fmt.Println(addFunc(1, 2))
	fmt.Println(subFunc(5, 4))

}
