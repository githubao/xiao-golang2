// go 切片
// author: baoqiang
// time: 2018/12/25 下午9:24
package lang

import "log"

func RunSlice() {
	s := []int{1, 2, 3}
	t := s[:2]
	t[1] = 999

	// 不同于Python，go的slice共享底部数据
	log.Printf("now s is: %v", s)

	//创建新的数据的两种方式
	t1 := append([]int{}, s[:2]...)
	t2 := make([]int, 2)
	copy(t2, s[:2])

	log.Printf("t1: %v", t1)
	log.Printf("t2: %v", t2)

}
