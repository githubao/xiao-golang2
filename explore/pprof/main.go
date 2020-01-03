// run file
// author: baoqiang
// time: 2019-07-29 15:50
package pprof

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

func Servit() {
	go func() {
		//ip := "0.0.0.0:6000"
		ip := "localhost:6000"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
			os.Exit(1)
		}
	}()

	tick := time.Tick(time.Second / 100)

	var buf []byte
	for range tick {
		buf = append(buf, make([]byte, 1024*1024)...)
		fmt.Println("tick...")
	}
}

/*
go tool pprof -raw -seconds 30 http://localhost:6000/debug/pprof/profile
go tool pprof /Users/baoqiang/pprof/pprof.samples.cpu.001.pb.gz
*/
