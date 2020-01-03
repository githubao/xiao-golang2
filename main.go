// main方法
// author: baoqiang
// time: 2018/7/27 下午5:50
package main

import (
	"fmt"
	"time"

	"github.com/githubao/xiao-golang2/practice/apollo"
)

func run() {
	//sample.RunLissajous()
	//fmt.Println("hello")
	//TestTime()
	//lang.RunParseFlag()
	//gopool.MultiRun()
	//gogo.RunSlice()
	//gogo.RunStructRace()
	//gogo.RunOcr()
	//packer.RunRender()
	//gogo.RunLock()
	//xcontext.ExampleWithCancel()
	//xcontext.ExampleWithTimeout()
	//xcontext.ExampleWithValue()
	//pprof.Servit()
	//uptr.ExampleUptr()
	//network.NetworkSample()
	//network.HttpGet()
	//network.SimpleTimeServer()
	//network.Codec()
	//network.RunServer()
	//network.RunClient()
	//network.Character()
	//network.Security2()
	//network.RunTempl()
	//network.RunRpcServer()
	//apollo.RunBloomFilter()
	apollo.RunLog()
}

func tmp() {
	var p *int
	p = new(int)
	fmt.Println(p)
	fmt.Println(*p)
}

func TestTime() {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	tm2 := tm1.AddDate(0, 0, 1)
	fmt.Println(tm2)
}

func main() {
	run()
	//tmp()
}
