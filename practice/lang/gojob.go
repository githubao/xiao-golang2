// 线程池
// author: baoqiang
// time: 2018/12/26 下午7:29
package lang

import "log"

type Job interface {
	Do() error
}

type JobChan chan Job
type WorkChan chan JobChan // 限定go routine的数目

var (
	JobQueue   JobChan
	WorkerPool WorkChan
)

type Worker struct {
	JobChannel JobChan
	quit       chan bool
}

func (w *Worker) Start() {
	go func() {
		for {
			WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				if err := job.Do(); err != nil {
					log.Printf("excute job failed with err: %v", err)
				}
			case <-w.quit:
				return
			}
		}
	}()
}

const MaxWorkerPoolSize = 5

//分发器
type Dispatcher struct {
	Workers []*Worker
	quit    chan bool
}

func (d *Dispatcher) Run() {
	for i := 0; i < MaxWorkerPoolSize; i++ {
		worker := &Worker{
			JobChannel: make(chan Job),
			quit:       make(chan bool),
		}
		d.Workers = append(d.Workers, worker)
		worker.Start()
	}

	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				jobChan := <-WorkerPool
				jobChan <- job
			}(job)
		case <-d.quit:
			return
		}

	}

}
