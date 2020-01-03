// 自己实现一个日志"框架"
// author: baoqiang
// time: 2020/1/3 7:49 下午
package apollo

import (
	"fmt"
	"os"
)

/*
1. Log接口使用interface{}，便于扩展
2. 负数命名可以末尾加s，也可以加List
3. 实例命名使用类的所有首字母，保证含义清晰
4. 提供类的New方法，返回该类的指针

*/

// user interface
func RunLog() {
	logger := createLogger()
	logger.Log("hello")
}

func createLogger() *Logger {
	g := &Logger{}

	cw := &ConsulWriter{}
	g.registerWriter(cw)

	cwd, _ := os.Getwd()
	fw := &FileWriter{filename: cwd + "/test.log"}
	g.registerWriter(fw)

	return g
}

// log struct
type Logger struct {
	writerList []LogWriter
}

func (g *Logger) registerWriter(w LogWriter) {
	g.writerList = append(g.writerList, w)
}

func (g *Logger) Log(s string) {
	for _, w := range g.writerList {
		w.Write(s)
	}
}

// log writer interface
type LogWriter interface {
	Write(interface{}) error
}

// impl Log writer
type ConsulWriter struct {
}

func (w *ConsulWriter) Write(data interface{}) error {
	s := fmt.Sprintf("%v", data)
	_, err := os.Stdout.Write([]byte(s))
	return err
}

type FileWriter struct {
	filename string
}

func (w *FileWriter) Write(data interface{}) error {
	s := fmt.Sprintf("%v", data)

	fw, err := os.Create(w.filename)
	if err != nil {
		fmt.Printf("fileWriter err: %v\n", err)
		return err
	}
	defer fw.Close()

	_, err = fw.Write([]byte(s))
	return err
}
