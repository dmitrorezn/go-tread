package tread

import (
	"errors"
	"sync"
	"sync/atomic"
)

type ITread interface {
	Go(fn func()) error
	SpotAndWait()
}

func NewTread() *Tread {
	return &Tread{
		closed: atomic.Bool{},
		wg:     new(sync.WaitGroup),
	}
}

type Tread struct {
	closed atomic.Bool
	wg     *sync.WaitGroup
}

func (t *Tread) SpotAndWait() {
	t.closed.Store(true)
	t.wg.Wait()
}

var ErrTreadClose = errors.New("tread closed already")

func (t *Tread) Go(fn func()) error {
	if t.closed.Load() {
		return ErrTreadClose
	}
	t.wg.Add(1)
	go func() {
		fn()
		t.wg.Done()
	}()

	return nil
}
