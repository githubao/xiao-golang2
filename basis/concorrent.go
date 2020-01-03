package main

import (
	"fmt"
	"math"
	"sync"
)

func sum(id int) {
	var x int64
	for i := 0; i < math.MaxUint16; i++ {
		x += int64(i)
	}
	fmt.Printf("id: %d, sum: %d\n", id, x)
}

func RoutineRun() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func(id int) {
			defer wg.Done()
			sum(id)
		}(i)
	}

	wg.Wait()
}

func ChannelRun() {
	data := make(chan int)
	exit := make(chan bool)

	go func() {
		for d := range data {
			fmt.Println(d)
		}

		fmt.Println("recv over.")
		exit <- true

	}()

	data <- 1
	data <- 2
	data <- 4

	close(data)

	fmt.Println("send over.")
	<-exit

}
