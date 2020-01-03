// go 定时任务
// author: baoqiang
// time: 2019/2/28 下午9:05
package lang

import (
	"fmt"
	"time"
)

func F(){
	fmt.Println(time.Now())
}

// 每天凌晨开始任务
func RunTimer(f func()){

	go func() {
		f()
		now := time.Now()

		//next := now.Add(time.Hour * 24)
		next := now.Add(time.Minute * 1)
		next = time.Date(next.Year(),next.Month(),next.Day(),0,0,0,0,next.Location())
		t := time.NewTimer(next.Sub(time.Now()))
		<- t.C
	}()
}