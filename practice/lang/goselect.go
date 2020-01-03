// go语言select
// author: baoqiang
// time: 2018/12/24 下午9:13
package lang

import (
	"log"
	"math/rand"
	"time"
)

// select case的chan，哪个chan先达到条件哪个就先执行

func eat() chan string {
	out := make(chan string)
	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		out <- "Mom call you eating"
		close(out)
	}()
	return out
}

func singleEat() {
	eatCh := eat()

	sleep := time.NewTimer(time.Second * 3)

	select {
	case s := <-eatCh:
		log.Printf("eat got: %s\n", s)
	case <-sleep.C:
		log.Printf("sleep time\n")
		//default:
		//	log.Printf("playing")
	}

}

func RunEat() {
	for i := 0; i < 10; i++ {
		singleEat()
	}
}
