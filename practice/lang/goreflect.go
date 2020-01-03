// 写一个简单的反射代码 
// author: baoqiang
// time: 2019/2/25 下午5:50
package lang

import (
	"reflect"
	"fmt"
)

type Employee struct {
	Id           int64
	Name         string
	PersonalAddr string
}

func createQuery(o interface{}) {
	// 具体的类型信息
	t := reflect.TypeOf(o)

	// 具体的值的信息
	v := reflect.ValueOf(o)

	switch t.Kind() {
	case reflect.Struct:
		fmt.Printf("income type is: %v\n", t)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int64:
				fmt.Printf("idx:%d name is %s, value is %d\n", i, t.Field(i).Name, v.Field(i).Int())
			case reflect.String:
				fmt.Printf("idx:%d name is %s, value is %s\n", i, t.Field(i).Name, v.Field(i).String())
			}
		}
	}
}

func RunReflect() {
	e := Employee{
		Id:           123,
		Name:         "xiaobao",
		PersonalAddr: "海淀",
	}

	createQuery(e)
}
