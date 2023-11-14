package main

import "fmt"

type Benz struct {
}

type BWM struct {
}

type Zhangsan struct {
}

func (z *Zhangsan) run(benz *Benz) {
	fmt.Println("张三开奔驰")
}

func main() {
	z := Zhangsan{}
	benz := Benz{}
	z.run(&benz)
}
