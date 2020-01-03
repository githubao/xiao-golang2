// 文件读写服务
// author: baoqiang
// time: 2018/12/21 下午9:19
package lang

import (
	"log"
	"os"
	"path"
	"bufio"
	"io/ioutil"
)

func ReadLine(path string) {
	f, err := os.Open(path)
	HandlerError(err)

	defer f.Close()

	//r := bufio.NewReader(f)
	//datas, err := r.ReadString('\n')
	//HandlerError(err)
	//
	//log.Printf("read data: %s", datas)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		log.Printf("data: %s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		HandlerError(err)
	}

}

func ReadFile(path string) {
	f, err := os.Open(path)
	HandlerError(err)

	defer f.Close()

	buf := make([]byte, 1024^2)
	//buf := make([]byte, 32)
	n, err := f.Read(buf)

	log.Printf("read data: %s, cnt: %d", string(buf[:n]), n)
}

func WriteFile(path string, data string) {
	fw, err := os.Create(path)
	HandlerError(err)

	defer fw.Close()

	n, err := fw.WriteString(data)
	HandlerError(err)

	log.Printf("write cnt: %d\n", n)
}

// 多种读取文件的方式
func ReadAll(path string) {
	b1, _ := ioutil.ReadFile(path)

	// 上面的readfile是第二种方式，可以buffer分块读取

	f, err := os.Open(path)
	HandlerError(err)
	b3, _ := ioutil.ReadAll(f)

	log.Printf("b1: %s\n", string(b1))
	log.Printf("b3: %s\n", string(b3))
}

func RunReadWrite() {
	filename := path.Join(os.Getenv("HOME"), "Desktop", "1.txt")
	data := "abc\n你好\n3.14\n"

	WriteFile(filename, data)
	//ReadFile(filename)
	ReadAll(filename)
}

func RunReadLine() {
	filename := path.Join(os.Getenv("HOME"), "Desktop", "1.txt")
	ReadLine(filename)
}
