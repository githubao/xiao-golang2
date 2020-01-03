// 归并排序
// author: baoqiang
// time: 2019/3/1 下午6:04
package impl

import (
	"fmt"
	"math/rand"
)

func RunMsort() {
	n := 10000
	data := make([]int, 0, n)

	for i := 0; i < n; i++ {
		data = append(data, rand.Intn(n))
	}

	//data := []int{3, 1, 2}

	res := mergeSort(data)

	for _, val := range res {
		fmt.Println(val)
	}
}

func mergeSort(arr []int) []int {
	length := len(arr)

	if length < 2 {
		return arr
	}

	middle := length / 2

	return merge(mergeSort(arr[0:middle]), mergeSort(arr[middle:]))
}

func merge(left []int, right []int) []int {
	var result []int

	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]

	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}
