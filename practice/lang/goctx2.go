// 通过ctx cancel防止内存泄漏
// author: baoqiang
// time: 2018/12/26 下午2:32
package lang

import (
	"context"
	"log"
)

func RunCtx() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				select {
				case <-ctx.Done():
					return
				default:
					ch <- i

				}
			}
		}()
		return ch
	}(ctx)

	for v := range ch {
		log.Printf("Got %v", v)
		if v == 5 {
			cancel()
			break
		}
	}
}
