// golang 实现一个自旋锁
// author: baoqiang
// time: 2019-05-07 11:20
package gogo

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type spinLock uint32

// 原来值是0，能够交换成功,跳出循环
// 不是0的话，交换不成功，runtime.Gosched()交出cpu的控制使用权
func (sl *spinLock) Lock() {
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		runtime.Gosched()
	}
}

func (sl *spinLock) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}

func NewSpinLock() sync.Locker {
	var lock spinLock
	return &lock
}

var globalCount int

func RunLock() {
	var wg sync.WaitGroup
	var lock = NewSpinLock()
	//var lock sync.Mutex

	for i := 0; i < 10000; i++ {
		wg.Add(1)

		go func() {
			lock.Lock()
			globalCount += 1
			lock.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("Got result: %v\n", globalCount)

}
