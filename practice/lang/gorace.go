// 竞态条件
// author: baoqiang
// time: 2018/12/21 下午10:11
package lang

import (
	"log"
	"sync"
	"runtime"
)

var counter int
var wg sync.WaitGroup
var mutex sync.Mutex

func CounterAdd() {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		mutex.Lock()

		value := counter

		//退出当前的goroutine,重新放回队列中
		runtime.Gosched()

		value ++
		counter = value

		mutex.Unlock()
	}
}

func RunCounter() {
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go CounterAdd()
	}

	wg.Wait()

	log.Printf("count is: %v\n", counter)
}
