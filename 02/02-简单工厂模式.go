package main

import "fmt"

type Fruit interface {
	Show()
}

type Apple struct {
	Fruit
}

func (apple *Apple) Show() {
	fmt.Println("我是Apple")
}

type Banana struct {
	Fruit
}

func (banana *Banana) Show() {
	fmt.Println("我是Banana")
}

type Pear struct {
	Fruit
}

func (pear *Pear) Show() {
	fmt.Println("我是Pear")
}

// 工厂
type Factory struct {
}

func (f *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit
	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	}
	return fruit
}
func main() {
	factory := Factory{}

	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()
}
