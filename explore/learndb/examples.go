// 一些测试相关
// author: baoqiang
// time: 2019/1/21 下午2:47
package learndb

import (
	"database/sql"
	"github.com/labstack/gommon/log"
)

func SqlName() {
	a := sql.NamedArg{Name: "age", Value: 18}
	//a := sql.NamedArg{struct{}{}, "age",18}
	log.Printf("sql arg: %v", a)
}
