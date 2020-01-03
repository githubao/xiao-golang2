// language explore
// author: baoqiang
// time: 2018/12/21 下午3:16
package lang

import (
	"log"
	"fmt"
	"time"
	"strconv"
)

// 使用 struct{} 实现无内存消耗的分配
// 比如 使用map实现set，发送无数据的signal的channel

func RunChanaaa() {
	//var readonly <-chan int = make(chan int)
	//var writeonly chan<- int = make(chan int)

	var readwrite = make(chan int)

	var exit = make(chan struct{})

	go func() {
		for data := range readwrite {
			fmt.Println(data)
		}

		exit <- struct{}{}
	}()

	readwrite <- 1
	readwrite <- 2
	readwrite <- 3

	close(readwrite)

	<-exit
}

func RunFnvHash() {
	//s := "爱😔t3"
	// "爱"是utf8长度是3(0-2),"😔"是unicode长度是4(3-6)
	// "t"长度1(7),"3"的长度是1(8)
	// 爱😔t3 对应的码点分别是29233，128532，116，51

	//for _, item := range []rune(s) {
	//	fmt.Println(item)
	//}

	//fmt.Println(fnv32(s))
	//fmt.Println(fnv32Max(s, 100))
	//fmt.Println(fnv32Len(s, 24))

	//爱 /U7231 0111 0010 0011 0001
	//char := "爱"
	////cp, _ := strconv.Atoi("爱")
	//cpStr := strconv.QuoteToASCII("爱")
	//cpStr2 := strings.Replace(cpStr, "\\u", "", 1)
	//cpStr3 := strings.Replace(cpStr2, "\"", "", 2)
	//cp, _ := strconv.ParseInt(cpStr3, 10, 64)
	//
	////b := "0111001000110001"
	////res, _ := strconv.ParseInt(b, 2, 32)
	//fmt.Printf("%s的二进制表示为: %b 码点为: %v\n", char, cp, cp)
	//
	//var c int64 = 65535
	////fmt.Println(strconv.FormatInt(c, 2))
	//fmt.Printf("%v的16,8,2进制表示: %x, %o, %b\n\n", c, c, c, c)

	//str := "世界"
	str := "世界"
	fmt.Printf("%v的unicode码：%v\n", str, []rune(str))
	fmt.Printf("%v的utf8码：%v\n", str, []byte(str))
	fmt.Printf("%v的unicode字符串表示：%v\n", str, strconv.QuoteToASCII(str))
	//fmt.Printf("%v的二进制表示：%032b\n", str, []rune(str)[0])
	fmt.Printf("%v的二进制表示：%016b\n", str, []rune(str)[0])
	fmt.Printf("\n")
}

func RunBinary() {
	//a := 0xFFFFFF
	a := 0xFF
	fmt.Println(a)

	length := 8
	b := 1<<(uint32(length)) - 1
	fmt.Println(b)
}

func RunMapInterface() {
	dic := make(map[string]interface{})

	var item1 string
	var item2 interface{}

	item1 = "item1"
	item2 = "item2"

	dic["key1"] = item1
	dic["key2"] = item2

	val1 := dic["key1"].(string)
	val2 := dic["key2"].(string)

	fmt.Println(val1)
	fmt.Println(val2)

}

func RunArray() {
	arr := [3]int{1, 2, 3}
	log.Printf("%T %+v\n", arr, arr)
}

//命名参数，需要注意函数内变量对其的修改
func Return() (a int) {
	//b := 2
	a = 1
	return
}

func Return2() (a int) {
	b := 2
	a = 1
	return b
}

func RunReturn() {
	log.Print(Return())
	log.Print(Return2())
}

var ticker = time.NewTicker(1 * time.Second)

func RunChan() {
	//read := make(<-chan int, 1)
	//write := make(chan<- int, 1)

	gogo := make(chan int, 1)

	go func() {
		fmt.Println("read")

		for {
			select {
			case data := <-gogo:
				fmt.Println(data)
			case <-ticker.C:
				return
			}
		}

	}()

	go func() {
		fmt.Println("write")
		gogo <- 3
		gogo <- 4
	}()

	time.Sleep(50 * time.Millisecond)
}
