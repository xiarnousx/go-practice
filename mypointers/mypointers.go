package mypointers

import (
	"fmt"
)

func Ex1() {
	// pointers and addresses
	x := 42
	fmt.Println(&x)
	fmt.Printf("%v\t%T\n", &x, &x)

	// using the astrix on it's own
	y := &x
	fmt.Println(y)
	// give us the value of that address
	fmt.Println(*y)
	fmt.Println(*&x)
	*y = 43
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*y)
}

func Ex2() {
	// asks for a pointer for an int i.e address to an int
	b := func(n *int) {
		//*n = 43
		c := 55
		// can we change the address of n?
		// I was expecting a to equal 55
		// n = &c
		*n = c
	}
	a := 42
	fmt.Println(a)
	// when calling b we give an address to an int
	b(&a)
	fmt.Println(a)
}

func Ex3() {
	// pointer semantics
	fmt.Println("Pointer Semantics")
	subOne := func(n *int) int {
		*n -= 1
		return *n
	}

	b := 1
	fmt.Println(b)
	fmt.Println(subOne(&b))
	fmt.Println(b)

	// value semantics
	fmt.Println("Value Semantics")
	addOne := func(n int) int {
		return n + 1
	}
	a := 1
	fmt.Println(addOne(a))
	fmt.Println(a)
}

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My Name is", d.first)
}

func (d dog) run() {
	fmt.Println("My name is ", d.first, " and running")
}

type doggy interface {
	walk()
	run()
}

func dogStyle(d doggy) {
	d.walk()
	d.run()
}

func Ex4() {

	d1 := dog{first: "Henrry"}
	d2 := dog{first: "Todd"}
	dogStyle(d1)
	fmt.Println(d1.first)

	fmt.Println("D2")
	dogStyle(d2)
}
