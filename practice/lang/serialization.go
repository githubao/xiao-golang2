// 序列化
// author: baoqiang
// time: 2018/12/21 下午12:34
package lang

import (
	"log"
	"encoding/json"
	"os"
	"strings"
)

type Student struct {
	Id      int64  `json:"id"`
	StuName string `json:"stu_name"`
	Age     int    `json:"age"`
}

func Serialize() {
	stu := Student{
		Id:      123,
		StuName: "xiaobao",
		Age:     18,
	}

	bdata, err := json.Marshal(stu)
	HandlerError(err)

	log.Printf("%s\n", string(bdata))
}

func DeSerialize() {
	str := `{"id":123,"stu_name":"xiaobao","age":18}`
	var stu Student

	err := json.Unmarshal([]byte(str), &stu)
	HandlerError(err)

	log.Printf("%+v\n", stu)

}

func Encode() {
	stu := Student{
		Id:      123,
		StuName: "xiaobao",
		Age:     18,
	}

	encoder := json.NewEncoder(os.Stdout)
	err := encoder.Encode(stu)
	HandlerError(err)
}

func Decode() {
	str := `{"id":123,"stu_name":"xiaobao","age":18}`
	var stu Student

	decoder := json.NewDecoder(strings.NewReader(str))
	err := decoder.Decode(&stu)
	HandlerError(err)
	log.Printf("%+v\n", stu)
}
