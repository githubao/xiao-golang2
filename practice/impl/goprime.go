// 素数筛
// author: baoqiang
// time: 2019/3/1 下午3:38
package impl

import "fmt"

func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// 有可能被放进dest里面的，都已经是前面没有被整除的了，所以他们都是素数
func filter(src <-chan int, dest chan<- int, prime int) {
	for i := range src {
		if i%prime != 0 {
			dest <- i
		}
	}
}

func RunPrime() {
	ch := make(chan int)
	go generate(ch)

	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Println(prime)
		out := make(chan int)
		go filter(ch, out, prime)
		ch = out
	}
}
