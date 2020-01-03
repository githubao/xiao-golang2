// test render
// author: baoqiang
// time: 2019-04-19 11:38
package packer

import (
	"encoding/json"
	"fmt"
	"github.com/luci/go-render/render"
)

type customType int
type testStruct struct {
	S string
	V *map[string]int
	I interface{}
}

func RunRender() {
	a := testStruct{
		S: "hello",
		V: &map[string]int{"foo": 0, "bar": 1},
		I: customType(42),
	}

	fmt.Println("Render test:")
	fmt.Printf("fmt.Printf:    %#v\n", a)
	fmt.Printf("render.Render: %s\n", render.Render(a))
	fmt.Printf("json format: %s\n", DumpJson(a))
}

func DumpJson(obj interface{}) string {
	bdata, _ := json.Marshal(obj)
	return string(bdata)
}
