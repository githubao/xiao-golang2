// sort
// author: baoqiang
// time: 2019/2/28 下午7:57
package lang

import "fmt"

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func RunSwap() {
	ints := []int{9, 2, 1, 4, 3, 8}
    Sort(si(ints))
	fmt.Println(ints)
}

type si []int
type ss []string

func (s si) Len() int           { return len(s) }
func (s si) Less(i, j int) bool { return s[j] < s[i] }
func (s si) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s ss) Len() int           { return len(s) }
func (s ss) Less(i, j int) bool { return s[j] < s[i] }
func (s ss) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func Sort(x Sorter) {
	for i := 0; i < x.Len()-1; i++ {
		for j := i + 1; j < x.Len(); j++ {
			if x.Less(i, j) {
				x.Swap(i, j)
			}
		}
	}
}
