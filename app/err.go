// err process
// author: baoqiang
// time: 2019-08-27 15:49
package app

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(-1)
	}
}
