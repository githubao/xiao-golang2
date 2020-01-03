// hash函数
// author: baoqiang
// time: 2019/3/1 下午2:57
package impl

import "fmt"

func HashCode(key string) int {
	var index = 0
	index = int(key[0])
	for k := 0; k < len(key); k++ {
		index += 11035152245 + int(key[k])
	}
	index >>= 27
	index &= 16 - 1
	return index
}

func RunHmap() {
	InitBuckets()
	AddKeyValue("a","hello world")
	AddKeyValue("abc","hello China")

	fmt.Println(GetValueByKey("a"))
	fmt.Println(GetValueByKey("abc"))
}

var buckets = make([]*Node, 16)

func InitBuckets() {
	for i := 0; i < 16; i++ {
		buckets[i] = CreateHead(KValue{"head", "start"})
	}
}

func AddKeyValue(key, value string) {
	var nIndex = HashCode(key)
	var headNode = buckets[nIndex]
	var tailNode = TailNode(headNode)
	AddNode(KValue{key, value}, tailNode)
}

func GetValueByKey(key string ) string {
	var nIndex = HashCode(key)
	var headNode = buckets[nIndex]
	var value = FindValueByKey(key,headNode)
	return value
}
