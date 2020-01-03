package main

import (
	"fmt"
)



func main() {
	//Range()
	//Map()
	//Other()
	PhoneRun()
}


type Phone interface {
	call()
}

type NokiaPhone struct {

}

type IPhone struct {

}

func (iPhone IPhone) call(){
	fmt.Println("i am iphone")
}

func (nokiaPhone NokiaPhone) call(){
	fmt.Println("i am nokia")
}

func PhoneRun(){
	var phone Phone
	phone = new(IPhone)
	phone.call()

	phone = new(NokiaPhone)
	phone.call()
}


func Other(){
	var sum int = 17
	var cnt int = 5
	var mean float64 = float64(sum) / float64(cnt)
	fmt.Println(mean)

}

func Range() {
	nums := []int{2, 3, 5}
	sum := 0
	for i, num := range nums {
		sum += num
		fmt.Println("index: ", i)
	}

	fmt.Println("sum: ",sum)

	kvs := map[string]string{"a":"apple","b":"banana"}
	for k,v := range kvs{
		fmt.Printf("%s -> %s\n",k,v)
	}

	for i,c := range "go"{
		fmt.Println(i,c)
	}
}

func Map(){
	var countryCapitalMap map[string]string = make(map[string]string)

	countryCapitalMap["France"]="Paris"
	countryCapitalMap["Italy"]="Rome"

	capital,ok := countryCapitalMap["Italy"]
	if(ok){
		fmt.Println("Italy capital is: "+capital)
	}else {
		fmt.Println("Italy capital is not present")
	}

	delete(countryCapitalMap,"Italy")

	capital,ok = countryCapitalMap["Italy"]
	if(ok){
		fmt.Println("Italy capital is: "+capital)
	}else {
		fmt.Println("Italy capital is not present")
	}
}


