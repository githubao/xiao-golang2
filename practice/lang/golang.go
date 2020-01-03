// go语言
// author: baoqiang
// time: 2018/12/21 下午12:59
package lang

import "log"

func Info() {
	str := `
	1. 2007年9月开始设计，2009年11月正式发布
	2. 语言特点：编译，垃圾回收，并发	
	3. 应用领域：系统编程，web-server，存储架构
	4. 包注释使用/**/，注释后面空一行，注释使用大写开头
	`
	log.Printf("%s\n", str)
}

func Keywords() {
	keys := `
	1. fmt
	2. errors	
	3. strings
	4. strconv
	5. regexp
	6. encoding/json
	
	1. bool
	2. int,int8,int16,int32,int64
	3. uint,uint8(short),uint16,uint32,uint64
	4. float32,float64
	5. string
	6. complex64,complex128
	7. array

	1. slice
	2. map
	3. chan

	1. append
	2. close
	3. delete
	4. panic
	5. recover
	6. imag
	7. real
	8. make
	9. new
	10. cap
	11. copy
	12. len

	1. error
	type error interface {
		Error() string
	}

	`
	log.Printf("%s\n", keys)
}

func Keywords2() {
	keys := `
	1. 压缩
	2. 缓冲
	3. 字节
	4. 容器
	5. 上下文
	6. 加解密
	7. 数据库
	8. 字符串
	9. 编码
	10. 错误
	11. 日志
	12. 格式化IO
	13. 图形
	14. IO
	15. 数学
	16. 网络
	17. 操作系统
	18. 反射
	19. 运行时
	20. 锁
	21. 原语
	22. 测试
	23. 模板
	24. 排序
	25. 类型转换
	26. 时间
	27. 字符编码
	`
	log.Printf("%s\n", keys)
}

func Usage() {
	able := `
	1. 处理日志
	2. 数据打包
	3. 虚拟机处理
	4. 文件系统
	5. 分布式系统
	6. 数据库代理
	7. Web应用
	8. 内存数据库
	9. 云平台
	`
	log.Printf("%s\n", able)

}
