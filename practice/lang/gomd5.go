// go md5
// author: baoqiang
// time: 2019/2/28 下午9:55
package lang

import (
	"bytes"
	"fmt"
	"crypto/md5"
	"encoding/hex"
)

//a104523eb37b289bfb51b5ee75be883e
func RunMd5() {
	s1 := "something"
	s2 := "good"

	buf := bytes.NewBuffer([]byte(s1))
	buf.WriteString(s2)
	fmt.Println(buf.String())

	h := md5.New()
	h.Write([]byte(buf.String()))
	fmt.Println(hex.EncodeToString(h.Sum(nil)))
}
