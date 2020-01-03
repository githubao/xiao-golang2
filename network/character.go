// 字符编码相关的问题
// author: baoqiang
// time: 2019-08-26 16:33
package network

import (
	"fmt"
	"unicode/utf16"
)

// 常见的ascii字符
// 10 \n
// 13 \r
// 32 space
// 43(2B) +
// 49 1
// 65 A
// 97 a

// ISO 8859-1也被称为Latin-1, 包含256个字符

// unicode实现了码点，统一表示世界上常用语言的常用字符

// utf-8是对unicode的一种实现，在golang中，rune是int32的一个别名

// 文本第一个字符的BOM，大端是0xFFFE,小端是0xFEFF
// 大端，数据的低字节放在内存的高位地址中

func Character() {
	//runeLen()
	runUtf16()
}

func runeLen() {
	var s = "你好,"
	fmt.Printf("[%v] rune length: %v\n", s, len([]rune(s)))
	fmt.Printf("[%v] byte length: %v\n", s, len(s))
}

func runUtf16() {
	// utf-8
	var s = "你好,"
	// rune to uint16
	ints := utf16.Encode([]rune(s))
	//bytes
	runes := utf16.Decode(ints)

	fmt.Printf("ints: %v\n", ints)
	fmt.Printf("runes: %v\n", runes)
}
