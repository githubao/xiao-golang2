// go协程竞争
// author: baoqiang
// time: 2018/12/26 下午2:06
package lang

import (
	"log"
	"time"
	"runtime"
)

func RunRace1() {
	// unreachable
	go log.Printf("hahaha")
}

func RunRace2() {
	go log.Printf("hahaha222")
	time.Sleep(1 * time.Second)
}

func RunRace3() {
	go log.Printf("hahaha333")

	// 让出竞态条件
	runtime.Gosched()
}

func RunRace4() {
	// 最大运行的cpu的数量
	runtime.GOMAXPROCS(1)

	// this is unreachable
	go log.Printf("lalala")

	for {

	}
}

func RunRace5() {
	// 最大运行的cpu的数量
	runtime.GOMAXPROCS(1)

	go log.Printf("lalala222")

	for {
		runtime.Gosched()
	}
}

func RunRace6() {
	// 最大运行的cpu的数量
	runtime.GOMAXPROCS(1)

	go log.Printf("lalala333")

	select {}
}

func RunRace7() {
	// 最大运行的cpu的数量
	runtime.GOMAXPROCS(1)

	done := false

	// may be unreachable
	go func() {
		log.Printf("papapa")
		done = true
	}()

	for {
		if done {
			break
		}
	}
}

func RunRace8() {
	// 最大运行的cpu的数量
	runtime.GOMAXPROCS(1)

	done := make(chan bool)

	// may be unreachable
	go func() {
		log.Printf("papapa222")
		done <- true
	}()

	for {
		if <-done {
			break
		}
	}
}

func RunRace() {
	//RunRace1()
	//RunRace2()
	//RunRace3()
	//RunRace4()
	//RunRace5()
	//RunRace6()
	//RunRace7()
	RunRace8()
}
