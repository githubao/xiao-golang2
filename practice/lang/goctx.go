// context相关
// author: baoqiang
// time: 2018/12/21 下午5:03
package lang

import (
	"context"
	"log"
	"time"
)

func someHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doSth(ctx)

	// do main work
	time.Sleep(time.Second * 5)
	cancel()
}

func timeoutHandler() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	go doSth(ctx)

	// do main work
	time.Sleep(time.Second * 5)
	cancel()
}

func doSth(ctx context.Context) {
	var i = 1
	for {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			log.Println("done")
		default:
			log.Printf("work has spend %d seconds", i)
		}
		i++
	}
}

func RunCtxCancel() {
	someHandler()
}

func RunCtxTimeout() {
	timeoutHandler()
}
