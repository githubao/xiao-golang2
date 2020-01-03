// 并发读写struct引起的panic行为
// author: baoqiang
// time: 2019/3/5 下午4:03
package gogo

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
)

type Student struct {
	Sid  int64
	Name string
}

var s = Student{}

func RunStructRace() {
	var wg = new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		//wg.Add(1) 不能写在里面，要不然还没来得及add，main就运行完成了
		defer wg.Done()

		for range [10000]int{} {
			//fmt.Println(1)
			changeName1(&s)
			//time.Sleep(time.Duration(rand.Float64()) * time.Millisecond * 100)
			fmt.Println(time.Duration(rand.Float64()) * time.Millisecond * 100)
		}
	}()

	go func() {
		defer wg.Done()

		for range [10000]int{} {
			//fmt.Println(2)
			changeName2(&s)
			time.Sleep(time.Duration(rand.Float64()* 100) * time.Millisecond)
			fmt.Println(time.Duration(rand.Float64()* 100) * time.Millisecond)
		}
	}()

	wg.Wait()

	//time.Sleep(3*time.Second)

	fmt.Println(s)

}

func changeName1(s *Student) {
	s.Name = "aaa"
	fmt.Println(s.Name)
}

func changeName2(s *Student) {
	s.Name = "bbb"
	fmt.Println(s.Name)
}
