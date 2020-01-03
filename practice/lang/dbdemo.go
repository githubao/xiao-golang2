// mysql连接
// author: baoqiang
// time: 2018/12/21 下午1:07
package lang

import (
	"database/sql"
	"log"
)

var db *sql.DB

//TODO https://studygolang.com/articles/17093
func Init() {
	var err error
	db, err = sql.Open("mysql", "root:123@/student?charset=utf8")
	HandlerError(err)
}

func Close() {
	db.Close()
}

func SelectOne() {
	row := db.QueryRow("select id,stu_name,age from student_info where id = 3;")
	log.Printf("%v\n", row)
}
