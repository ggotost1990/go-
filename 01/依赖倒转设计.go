package main

import "fmt"

// 抽象层
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// 实现层
type Benz struct {
}

type Bmw struct {
}

func (b *Benz) Run() {
	fmt.Println("Benz is running...")
}

func (b *Bmw) Run() {
	fmt.Println("Bmw is running...")
}

type Zhangsan struct {
}

type Lisi struct {
}

func (z *Zhangsan) Drive(car Car) {
	fmt.Println("张三开车..")
	car.Run()
}

func (l *Lisi) Drive(car Car) {
	fmt.Println("李四开车")
	car.Run()
}

func main() {
	// 业务层
	var benz Car
	benz = new(Benz)
	var zhangsan Driver
	zhangsan = new(Zhangsan)
	zhangsan.Drive(benz)
	var li4 Driver
	li4 = new(Lisi)
	var bwm Car
	bwm = new(Bmw)
	li4.Drive(bwm)
}
