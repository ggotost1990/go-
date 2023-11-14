package main

import "fmt"

type Cat struct {
}

func (c *Cat) Eat() {
	fmt.Println("小猫吃饭")
}

type Catc struct {
}

func (cc *Catc) Eat(c *Cat) {
	// 通过传递父类的指针来调用父类的方法进行组合
	c.Eat()
}

func main() {
	//cat := new(Cat)
	//cat.Eat()
	cc := &Catc{}
	cc.Eat(&Cat{})

}
