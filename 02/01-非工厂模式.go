package main

import "fmt"

type Fruit struct {
}

func NewFruit(name string) *Fruit {
	fruit := new(Fruit)
	if name == "apple" {

	} else if name == "pear" {

	} else if name == "banana" {

	}
	return fruit
}

func (f *Fruit) Show(name string) {
	if name == "apple" {
		fmt.Println("apple")
	} else if name == "pear" {
		fmt.Println("pear")
	} else if name == "banana" {
		fmt.Println("banana")
	}
}

func main() {
	apple := NewFruit("apple")
	apple.Show("apple")

}
