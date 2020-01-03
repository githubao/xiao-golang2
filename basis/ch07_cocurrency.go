// Created by BaoQiang at 2017/4/29 20:16

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//RoutineDemo()
	//RoutineDemo2()
	//RoutineDemo3()
	//CoccurencyDemo()
	//CoccurencyDemo2()
	SelelctDemo()
	//timeOutDemo()
}

func RoutineDemo() {
	//goroutine
	ch := make(chan bool)

	go func() {
		fmt.Println("Go Go Go")
		ch <- true
	}()
	<-ch
}

func RoutineDemo2() {
	//goroutine
	ch := make(chan bool)

	go func() {
		fmt.Println("Go Go Go")
		ch <- true
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

//有缓存的时候是异步的，非阻塞的，"爱读不读"，所以这时候不会输出"Go Go Go"
func RoutineDemo3() {
	//goroutine
	ch := make(chan bool, 1)

	go func() {
		fmt.Println("Go Go Go")
		<-ch
	}()
	ch <- true
}

// 通过设置缓存的大小，让并发都能够执行完成
func CoccurencyDemo() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Task(c, i)
	}

	//只有接收到了10次消息，这个函数才可以被结束
	for i := 0; i < 10; i++ {
		<-c
	}

}

func Task(c chan bool, index int) {
	a := 0
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	//发送消息
	c <- true
}

// 使用waitGroup实现同步
func CoccurencyDemo2() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go Task2(&wg, i)
	}

	wg.Wait()

}

func Task2(wg *sync.WaitGroup, index int) {
	a := 0
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	wg.Done()
}

// select
func SelelctDemo() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)

	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "Hello"
	c1 <- 3
	c2 <- "Str"

	close(c1)
	for i := 0; i < 2; i++ {
		<-o
	}
}

func timeOutDemo() {
	c := make(chan bool)
	select {
	case v := <-c:
		fmt.Println(v)
	case <-time.After(2 * time.Second):
		fmt.Println("Time out ")
	}

}
