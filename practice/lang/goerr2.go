// go语言的错误处理机制
// author: baoqiang
// time: 2018/12/24 下午9:06
package lang

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var ErrHttpConnect = errors.New("http connect err")
var ErrInvalidParam = fmt.Errorf("invalid param err")

func Recover() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("err is : %v", err)
			log.Print("nothing happend")
		}
	}()
	panic("divided by zero")
}

// 不要在循环里面使用defer，如果需要的话，在循环里面包一个函数
func DeferInCycle() {
	for i := 0; i < 5; i++ {
		func() {
			f, err := os.Open("/path/to/file")
			if err != nil {
				log.Printf("err open file: %v", err)
			}
			defer f.Close()
		}()
	}
}

// 不要返回error的指针
func NoNilErr() *error {
	var p error = nil
	//if true{
	//	p = errors.New("some thing happens")
	//}
	return &p
}

func RunNoNilErr() {
	if err := NoNilErr(); err != nil {
		log.Printf("got err: %v", err)
	}
}
