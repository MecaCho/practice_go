package singleton

import (
	"fmt"
	"sync"
)

type Manager struct{}

var m *Manager
var once sync.Once

func GetInstance() *Manager {
	once.Do(func() {
		m = &Manager{}
	})
	return m
}

func (p Manager) Manage() {
	fmt.Println("manage...")
}
