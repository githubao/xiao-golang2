// 快速排序
// author: baoqiang
// time: 2019/2/28 下午1:30
package impl

import (
	"math/rand"
	"fmt"
)

func RunQsort() {
	n := 10000
	data := make([]int, 0, n)

	for i := 0; i < n; i++ {
		data = append(data, rand.Intn(n))
	}

	qsort(data)

	for _, val := range data {
		fmt.Println(val)
	}
}

func qsort(lst []int) {
	if len(lst) <= 1 {
		return
	}

	head, tail := 0, len(lst)-1
	mid := lst[0]

	for i := 1; i <= tail; {
		if lst[i] > mid {
			lst[i], lst[tail] = lst[tail], lst[i]
			tail--
		} else {
			lst[head], lst[i] = lst[i], lst[head]
			head++
			i++
		}
	}

	qsort(lst[:head])
	qsort(lst[head+1:])

}
