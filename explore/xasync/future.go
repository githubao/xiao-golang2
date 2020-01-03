// 结果类型
// author: baoqiang
// time: 2019/1/14 上午11:42
package xasync

import "context"

type TaskResult struct {
	Result interface{}
	Err    error
}

// 用于接收数据的返回结果
type Future struct {
	FutureResult chan TaskResult
	cancel       context.CancelFunc
}

func (f *Future) GetResults() []TaskResult {
	var results []TaskResult

	defer f.cancel()

	for res := range f.FutureResult {
		results = append(results, res)
	}

	return results
}
