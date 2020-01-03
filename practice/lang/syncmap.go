// map的并发锁实现
// author: baoqiang
// time: 2018/12/21 下午1:23
package lang

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

var n = 20

// fatal error: concurrent map writes
type M struct {
	lock sync.RWMutex
	Map  map[string]string
}

// Store
func (m *M) Set(key, value string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.Map[key] = value
}

// Load
func (m *M) Get(key string) (value string, ok bool) {
	value, ok = m.Map[key]
	return value, ok
}

func RunMap() {
	c := &M{
		Map: make(map[string]string),
	}
	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(num int) {
			c.Set(strconv.Itoa(num), fmt.Sprintf("my map: %d", num))
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Printf("job done")

	printMap(c)
}

func printMap(m *M) {
	for i := 0; i < n; i++ {
		v, _ := m.Get(strconv.Itoa(i))
		log.Printf("%d -> %v", i, v)
	}
}

func RunSyncMap(){
	c := &sync.Map{}
	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(num int) {
			c.Store(strconv.Itoa(num), fmt.Sprintf("sync map: %d", num))
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Printf("job done")

	c.Range(func(key, value interface{}) bool {
		log.Printf("%v -> %v", key, value)
		return true
	})

}