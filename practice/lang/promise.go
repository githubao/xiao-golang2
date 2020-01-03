// 封装wait group表现出promise的效果
// author: baoqiang
// time: 2018/12/21 下午10:43
package lang

import (
	"log"
	"sync"
)

type Promise struct {
	wg sync.WaitGroup
}

func (p *Promise) Add(f func()){
	p.wg.Add(1)

	go func() {
		defer p.wg.Done()
		f()
	}()
}

func (p *Promise) End(){
	p.wg.Wait()
}

func t(i int) func(){
	return func() {
		log.Printf("%d\n", i)
	}
}

func RunPromise(){
	var promise Promise

	for i := 0; i < 3; i++ {
		promise.Add(t(i))
	}

	promise.End()

}
