// 堆排序
// author: baoqiang
// time: 2019/2/28 下午3:15
package impl

import (
	"fmt"
	"math/rand"
)

func RunHeapsort() {
	const n = 100

	datas := make([]int, 0, 10)
	for range [n]int{} {
		datas = append(datas, rand.Intn(n))
	}

	heapSort(datas)

	fmt.Println(datas)

}

func heapSort(values []int) {
	buildHeap(values)

	for i := len(values); i > 1; i-- {
		values[0], values[i-1] = values[i-1], values[0]
		adjustHeap(values[:i-1], 0)

		//fmt.Println("the heap is ", values)
	}
}

func buildHeap(values []int) {
	for i := len(values); i >= 0; i-- {
		adjustHeap(values, i)
	}
}

func adjustHeap(values []int, pos int) {
	node := pos
	length := len(values)

	for node < length {
		var child = 0
		if 2*node+2 < length {
			if values[2*node+1] > values[2*node+2] {
				child = 2*node + 1
			} else {
				child = 2*node + 2
			}
		} else if 2*node+1 < length {
			child = 2*node + 1
		}

		if child > 0 && values[child] > values[node] {
			values[node], values[child] = values[child], values[node]
			node = child
		} else {
			break
		}

	}

}
