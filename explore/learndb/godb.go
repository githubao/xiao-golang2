// db查询相关
// author: baoqiang
// time: 2019/3/1 上午11:17
package learndb

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/pkg/errors"
)

var CREATE_TABLE = "CREATE TABLE `test_info`(" +
	"`id` int(11) NOT NULL AUTO_INCREMENT," +
	"`info` varchar(255) DEFAULT NULL," +
	"PRIMARY KEY (`id`)" +
	")ENGINE=InnoDB"

// 1 1551412864
// 2 1551412885
// 3 1551413216
// 1 <nil>
// map[id:1 info:1551412864]
func RunSql() {
	fmt.Println(CREATE_TABLE)

	db, err := sql.Open("mysql", "root:00@tcp(127.0.0.1:3306)/hello?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, _ := fetchRows(db, "SELECT * FROM test_info")
	for _, v := range *rows {
		fmt.Println(v["id"], v["info"])
	}

	fmt.Println(insert(db, "INSERT INTO test_info (info) VALUES (?) ", fmt.Sprintf("%d", time.Now().Unix())))

	row, _ := fetchRow(db, "SELECT * FROM test_info where id = ?", 1)
	fmt.Println(*row)

}

func insert(db *sql.DB, sqlstr string, args ...interface{}) (int64, error) {
	stmtIns, err := db.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(args...)
	if err != nil {
		panic(err.Error())
	}

	return result.RowsAffected()
}

func fetchRow(db *sql.DB, sqlstr string, args ...interface{}) (*map[string]string, error) {
	ret, err := fetchRows(db, sqlstr, args...)
	if err == nil && ret != nil && len(*ret) > 0 {
		return &(*ret)[0], nil
	}
	return nil, errors.New("no rows")
}

func fetchRows(db *sql.DB, sqlstr string, args ...interface{}) (*[]map[string]string, error) {
	stmtOut, err := db.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	ret := make([]map[string]string, 0)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		vmap := make(map[string]string, len(scanArgs))

		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			vmap[columns[i]] = value
		}

		ret = append(ret, vmap)
	}

	return &ret, nil
}
