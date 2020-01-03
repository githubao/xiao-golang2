// 运行实例
// author: baoqiang
// time: 2019/3/3 下午3:38
package gopool

import (
	"fmt"
	"time"
	"runtime"
	"math/rand"
)

type Score struct {
	Num int
}

func (s *Score) Do() {
	//fmt.Println("num: ", s.Num)
	time.Sleep(time.Duration(rand.Float32()) * time.Second)
}

func MultiRun() {
	workerNum := 1000

	p := NewWorkerPool(workerNum)
	p.Run()

	taskNum := 1000 * 1000

	go func() {
		for i := 1; i <= taskNum; i++ {
			sc := &Score{Num: i}
			p.JobQueue <- sc
		}
	}()

	for {
		fmt.Println("goroutine count: ", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}

}
