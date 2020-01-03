// golang exec command
// author: baoqiang
// time: 2019/2/28 下午2:55
package lang

import (
	"os/exec"
	"fmt"
)

const pathname = "/dev/ttys000"

func RunCommand(){
	lsof := exec.Command("lsof", pathname)
	wc := exec.Command("wc","-l")

	lsofOut,_ := lsof.StdoutPipe()

	// start cmd
	lsof.Start()

	wc.Stdin = lsofOut

	out,_ := wc.Output()

	fmt.Println(string(out))
}
