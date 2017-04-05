package singleton

import (
	"fmt"
	"sync"
)

var instance *Manager

var once sync.Once

func GetInstance() *Manager {
	once.Do(func() {
		fmt.Println("Create new instance")
		instance = &Manager{}
	})
	return instance
}

type Manager struct{}

func (m Manager) SayHello() {
	fmt.Println("Hello")
}
