// 锁机制
// author: baoqiang
// time: 2019-05-06 17:04
package gogo

import (
	"sync/atomic"
)

type XMutex struct {
	key  int32
	sema int32
}

func xadd(val *int32, delta int32) (new int32) {
	for {
		v := *val
		if atomic.CompareAndSwapInt32(val, v, v+delta) {
			return v + delta
		}
	}
	panic("unreached")
}

func (m *XMutex) Lock() {
	if xadd(&m.key, 1) == 1 {
		return
	}
	Semacquire(&m.sema)
}

func (m *XMutex) Unlock() {
	if xadd(&m.key, -1) == 0 {
		return
	}
	Semrelease(&m.sema)
}

func Semacquire(sema *int32) {
}

func Semrelease(sema *int32) {
}
