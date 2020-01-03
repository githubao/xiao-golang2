package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

// 接口
type Tester interface {
	Test()
	Eat()
}

// 实现了接口里面的所有方法就实现了接口
func (p *Person) Test() {
	fmt.Println("Test from Person")
}

func (p *Person) Eat() {
	fmt.Println("Eat from Person")
}

// 使用组合实现 类似继承 的功能
type Employee struct {
	Person
}

// 重写了同名方法就相当于override了
func (e *Employee) Eat() {
	fmt.Println("Eat from Employee")
}

//Eat from Employee
//Test from Person
//
//Eat from Employee
//Test from Person
//
//Eat from Person
//Test from Person
func GenericSample() {
	var e = new(Employee)
	e.Id = 1
	e.Name = "xiaoyu"

	e.Eat()
	e.Test()
	fmt.Println()

	var t Tester
	t = e
	t.Eat()
	t.Test()
	fmt.Println()

	var p Person
	p = e.Person
	p.Eat()
	p.Test()
	fmt.Println()
}
