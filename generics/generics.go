package generics

import (
	"fmt"
)

// type set
// ~ underling types
type numerics interface {
	~int | ~float64
}

// type constraints
func add[T numerics](a, b T) T {
	return a + b
}

// alias
type someAlias int

func Ex1() {
	var n someAlias = 42
	fmt.Println(add(n, 2))
	fmt.Println(add(1, 3.2))
}

// package constraint
