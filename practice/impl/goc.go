// go调用c
// author: baoqiang
// time: 2019/3/1 下午8:10
package impl

/*
# include <goc.h>
*/
import "C"
import "fmt"


func RunC() {
	f := C.initFunc(C.forty_two)
	fmt.Println(int(C.bridge_int_func(f)))
}
