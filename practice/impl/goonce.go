// 实现once
// author: baoqiang
// time: 2019/3/3 上午11:28
package impl

import (
	"sync"
	"sync/atomic"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}

	o.m.Lock()
	defer o.m.Unlock()

	if o.done == 0 {
		f()
		atomic.StoreUint32(&o.done, 1)
	}

}

func (o *Once) Do2(f func()){
	if atomic.LoadUint32(&o.done) == 1{
		return
	}
	if atomic.CompareAndSwapUint32(&o.done,0,1){
		f()
	}
}