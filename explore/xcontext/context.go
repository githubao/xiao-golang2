// context examples
// author: baoqiang
// time: 2019-07-30 17:27
package xcontext

import (
	"context"
	"fmt"
	"time"
)

// 其实调用cancel说白了就是往done的chan里面写值，
// parent知道cancel方法因为chan广播的方式所有的ctx.Done的goroutine都能读取到done的值
func ExampleWithCancel() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)

		n := 1

		go func() {
			for {
				select {
				case <-ctx.Done():
					// safe quit goroutine
					return
				case dst <- n:
					n += 1
				}
			}
		}()

		return dst
	}

	// ctx instance
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// read from gen
	for n := range gen(ctx) {
		fmt.Println(n)

		if n == 5 {
			break
		}

	}
}

// context deadline exceeded
// time.AfterFunc过了timeout的时间就会用cancelCtx关闭close chan，所以ctx.Done()就能读取到相应的值了
func ExampleWithTimeout() {
	timeout := 50 * time.Millisecond

	// ctx instance
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		// do time consumer task
		fmt.Println("over slept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

// ctx传递多个goroutine使用的全局共享变量的信息
func ExampleWithValue() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Printf("found value: %v, %v \n", k, v)
			return
		}
		fmt.Println("not found: ", k)
	}

	// init a key
	k1 := favContextKey("language")
	k2 := favContextKey("color")

	// ctx instance
	ctx := context.WithValue(context.Background(), k1, "Go")

	// read one
	f(ctx, k1)
	f(ctx, k2)

	fmt.Println(ctx)
}
