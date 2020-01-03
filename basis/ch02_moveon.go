package main

import "fmt"

func main() {
	//Array()
	//MyPtr()
	//Struct()
	Slice()
}

func Slice() {
	slice := make([]int, 2, 5)
	PrintSlice(slice)
	slice = append(slice, 2)
	slice = append(slice, 3)
	PrintSlice(slice)
	number1 := slice[2:3]
	PrintSlice(number1)
	copy(slice,number1)
	PrintSlice(slice)

}

func PrintSlice(slice []int) {
	var j int;
	for j = 0; j < len(slice); j++ {
		fmt.Print(slice[j])
		fmt.Print(" ")
	}
	fmt.Println()
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
}

type Book struct {
	title   string
	author  string
	subject string
	book_id int
}

func Struct() {
	var book Book
	book.author = "xiao"
	book.book_id = 42

	fmt.Println(book.author, book.book_id)
}

func Explore() {
	s := "你好"
	println(s)
}

func MyPtr() {
	a := 10
	p := &a
	println(a, &a, p, *p)
}

func Array() {
	var a [10] float64 = [10]float64{}
	var b = [3]string{"a", "b", "cc"}

	var i, j int;

	for i = 0; i < len(a); i++ {
		println(a[i])
	}

	for j = 0; j < len(b); j++ {
		println(b[j])
	}
}



