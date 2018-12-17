package main

import "fmt"

type Hello interface {
	out()
}

type MyInt int

func (myInt *MyInt) out() {
	fmt.Println("MyInt.out():", *myInt)
}

func foo(i MyInt) {
	fmt.Println("foo():", i)
}

func main() {
	var i MyInt
	i = 5
	i.out()
	foo(i)
	foo(1)

	var x int
	x = 100
	// foo(x) cannot use type int as type MyInt in a argument to foo
	foo(MyInt(x))
}
