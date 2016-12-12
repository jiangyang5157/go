package singleton

import "testing"

func Test_singleton(t *testing.T) {
	GetInstance().SayHello()
	GetInstance().SayHello()
	GetInstance().SayHello()
	GetInstance().SayHello()
	go GetInstance().SayHello()
	go GetInstance().SayHello()
	go GetInstance().SayHello()
	go GetInstance().SayHello()
}