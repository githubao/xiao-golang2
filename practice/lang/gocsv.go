// csv实现
// author: baoqiang
// time: 2019/2/28 下午8:22
package lang

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func RunCsv() {
	filename := "/Users/baoqiang/1.csv"
	buf := new(bytes.Buffer)
	r2 := csv.NewWriter(buf)

	for i := 0; i < 10; i++ {
		s := make([]string,3)
		s[0] = fmt.Sprintf("user_id: %d", i)
		s[1] = fmt.Sprintf("name: %d", i)
		s[2] = fmt.Sprintf("depart: %d", i)
		r2.Write(s)
		r2.Flush()
	}

	fmt.Println(buf)
	fout,err := os.Create(filename)
	defer fout.Close()
	if err != nil{
		fmt.Println(filename,err)
		return
	}

	fout.WriteString(buf.String())

	//read file
	cntb,err := ioutil.ReadFile(filename)
	if err != nil{
		panic(err.Error())
	}

	r1 := csv.NewReader(strings.NewReader(string(cntb)))
	ss,_ := r1.ReadAll()

	for i := 0; i<len(ss); i++ {
		fmt.Println(ss[i])
	}

}
