// go语言学习网的文章，每次随机10条数据
// author: baoqiang
// time: 2019/2/14 下午9:09
package impl

import (
	"math/rand"
	"fmt"
)

const urlFmt = "https://studygolang.com/articles/%d"

func shufSlice() {
	//num := 18221
	num := 10

	list := rand.Perm(num)
	for i := range list {
		list[i]++
	}

	fmt.Print(list)

}
