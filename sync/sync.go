package sync

import "sync"

type Counter struct {
	lock  sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.lock.Lock()
	c.value++
	c.lock.Unlock()
}

func (c *Counter) Value() int {
	return c.value
}
