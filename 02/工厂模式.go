package main

import "fmt"

// 抽象层
type Fruit interface {
	Show()
}

type AbstrctFactory interface {
	CreateFactory() Fruit
}

// 工厂基础层
type AppleFactory struct {
	AbstrctFactory
}

type BananaFactory struct {
	AbstrctFactory
}

type PearFactory struct {
	AbstrctFactory
}

func (appleFac *AppleFactory) CreateFactory() Fruit {
	var fruit Fruit
	fruit = new(Apple)
	return fruit
}

func (bananaFac *BananaFactory) CreateFactory() Fruit {
	var fruit Fruit
	fruit = new(Banana)
	return fruit
}

func (pearFac *PearFactory) CreateFactory() Fruit {
	var fruit Fruit
	fruit = new(Pear)
	return fruit
}

// 水果基础层
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

//
func main() {
	// 业务层
	var factory AbstrctFactory
	factory = new(AppleFactory)
	var apple Fruit
	apple = factory.CreateFactory()
	apple.Show()

	//factory = new(BananaFactory)
	//banana := factory.CreateFactory()
	//banana.Show()
	//
	//factory = new(PearFactory)
	//pear := factory.CreateFactory()
	//pear.Show()

}
