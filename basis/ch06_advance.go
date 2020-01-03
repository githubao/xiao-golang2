// Created by BaoQiang at 2017/4/29 15:21

package main

import (
	"fmt"
	"reflect"
)

func main() {
	//StructDemo()
	//MethodDemo()
	//InterfaceDemo()
	ReflectDemo()
}

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func ReflectDemo() {
	u := User{1, "OK", 12}
	//Set(&u)
	//Info(u)
	MethodInvoke(&u)
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("not ptr")
	} else {
		v = v.Elem()
		//set value
	}

	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("Bad")
		return
	}

	if f.Kind() == reflect.String {
		f.SetString("new value")
	}
}

func MethodInvoke(o interface{}) {
	v := reflect.ValueOf(o)
	mv := v.MethodByName("Hello")
	args := [] reflect.Value{reflect.ValueOf("Joe")}
	mv.Call(args)

}

func (u User) Hello(name string) {
	fmt.Println("Hello ", name)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Name: ", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Print("Not struct")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fileds: ")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v =%v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}

type IntAlias int

type Person struct {
	name string
	age  int
}

type USB interface {
	Name() string
	Connect()
}

type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("Connect: ", pc.name)
}

func disConnected(usb interface{}) {
	// usb 是不是 PhoneConnector 类型
	if pc, ok := usb.(PhoneConnector); ok {
		fmt.Println("DisConnected: ", pc.name)
	} else {
		fmt.Println("Unknown Device")
	}

	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("DisConnected: ", v.name)
	default:
		fmt.Println("Unknown Device")

	}

}

func InterfaceDemo() {
	var a USB

	a = PhoneConnector{"pcConnector"}

	a.Connect()

	disConnected(a)

}

func Add() {
	var a IntAlias
	a.Increase(100)
	fmt.Println(a)
}

func (intAlias *IntAlias) Increase(num int) {
	*intAlias += IntAlias(num)
}

func StructDemo() {
	p := Person{"xiao", 23}
	fmt.Println(p)
}

func MethodDemo() {
	//p := Person{"xiao", 3}
	//p.Print()
	//
	//Person.Print(p)

	Add()
}

//方法没有重载的概念
//Receiver
func (p Person) Print() {
	fmt.Println(p)
}
