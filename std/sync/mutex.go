package sync

import (
	"sync"
	"time"
)

const cost = time.Microsecond

type RW interface {
	Read()
	Write()
}

type Lock struct {
	count int
	mu    sync.Mutex
}

func (l *Lock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost)
	l.mu.Unlock()
}

func (l *Lock) Read() {
	l.mu.Lock()
	time.Sleep(cost)
	_ = l.count
	l.mu.Unlock()
}

type RWLock struct {
	count int
	mu    sync.RWMutex
}

func (l *RWLock) Write() {
	l.mu.RLock()
	l.count++
	time.Sleep(cost)
	l.mu.RUnlock()
}

func (l *RWLock) Read() {
	l.mu.RLock()
	time.Sleep(cost)
	_ = l.count
	l.mu.RUnlock()
}
