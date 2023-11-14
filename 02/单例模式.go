package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var lock sync.Mutex

type singelton struct {
}

// 饿汉模式， 编译期间就会生成一个实例对象
//var instance *singelton = new(singelton)

var instance *singelton
var initialized uint32

func GetInstance() *singelton {
	// 饿汉模式， 只有没有实例的时候才会创建
	// 为了线程安全 增加互斥
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = new(singelton)
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

func (s *singelton) Something() {
	fmt.Println("something...")
}
func main() {
	s1 := GetInstance()
	s1.Something()
	s2 := GetInstance()
	s2.Something()
}
