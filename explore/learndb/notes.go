// 学到的一些笔记
// author: baoqiang
// time: 2019/1/21 下午2:55
package learndb

import "fmt"

func Noteit() {
	notes := []string{
		"NamedArg添加了不可导出的变量_Named_Fields_Required，它是一个结构体类型，是为了外部使用初始化的时候必须要指定匿名字段的值进行初始化",
		"IsolationLevel实现的是数据库的隔离性对应的隔离级别的枚举常量，通常有可串行化，可重复读，提交读，未提交读等",
		"数据库select出来的可能的数据类型有int64,float64,bool,[]byte,string,time.Time,nil",
		"Out结构体用于sql的存储过程，如果In为true那么Dest就是一个inout的变量，作为输入值之后然后把结果值替换掉",
		"",
		"",
		"",
	}

	fmt.Print(notes)
}
