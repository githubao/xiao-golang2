// 测试用例的编写
// author: baoqiang
// time: 2019/2/28 下午3:41
package lang

import "testing"

func Add(x, y int) int {
	return x + y
}

// go test -run="TestAdd"
// go test -v
func TestAdd(t *testing.T) {
	if c := Add(2, 3); c != 5 {
		t.Fatalf("failed")
	} else {
		t.Log("good")
	}
}

// go test -v -test.bench="."
// go test -v -test.bench="BenchmarkAdd"
func BenchmarkAdd(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Add(3, 4)
	}
}
