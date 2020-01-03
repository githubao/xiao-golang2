// 一致性hash对应的实现
// author: baoqiang
// time: 2019-07-17 18:48
package gogo

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type UInt32Slice []uint32

func (s UInt32Slice) Len() int {
	return len(s)
}

func (s UInt32Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s UInt32Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int               // 复制因子
	keys     UInt32Slice       //已排序的节点哈希切片
	hashMap  map[uint32]string // 节点hash和key的map，键是hash值，值是节点key
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[uint32]string),
	}
	//默认使用CRC32算法
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}

	return m
}

func (m *Map) IsEmpty() bool {
	return len(m.keys) == 0
}

// 增加缓存节点
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		// 虚拟节点的复制因子
		for i := 0; i < m.replicas; i++ {
			hash := m.hash([]byte(strconv.Itoa(i) + key))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	// 排序
	sort.Sort(m.keys)
}

func (m *Map) Get(key string) string {
	if m.IsEmpty() {
		return ""
	}

	hash := m.hash([]byte(key))

	// 二分查找, 第一个大于hash的节点
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	if idx == len(m.keys) {
		idx = 0
	}

	return m.hashMap[m.keys[idx]]
}
