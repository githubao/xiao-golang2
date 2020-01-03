// 遍历dir，找到导入
// author: baoqiang
// time: 2019/2/13 下午2:26
package impl

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const outfile = "/Users/baoqiang/a.txt"

func walkDir(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("read dir failed: %v", err)
		os.Exit(1)
	}

	fw, err := os.Create(outfile)
	if err != nil {
		log.Fatalf("create fail failed: %v", err)
		os.Exit(1)
	}

	for _, file := range files {
		fullname := filepath.Join(dir, file.Name())

		filepath.Walk(fullname, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				//log.Printf("is dir: %v", path)
				return nil
			}

			if err != nil {
				//log.Printf("read file failed: %v", path)
				return nil
			}

			if !strings.HasSuffix(path, "go") {
				//log.Printf("does not has suffix go: %v", path)
				return nil
			}

			//if !strings.Contains(path, "walkdir") {
			//	return nil
			//}

			data, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatalf("read file failed: %v\n", err)
				return nil
			}

			// 处理导入文件
			log.Printf("process file: %v\n", path)

			start := false
			for _, line := range strings.Split(string(data), "\n") {
				if line == "import (" {
					start = true
					continue
				}
				if line == ")" {
					return nil
				}
				if start {
					trimed := strings.Replace(line, "\t", "", -1)
					trimed = strings.Replace(trimed, "\"", "", -1)
					if trimed == "" {
						continue
					}
					//fmt.Println(trimed)
					fw.WriteString(trimed)
					fw.WriteString("\n")
				}

			}

			return nil

		})
	}

	//关闭文件
	fw.Close()

}

func countImport() {
	dic := make(map[string]int)

	data, _ := ioutil.ReadFile(outfile)
	for _, line := range strings.Split(string(data), "\n") {
		dic[line] += 1
	}

	sortMap(dic)
}

type kv struct {
	Key   string
	Value int
}

func sortMap(dic map[string]int) {
	var ss []kv
	for k, v := range dic {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value //降序
	})

	for _, kv := range ss {
		fmt.Printf("%s\t%d\n", kv.Key, kv.Value)
	}

}
