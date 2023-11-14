package main

import (
	"fmt"
	"sync"
)

type singelton struct {
}

// 饿汉模式， 编译期间就会生成一个实例对象
//var instance *singelton = new(singelton)

var instance *singelton
var once sync.Once

func GetInstance() *singelton {
	// 饿汉模式， 只有没有实例的时候才会创建
	// 为了线程安全 增加互斥
	once.Do(func() {
		instance = new(singelton)
	})

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
