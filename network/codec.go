// 编解码
// author: baoqiang
// time: 2019-08-23 18:55
package network

import (
	"bytes"
	"encoding/asn1"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

func Codec() {
	//asn1Code()
	gobCode()
	jsonCode()
	base64Code()
}

func asn1Code() {
	var src = 234

	mdata, err := asn1.Marshal(src)
	checkError(err)

	var n int
	_, err = asn1.Unmarshal(mdata, &n)
	checkError(err)

	fmt.Printf("asn1 got_data: %v\n", n)
}

func jsonCode() {
	s := map[string]interface{}{
		"name": "xiao",
		"age":  18,
	}

	var buffer = new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)

	err := encoder.Encode(s)
	checkError(err)

	var out Student
	decoder := json.NewDecoder(buffer)
	err = decoder.Decode(&out)
	checkError(err)

	fmt.Printf("json got_data: %v\n", out)
}

func gobCode() {
	s := Student{Name: "xiao", Age: 18}

	var buffer = new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)

	err := encoder.Encode(s)
	checkError(err)

	//var out map[string]interface{} // un support
	var out Student
	decoder := gob.NewDecoder(buffer)
	err = decoder.Decode(&out)
	checkError(err)

	fmt.Printf("gob got_data: %v\n", out)
}

func base64Code() {
	s := Student{Name: "xiao", Age: 18}
	b, _ := json.Marshal(s)

	var buffer = new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.StdEncoding, buffer)

	_, err := encoder.Write(b)
	checkError(err)
	encoder.Close()

	//fmt.Printf("base64 buffer_data: %v\n", buffer)

	var out = make([]byte, 256)
	decoder := base64.NewDecoder(base64.StdEncoding, buffer)
	_, err = decoder.Read(out)
	checkError(err)

	fmt.Printf("base64 got_data: %v\n", out)
}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
