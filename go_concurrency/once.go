package go_concurrency

import (
	"fmt"
	"sync"
	"time"
)

type MuOnce struct {
	sync.RWMutex
	sync.Once
	mtime time.Time
	vals  []string
}

func (m *MuOnce) refresh() {
	m.Lock()
	defer m.Unlock()
	m.Once = sync.Once{}
	m.mtime = time.Now()
	m.vals = []string{m.mtime.String()}
}

func (m *MuOnce) strings() []string {
	now := time.Now()
	m.RLock()
	if now.After(m.mtime) {
		defer m.Do(m.refresh)
	}
	vals := m.vals
	m.RUnlock()
	return vals
}

func OnceMain() {
	fmt.Println("Hello, playground")
	m := new(MuOnce)
	fmt.Println(m.strings())
	fmt.Println(m.strings())
}
