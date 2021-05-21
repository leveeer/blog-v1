package common

import "sync"

// goroutine pool
type GoroutinePool struct {
	c  chan struct{}
	wg *sync.WaitGroup
}

// 采用有缓冲channel实现,当channel满的时候阻塞
func NewCoroutinePool(maxSize int) *GoroutinePool {
	if maxSize <= 0 {
		panic("invalid maxSize")
	}
	return &GoroutinePool{
		c:  make(chan struct{}, maxSize),
		wg: new(sync.WaitGroup),
	}
}

//add goroutine
func (g *GoroutinePool) Add(delta int) {
	g.wg.Add(delta)
	for i := 0; i < delta; i++ {
		g.c <- struct{}{}
	}
}

//donne goroutine
func (g *GoroutinePool) Done() {
	<-g.c
	g.wg.Done()
}

//wait goroutine
func (g *GoroutinePool) Wait(done chan struct{}) {
	go func() {
		g.wg.Wait()
		close(done)
	}()
}

func (g *GoroutinePool) AddGoroutine(handle func()) {
	g.Add(1)
	go func() {
		handle()
		g.Done()
	}()

}
