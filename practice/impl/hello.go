//
// author: baoqiang
// time: 2018/7/22 下午1:07
package impl

import (
	"encoding/json"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
	"testing"
	"unsafe"
	"io/ioutil"
)

func testIoCloser() {
	s := "hello"
	r := strings.NewReader(s)
	c := ioutil.NopCloser(r)
	data, _ := ioutil.ReadAll(c)
	fmt.Println(string(data))
}

func testString() {
	b := []byte{65}
	fmt.Println(string(b))
	fmt.Println(*(*string)(unsafe.Pointer(&b)))
}

func TestAbc(t *testing.T) {
	fmt.Println("aaa")
}

func test9() {
	a := true
	if a == true {
		fmt.Println(a)
	}

	fmt.Printf("%s", "aaa")

}

func test8() {
	var a []int
	if len(a) == 0 {
		fmt.Println("a切片是nil")
	}

	//a[0] = 1
	a = append(a, 2, 3)
	fmt.Println(a)

	b := make([]int, 0)
	if len(b) == 0 {
		fmt.Println("b切片是nil")
	}

	b = append(b, 2, 3)
	fmt.Println(b)

	c := make([]int, 2, 3)
	c[0] = 1
	add(c)
	fmt.Println("c: ", c)

	fmt.Println(c[2:3])

	//d := append(c, 1, 2)
	//fmt.Println(&c == &d)
}
func add(ints []int) {
	d := append(ints, 1, 2)
	fmt.Println(d)
}

func test7() {
	p := &Per{}
	for i := 0; i < 100; i++ {
		p.bh = i
		go func(p *Per) {
			fmt.Println(p)
		}(p)
	}

	println()
}

type Per struct {
	name string
	bh   int
}

func test6() {
	for i := 0; i < 5; i++ {
		go func() { fmt.Println(i) }() //错误的做法
		//go func(i int) { Hello(i) }(i) //正确的做法
	}

	time.Sleep(time.Second * 1)
}

func test5() {
	i := NewMyInt(4)
	i.String()
}

type MyInt struct {
	a int
}

func NewMyInt(i int) *MyInt {
	return &MyInt{
		a: i,
	}
}

func (i MyInt) String() string {
	//return fmt.Sprint(i)
	return string(i.a)
}

func test4() {
	s := "abc"
	fmt.Println(strings.Contains(s, "b"))
	fmt.Println(strings.ContainsAny(s, "b"))

	//Hello(3)
}

func test3() {
	db, err := sql.Open("mysql", "root:00@tcp(localhost:3306)/hello?charset=utf8")
	checkErr(err)

	defer db.Close()

	//obj := map[string]interface{}{
	//	"id":   "123\\",
	//	"name": "hello\t",
	//}

	obj := "{\"id\":\"123\\\",\"name\":\"你好\t\"}"
	//obj := `{"id":"123\\","name":"你好\t"}`

	//src, err := json.Marshal(obj)
	//checkErr(err)

	src := obj

	stmt, err := db.Prepare("INSERT INTO aaa(extra) VALUES (?)")
	checkErr(err)

	inserted, err := stmt.Exec(src)
	checkErr(err)

	affected, err := inserted.LastInsertId()
	checkErr(err)
	fmt.Printf("%d\n", affected)

	row := db.QueryRow("select * from aaa ORDER by id desc limit 1")
	var uid int
	var extra string
	err = row.Scan(&uid, &extra)
	checkErr(err)

	var res map[string]interface{}
	err = json.Unmarshal([]byte(extra), &res)
	checkErr(err)

	fmt.Println(uid, res)

}

func test2() {

	extraInfo := map[string]interface{}{
		"id":   "123\\",
		"name": "hello	",
	}

	res, err := json.Marshal(extraInfo)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(res))
	}

	//src := "{\"id\":\"123\\\\\",\"name\":\"你好\\t\"}"
	//res := []byte(src)

	var extraInfo2 map[string]interface{}

	if err2 := json.Unmarshal(res, &extraInfo2); err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(extraInfo2)
	}

}

func test1() {
	//var extraInfo model.WithdrawHistoryExtra

	var extraInfo map[string]interface{}

	src := "{\"activity_award\":0,\"ad_income\":9150,\"bank_branch_location\":\"广东广州\",\"bank_branch_name\":\"中国工商银行天河支行\",\"bank_name\":\"中国工商银行\t\",\"contract_list\":[{\"contract_article_amount\":2000000,\"contract_article_num\":200,\"contract_id\":\"ZN20171027008\\ ZN20180425023\",\"contract_reach_num\":190,\"contract_type\":1,\"real_amount\":1900000,\"sub_contract_type\":0}],\"pay_time\":1528992000,\"praise_income\":0,\"pre_contract_amount\":0,\"real_pay_amount_after_tax\":1603686,\"real_pay_amount_before_tax\":1909150,\"real_pay_amount_tax\":305464,\"video_income\":0}"
	//src := "{\"activity_award\":0,\"ad_income\":9150,\"bank_branch_location\":\"广东广州\",\"bank_branch_name\":\"中国工商银行天河支行\",\"bank_name\":\"中国工商银行\",\"contract_list\":[{\"contract_article_amount\":2000000,\"contract_article_num\":200,\"contract_id\":\"ZN20171027008 ZN20180425023\",\"contract_reach_num\":190,\"contract_type\":1,\"real_amount\":1900000,\"sub_contract_type\":0}],\"pay_time\":1528992000,\"praise_income\":0,\"pre_contract_amount\":0,\"real_pay_amount_after_tax\":1603686,\"real_pay_amount_before_tax\":1909150,\"real_pay_amount_tax\":305464,\"video_income\":0}"
	//src := "{\"activity_award\":0}"
	bsrc := []byte(src)

	if err := json.Unmarshal(bsrc, &extraInfo); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(extraInfo)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
