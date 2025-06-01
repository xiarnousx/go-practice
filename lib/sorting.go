package lib

import (
	"fmt"
	"sort"
)

func Ex1() {
	s := []int{22, 4, 44, 11, 1, 0}
	sort.Ints(s)
	fmt.Println(s)
}

func Ex2() {
	ss := []string{"a", "k", "m", "l"}
	sort.Strings(ss)
	fmt.Println(ss)
}
