package main

//import "fmt"
import "unsafe"

func main() {
	//First()
	//Data()
	//Ptr()
	//Cnt()
	//Increase()
	CalBit()
}


func CalBit()  {
	A := 0
	B := 1
	println(A&B)
	println(A|B)
	println(A^B)
	println(B<<4)
	println(15>>2)

	/*
	priority:
	^ !
	* / % << >> &
	+ - |
	== != < <= >= >
	<-
	&&
	||
	 */

	/*
	C language priority:
	() [] -> .
	! ~ ++ -- + - * & sizeof
	* / %
	+ -
	<< >>
	< <= >= >
	== !=
	&
	^
	|
	&&
	||
	?:
	= += -+ *= /= &= ^= != <<= >>=
	,
	 */



}

func Increase(){
	const (
		a = iota
		b
		c
		d="haha"
		e
		f=100
		g
		h=iota
		i
	)
	println(a,b,c,d,e,f,g,h,i)

	const(
		a1 = 1<<iota
		a2 = 3<<iota
		a3
		a4
	)
	println(a1,a2,a3,a4)
}

//const
func Cnt(){
	const (
		UNKNOWN = 0
		FEMALE = 1
		MALE = 2
	)
	var s = [2]string{"1","e"}
	println(unsafe.Sizeof(FEMALE))
	println(len(s))
	println(cap(s))

}

func Ptr(){
	var i int = 4
	var p *int= &i
	println(i,p)
}


func Data() {
	var flag bool = true
	var i int64 = 1
	var f float64 = 1.030
	var c complex64 = 1+2i
	var s string = "er"
	var _,b = 1,2
	/*
	pointer [] struc union func
	slice interface map channel
	int=uint uintptr
	 */
	println(flag,i,f,c,s,b)

	var a1 int = 10
	var a2 = 11
	a3 := 12
	println(a1,a2,a3)

	const A string = "tr"
}

func First() {
	println("hello")
}