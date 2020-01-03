// 通道相关
// author: baoqiang
// time: 2019/2/28 下午9:59
package lang

import (
	"fmt"
	"math/rand"
	"time"
)

func RunChan111() {
	// step0
	go say("world")
	say("hello")

	// step1
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y)

	// step2
	c2 := make(chan int, 2)
	c2 <- 1
	c2 <- 2
	fmt.Println(<-c2)
	fmt.Println(<-c2)

	//step3
	c3 := make(chan int, 10)
	go fibonacci(cap(c3), c3)
	for i := range c3 {
		fmt.Println(i)
	}

	// step4
	c4 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			fmt.Println(<-c4)
		}
		quit <- 0
	}()

	fibonacci2(c4, quit)

	// step5
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom")
			return
		default:
			fmt.Println("nothing")
			time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
		}
	}

}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}

	c <- sum
}

func fibonacci(n int, c chan<- int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci2(c chan<- int, quit <-chan int) {
//func fibonacci2(c, quit chan int) {
	x, y := 0, 1

out:
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			break out
		}
	}

}
