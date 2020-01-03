// 布隆过滤器
// author: baoqiang
// time: 2019-09-03 14:24
package apollo

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math"

	"github.com/spaolacci/murmur3"
	"github.com/willf/bitset"
)

// test func
func RunBloomFilter() {
	n := uint(10000)
	p := 0.01

	f := NewWithEstimates(n, p)

	// m: 95851, k:7
	fmt.Printf("m: %d, k:%d\n", f.m, f.k)

	s1 := []byte("aaa")
	s2 := []byte("bbb")

	// add
	f.Add(s1)

	// test
	fmt.Println(f.Test(s1))
	fmt.Println(f.Test(s2))
}

type BloomFilter struct {
	m uint
	k uint
	b *bitset.BitSet
}

// new func
func NewBloomFilter(m uint, k uint) *BloomFilter {
	return &BloomFilter{max(1, m), max(1, k), bitset.New(m)}
}

func FromBloomFilter(data []uint64, k uint) *BloomFilter {
	m := uint(len(data) * 64)
	return &BloomFilter{m, k, bitset.From(data)}
}

func NewWithEstimates(n uint, fp float64) *BloomFilter {
	m, k := EstimateParameters(n, fp)
	return NewBloomFilter(m, k)
}

// get fields
func (f *BloomFilter) Cap() uint {
	return f.m
}

func (f *BloomFilter) K() uint {
	return f.k
}

// impl methods

// add method
func (f *BloomFilter) Add(data []byte) *BloomFilter {
	h := baseHashes(data)
	for i := uint(0); i < f.k; i++ {
		f.b.Set(f.location(h, i))
	}
	return f
}

func (f *BloomFilter) AddString(data string) *BloomFilter {
	return f.Add([]byte(data))
}

// test method
// 不在是的确不在，在的话是"可能在"
func (f *BloomFilter) Test(data []byte) bool {
	h := baseHashes(data)
	for i := uint(0); i < f.k; i++ {
		if !f.b.Test(f.location(h, i)) {
			return false
		}
	}
	return true
}

func (f *BloomFilter) TestString(data string) bool {
	return f.Test([]byte(data))
}

func (f *BloomFilter) TestLocations(locs []uint64) bool {
	for i := 0; i < len(locs); i++ {
		if !f.b.Test(uint(locs[i] % uint64(f.m))) {
			return false
		}
	}
	return true
}

// test&add method
func (f *BloomFilter) TestAndAdd(data []byte) bool {
	present := false

	h := baseHashes(data)
	for i := uint(0); i < f.k; i++ {
		l := f.location(h, i)
		if !f.b.Test(l) {
			present = false
		}
		f.b.Set(l)
	}

	return present
}

func (f *BloomFilter) TestAndAddString(data string) bool {
	return f.TestAndAdd([]byte(data))
}

func (f *BloomFilter) ClearAll() *BloomFilter {
	f.b.ClearAll()
	return f
}

// copy method
func (f *BloomFilter) Copy() *BloomFilter {
	fc := NewBloomFilter(f.m, f.k)
	fc.Merge(f)
	return fc
}

func (f *BloomFilter) Merge(g *BloomFilter) error {
	if f.m != g.m {
		return fmt.Errorf("m's don't match: %d != %d", f.m, g.m)
	}
	if f.k != g.k {
		return fmt.Errorf("k's don't match: %d != %d", f.k, g.k)
	}

	f.b.InPlaceUnion(g.b)

	return nil
}

// 统计当前过滤器的错误率
func (f *BloomFilter) EstimateFalsePositiveRate(n uint) (fpRate float64) {
	rounds := uint32(100000)
	f.ClearAll()

	n1 := make([]byte, 4)
	for i := uint32(0); i < uint32(n); i++ {
		binary.BigEndian.PutUint32(n1, i)
		f.Add(n1)
	}

	fp := 0

	for i := uint32(0); i < rounds; i++ {
		binary.BigEndian.PutUint32(n1, i+uint32(n)+1)
		if f.Test(n1) {
			fp++
		}
	}

	// rating
	fpRate = float64(fp) / (float64(rounds))

	f.ClearAll()

	return
}

// equal method
func (f *BloomFilter) Equal(g *BloomFilter) bool {
	return f.m == g.m && f.k == g.k && f.b.Equal(g.b)
}

// private methods
func (f *BloomFilter) location(h [4]uint64, i uint) uint {
	return uint(location(h, i) % uint64(f.m))
}

// help func
func max(x, y uint) uint {
	if x > y {
		return x
	}
	return y
}

func baseHashes(data []byte) [4]uint64 {
	a1 := []byte{1}
	hasher := murmur3.New128()

	hasher.Write(data)
	v1, v2 := hasher.Sum128()

	hasher.Write(a1)
	v3, v4 := hasher.Sum128()

	return [4]uint64{
		v1, v2, v3, v4,
	}
}

func location(h [4]uint64, i uint) uint64 {
	ii := uint64(i)
	return h[ii%2] + ii*h[2+(((ii+(ii%2))%4)/2)]
}

// get the data locations representation
func Locations(data []byte, k uint) []uint64 {
	locs := make([]uint64, k)

	h := baseHashes(data)

	for i := uint(0); i < k; i++ {
		locs[i] = location(h, i)
	}
	return locs
}

// 根据准确率和数据大小，估计容量m和hash函数的个数k
func EstimateParameters(n uint, p float64) (m uint, k uint) {
	m = uint(math.Ceil(-1 * float64(n) * math.Log(p) / math.Pow(math.Log(2), 2)))
	k = uint(math.Ceil(math.Log(2) * float64(m) / float64(n)))
	return
}

// json marshal
type bloomFilterJSON struct {
	M uint           `json:"m"`
	K uint           `json:"k"`
	B *bitset.BitSet `json:"b"`
}

func (f *BloomFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(bloomFilterJSON{f.m, f.k, f.b})
}

func (f *BloomFilter) unmarshalJSON(data []byte) error {
	var j bloomFilterJSON

	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}
	f.m = j.M
	f.k = j.K
	f.b = j.B

	return nil
}

// read & write
func (f *BloomFilter) WriteTo(stream io.Writer) (int64, error) {
	err := binary.Write(stream, binary.BigEndian, uint64(f.m))
	if err != nil {
		return 0, err
	}

	err = binary.Write(stream, binary.BigEndian, uint64(f.k))
	if err != nil {
		return 0, err
	}

	numBytes, err := f.b.WriteTo(stream)
	return numBytes + int64(2*binary.Size(uint64(0))), err
}

func (f *BloomFilter) ReadFrom(stream io.Reader) (int64, error) {
	var m, k uint64

	err := binary.Read(stream, binary.BigEndian, &m)
	if err != nil {
		return 0, err
	}

	err = binary.Read(stream, binary.BigEndian, &k)
	if err != nil {
		return 0, err
	}

	b := &bitset.BitSet{}
	numBytes, err := b.ReadFrom(stream)
	if err != nil {
		return 0, err
	}

	// set value
	f.m = uint(m)
	f.k = uint(k)
	f.b = b

	return numBytes + int64(2*binary.Size(uint64(0))), err
}

// gob impl
func (f *BloomFilter) GobEncode() ([]byte, error) {
	var buf bytes.Buffer
	_, err := f.WriteTo(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (f *BloomFilter) GobDecode(data []byte) error {
	buf := bytes.NewBuffer(data)
	_, err := f.ReadFrom(buf)

	return err
}
