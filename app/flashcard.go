// card卡片
// author: baoqiang
// time: 2019-08-27 15:34
package app

type FlashCard struct {
	Simplified string
	English    string
	Dictionary *Dictionary
}

type FlashCards struct {
	Name      string
	CardOrder string // random sequential
	ShowHalf  string // random_half english_half chinese_half
	Cards     []*FlashCard
}
