// 任务
// author: baoqiang
// time: 2019/1/14 上午11:42
package xasync

import (
	"context"
)

type Tasker interface {
	Call(ctx context.Context, req interface{}) (interface{}, error)
}

type TaskFunc func(ctx context.Context, req interface{}) (interface{}, error)

func (t *TaskFunc) Call(data interface{}) (interface{}, error) {
	return t.Call(data)
}
