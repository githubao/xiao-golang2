// 资源连接池
// author: baoqiang
// time: 2019/3/3 下午2:56
package impl

type Pool struct {
	chch chan chan []byte
	back chan chan []byte
	exit chan bool
}

func NewPool(count int) *Pool {
	p := new(Pool)
	p.back = make(chan chan []byte, count)
	p.chch = make(chan chan []byte, count)
	for i := 0; i < count; i++ {
		p.chch <- make(chan []byte, 1)
	}
	p.exit = make(chan bool)
	go p.run()
	return p
}

// 读一条看能不能读的到
func (p *Pool) Alloc() chan []byte {
	select {
	case ch := <-p.chch:
		return ch
	default:
		break
	}
	return make(chan []byte, 1)
}

// 写一条看能不能写的进去，写不进去就写进back
func (p *Pool) Free(ch chan []byte) {
	select {
	case p.chch <- ch:
		return
	default:
		p.back <- ch
	}
}

func (p *Pool) Close() {
	if p.exit != nil {
		close(p.exit)
		p.exit = nil
	}
}

// 如果back里面有的话，就读出来放进chs
// chch能写进去的话，就从chs里面取一个放进去
func (p *Pool) run() {
	var chs []chan []byte
	var chch chan chan []byte
	var next chan []byte

	for {
		select {
		case <-p.exit:
			return
		case ch := <-p.back:
			if chch == nil {
				chch = p.chch
				next = ch
			} else {
				chs = append(chs, ch)
			}
		case chch <- next:
			if len(chs) == 0 {
				chch = nil
				next = nil
			} else {
				next = chs[len(chs)-1]
				chs = chs[:len(chs)-1]
			}


		}
	}

}

