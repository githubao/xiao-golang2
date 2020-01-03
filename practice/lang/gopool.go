// 协程池
// author: baoqiang
// time: 2018/12/21 下午1:44
package lang

import (
	"fmt"
	"log"
	"time"
)

func workerPool(n int, jobCh <-chan interface{}, retCh chan<- interface{}) {
	for i := 0; i < n; i++ {
		go worker(i, jobCh, retCh)
	}
}

func worker(id int, jobCh <-chan interface{}, retCh chan<- interface{}) {
	for job := range jobCh {
		ret := fmt.Sprintf("task %d processed job: %v", id, job)
		retCh <- ret
	}
}

func genJob(n int) <-chan interface{} {
	jobCh := make(chan interface{}, 200)
	for i := 0; i < n; i++ {
		jobCh <- i
	}
	close(jobCh)
	return jobCh
}

func RunPool() {
	jobCh := genJob(100)
	retCh := make(chan interface{}, 200)
	workerPool(5, jobCh, retCh)

	//wait to complete
	time.Sleep(time.Second * 1)
	close(retCh)

	for ret := range retCh {
		log.Printf("Got: %v\n", ret)
	}
}
