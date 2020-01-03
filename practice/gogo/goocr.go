// 文本识别
// author: baoqiang
// time: 2019/3/14 下午10:13
package gogo

import (
	"fmt"

	"github.com/otiai10/gosseract"
)

func RunOcr() {
	filename := "/Users/baoqiang/Downloads/3.png"

	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(filename)
	text, _ := client.Text()
	fmt.Println(text)
	// Hello, World!
}
