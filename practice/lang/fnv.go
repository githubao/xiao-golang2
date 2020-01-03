// 随机数 rand32
// author: baoqiang
// time: 2019/2/1 下午2:42
package lang

import _ "hash/fnv"


func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

// 最大值
func fnv32Max(key string, max uint32)uint32{
	hash := fnv32(key)
	return hash % max
}

//固定位数，高(32-length)位与低的(length)位做异或
func fnv32Len(key string, length uint32)uint32{
	hash := fnv32(key)
	return (hash >> length) ^ (hash & (1<<length-1))
}


const fnv =
`
32 bit FNV_prime = 2^24 + 2^8 + 0x93 = 16777619

64 bit FNV_prime = 2^40 + 2^8 + 0xb3 = 1099511628211

128 bit FNV_prime = 2^88 + 2^8 + 0x3b = 309485009821345068724781371

256 bit FNV_prime = 2^168 + 2^8 + 0x63 = 374144419156711147060143317175368453031918731002211

512 bit FNV_prime = 2^344 + 2^8 + 0x57 = 
35835915874844867368919076489095108449946327955754392558399825615420669938882575
126094039892345713852759

1024 bit FNV_prime = 2^680 + 2^8 + 0x8d = 
50164565101131186554345988110352789550307653454047907443030175238311120551081474
51509157692220295382716162651878526895249385292291816524375083746691371804094271
873160484737966720260389217684476157468082573

`