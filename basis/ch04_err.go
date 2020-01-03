package main

import (
	"fmt"
	"errors"
	"math"
)

func main() {
	runErr()
}

func runErr(){
	result,err := Sqrt(-1)
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}

}

type error interface {
	Error() string
}

func Sqrt(f float64) (float64,error){
	if f<0 {
		return 0,errors.New("math: squre root of negative num")
	}

	return math.Sqrt(f),nil
}
