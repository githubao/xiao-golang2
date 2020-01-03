package main

import "fmt"

func Greeting() {
	fmt.Println("hello go!")
}

/**
0
1
10
10
*/
func ConstRun() {
	const (
		a = iota
		b
		c = 10
		d
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	const (
		B        = iota
		KB int64 = 1 << (iota * 10)
		MB
		GB
	)
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
}

func StringConvert() {
	s := "hello"
	c := []rune(s)
	c[1] = 'a'
	s2 := string(c)

	fmt.Println(c)
	fmt.Println(s2)
}

func StringMulti() {
	s1 := "first" +
		"second"

	s2 := `first
		second`

	fmt.Println(s1)
	fmt.Println(s2)
}

func ComplexRun() {
	c := 1 + 1i
	fmt.Println(c)
}

func MultiAssignment() {
	i, data := 0, []int{1, 2, 3}
	i, data[i] = 2, 100
	fmt.Println(data)
}

func MakeSlice() {
	s := make([]int, 3)
	s[1] = 3
	fmt.Println(s)
}
