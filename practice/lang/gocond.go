// go cond
// author: baoqiang
// time: 2019/1/18 下午2:07
package lang

import (
	"bytes"
	"sync"
	"io"
	"fmt"
	"time"
)

type MyDataBucket struct {
	br     *bytes.Buffer
	gmutex *sync.RWMutex
	rcond  *sync.Cond
}

func NewDataBucket() *MyDataBucket {

	db := &MyDataBucket{
		br:     bytes.NewBuffer(make([]byte, 0)),
		gmutex: new(sync.RWMutex),
	}
	db.rcond = sync.NewCond(db.gmutex.RLocker())

	return db
}

func (db *MyDataBucket) Read(i int) {
	db.gmutex.RLock()
	defer db.gmutex.RUnlock()

	var data []byte
	var d byte
	var err error

	for {
		if d, err = db.br.ReadByte(); err != nil {
			if err == io.EOF {
				if string(data) != "" {
					fmt.Printf("reader-%d: %s\n", i, data)
				}
				db.rcond.Wait()
				data = data[:0]
				continue
			}
		}

		data = append(data, d)
	}

}

func (db *MyDataBucket) Put(d []byte) (int, error) {
	db.gmutex.Lock()
	defer db.gmutex.Unlock()

	n, err := db.br.Write(d)
	db.rcond.Broadcast()

	return n, err
}

func RunCond() {
	db := NewDataBucket()

	go db.Read(1)
	go db.Read(2)

	for i := 0; i < 1000; i++ {
		go func(i int) {
			d := fmt.Sprintf("data=%d", i)
			db.Put([]byte(d))
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
}
