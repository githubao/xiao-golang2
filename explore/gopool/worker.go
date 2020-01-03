// 具体的干活的人，一直从jobqueue里面取任务
// author: baoqiang
// time: 2019/3/3 下午3:37
package gopool

type Worker struct {
	JobQueue chan Job
}

func NewWorker() Worker {
	return Worker{JobQueue: make(chan Job)}
}

func (w Worker) Run(wq chan chan Job) {
	go func() {
		for{
			// 把这个队列写进workpool的工作队列中去
			wq <- w.JobQueue
			select {
			case job := <- w.JobQueue:
				job.Do()

			}
		}
	}()
}
