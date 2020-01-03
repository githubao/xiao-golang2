// go语言中json的坑
// author: baoqiang
// time: 2018/12/24 下午7:57
package lang

import (
	"encoding/json"
	"log"
	"reflect"
	"bytes"
)

// 1. 数字格式使用json.Number对应的类型
// 2. 格式化时间的时候，自定义format格式
// 3. 需要注意合理的零值会被json的omitempty的tag忽略！使用中划线表示无论如何都会过滤该字段

import (
	"time"
	"fmt"
)

type Person struct {
	Name   string   `json:"name,omitempty"`
	Age    int      `json:"age,omitempty"`
	Weight float64  `json:"weight,omitempty"`
	Date   jsonTime `json:"date,omitempty"`
}

type Person2 struct {
	Name   string   `json:"-"`
	Age    int      `json:"-"`
	Weight float64  `json:"-"`
	Date   jsonTime `json:"-"`
}

func EnDe() {
	p := Person{
		Name:   "xiaobao",
		Age:    18,
		Weight: 66.6,
	}

	bdata, _ := json.Marshal(p)

	var s map[string]interface{}

	_ = json.Unmarshal(bdata, &s)

	log.Printf("%v\n", s)
	log.Printf("%v\n", reflect.TypeOf(s["Age"]).Name())
}

func EnDe2() {
	p := Person{
		Name:   "xiaobao",
		Age:    18,
		Weight: 66.6,
	}

	bdata, _ := json.Marshal(p)

	var s map[string]interface{}

	de := json.NewDecoder(bytes.NewReader(bdata))
	de.UseNumber()
	de.Decode(&s)

	log.Printf("%v\n", s)
	log.Printf("%v\n", reflect.TypeOf(s["Age"]).Name())
}

type jsonTime time.Time

func (jt jsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(jt).Format(Layout))
	return []byte(stamp), nil
}

func RunTimeJson() {
	p := Person{
		Name:   "xiaobao",
		Age:    18,
		Weight: 66.6,
		Date:   jsonTime(time.Now()),
	}

	bdata, _ := json.Marshal(p)
	fmt.Printf("%v\n", string(bdata))
}

func OmitEmpty() {
	p := Person2{
		Age: 17,
	}

	bdata, _ := json.Marshal(p)
	fmt.Printf("%v\n", string(bdata))
}
