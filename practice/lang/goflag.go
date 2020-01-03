// 使用flag解析命令行参数
// author: baoqiang
// time: 2019/2/28 下午2:30
package lang

import (
	"flag"
	"fmt"
)

// ./main -name xiao -age 16
func RunParseFlag() {
	name := flag.String("name", "", "student name")
	age := flag.Int("age", 0, "student age")

	flag.Parse()

	fmt.Println(*name)
	fmt.Println(*age)

}
