// 自己写一个http server
// author: baoqiang
// time: 2018/12/21 下午10:25
package lang

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const addr = ":9527"

func RunServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "just another http server...")
	})

	srv := http.Server{
		Addr:    addr,
		Handler: http.DefaultServeMux,
	}

	var wg sync.WaitGroup
	exit := make(chan os.Signal)

	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-exit
		wg.Add(1)
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		HandlerError(err)

		wg.Done()
	}()

	err := srv.ListenAndServe()

	log.Printf("waiting for the remaining connections to finish...")

	wg.Wait()

	HandlerError(err)

	log.Print("gracefully shutdown the http server...")

}
