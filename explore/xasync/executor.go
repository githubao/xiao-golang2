// 跑任务
// author: baoqiang
// time: 2019/1/14 下午12:02
package xasync

import (
	"time"
	"context"
	"sync"
	"runtime"
	"fmt"
)

type TaskExecutor struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewTaskExecutor(timeout int) *TaskExecutor {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Millisecond)
	return &TaskExecutor{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (e *TaskExecutor) Submit(req interface{}, tasks ...Tasker) *Future {
	var wg sync.WaitGroup

	results := make(chan TaskResult, len(tasks))

	for _, task := range tasks {
		wg.Add(1)

		go func(t Tasker) {
			select {
			case results <- func() (result TaskResult) {
				defer func() {
					if err := recover(); err != nil {
						result = TaskResult{nil, ErrorStack(err)}
						return
					}
				}()

				resp, err := t.Call(e.ctx, req)
				result = TaskResult{resp, err}
				return
			}():
			case e.ctx.Done():
				results <- TaskResult{nil, e.ctx.Err()}
			}

		}(task)

		wg.Done()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return &Future{
		FutureResult: results,
		cancel:       e.cancel,
	}
}

func ErrorStack(err interface{}) error {
	stack := make([]byte, 4096)
	runtime.Stack(stack, false)
	return fmt.Errorf("err: %v\nstack: %s", err, stack)

}
