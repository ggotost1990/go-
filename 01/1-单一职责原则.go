package main

import "fmt"

type ClothesWork struct {
}

type ClothesShop struct {
}

func (cw *ClothesWork) OnWork() {
	fmt.Println("工作时穿的衣服...")
}

func (cs *ClothesShop) OnShop() {
	fmt.Println("逛街时穿的衣服...")
}

func main() {
	// 工作时
	cw := ClothesWork{}
	cw.OnWork()
	// 逛街时
	cs := ClothesShop{}
	cs.OnShop()
}
