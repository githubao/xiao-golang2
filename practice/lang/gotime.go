// 时间相关
// author: baoqiang
// time: 2018/12/24 下午7:38
package lang

import (
	"log"
	"time"
)

const Layout = "2006-01-02 15:04:05"

// 统一用utc时间解析
func CompareTime() {
	current := time.Now()

	timeStr := "2018-12-24 19:39:08"
	utc, _ := time.Parse(Layout, timeStr)
	local, _ := time.ParseInLocation(Layout, timeStr, time.Local)

	log.Printf("%v\n%v\n%v\n", current.Unix(), utc.Unix(), local.Unix())

}
