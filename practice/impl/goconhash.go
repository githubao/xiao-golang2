// 一致性hash的实现
// author: baoqiang
// time: 2019/3/1 下午5:09
package impl

import (
	"sync"
	"strconv"
	"sort"
	"github.com/klauspost/crc32"
	"fmt"
)

type SortKeys []uint32

func (sk SortKeys) Len() int           { return len(sk) }
func (sk SortKeys) Less(i, j int) bool { return sk[i] < sk[j] }
func (sk SortKeys) Swap(i, j int)      { sk[i], sk[j] = sk[j], sk[i] }

const DEFAULT_REPLICAS = 100

type HashRing struct {
	Nodes map[uint32]string
	Keys  SortKeys
	sync.RWMutex
}

func RunConHash() {
	nodes := []string{"a", "b", "c", "d", "e"}

	hr := NewHashRing(nodes)

	fmt.Println(hr.GetNode("a1"))
	fmt.Println(hr.GetNode("b6"))

}

func NewHashRing(nodes []string) *HashRing {
	hr := &HashRing{}

	if nodes == nil {
		return hr
	}

	hr.Nodes = make(map[uint32]string)
	hr.Keys = SortKeys{}

	for _, node := range nodes {
		// 添加所有虚拟节点的hash值
		for i := 0; i < DEFAULT_REPLICAS; i++ {
			str := node + strconv.Itoa(i)
			hr.Nodes[hr.hashStr(str)] = node
			hr.Keys = append(hr.Keys, hr.hashStr(str))
		}
	}

	sort.Sort(hr.Keys)

	return hr
}

func (hr *HashRing) hashStr(s string) uint32 {
	return crc32.ChecksumIEEE([]byte(s))
}

//寻找离key最近的那个节点
func (hr *HashRing) GetNode(key string) string {
	hr.RLock()
	defer hr.RUnlock()

	hash := hr.hashStr(key)
	i := hr.getPosition(hash)
	return hr.Nodes[hr.Keys[i]]
}

func (hr *HashRing) getPosition(hash uint32) int {
	i := sort.Search(len(hr.Keys), func(i int) bool {
		return hr.Keys[i] >= hash
	})

	if i < len(hr.Keys) {
		if i == len(hr.Keys)-1 {
			return 0
		} else {
			return i
		}
	} else {
		return len(hr.Keys) - 1
	}

}
