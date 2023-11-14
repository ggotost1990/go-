package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Null struct {
}

func add(nums ...int) {
	fmt.Println(nums)
}

func main() {

	add(1, 2, 3)
	//p := Person{"ggotost", 123}
	//fmt.Println(p)
	//fmt.Println(p.name, p.age)
	//
	//p.name = "jack"
	//fmt.Println(p.name)
	//pp := p
	//pp.age = 22
	//fmt.Println(pp.age)
	//ppp := &p
	//ppp.age = 32
	//fmt.Println(p.age)
	//colors := make([]string, 2)
	//colors[0] = "yellow"
	//colors[1] = "blue"
	//fmt.Println(colors, len(colors), cap(colors))
	//colors = append(colors, "red")
	//fmt.Println(colors, len(colors), cap(colors))
	//p2 := new(Person)
	//p2.name = "hell"
	//fmt.Println(p2)
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	names2 := names[0:2]
	fmt.Println(names, names2)
	names2[0] = "ggotost"
	fmt.Println(names, names2)
	//fmt.Println(names)
	//a := names[0:2]
	//b := names[1:3]
	//
	//fmt.Println(a, b)
	//a[0] = "xxx"
	//fmt.Println(names, a, b)
	//s1 := []int{1, 2, 3, 4}
	//s2 := make([]int, len(s1), (cap(s1)+1)*2)
	//copy(s2, s1)
	//fmt.Println(s2, len(s2), cap(s2))
}
