package main

import (
	"fmt"
)

func main() {
	bin := person{10, "bin"}
	fmt.Println(bin)
	modify(&bin)
	fmt.Println(bin)
}

func modify(p *person) {
	p.age = p.age + 10
}

type person struct {
	age  int
	name string
}
