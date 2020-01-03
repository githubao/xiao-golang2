// iota
// author: baoqiang
// time: 2019/2/28 下午2:45
package lang

import "fmt"

const (
	a = iota
	b = 3 << iota
	c
	d
)

func RunIota() {
	fmt.Println(a, b, c, d)
}
