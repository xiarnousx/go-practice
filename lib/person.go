package lib

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func Ex3() {
	p1 := Person{Name: "James", Age: 45}
	p2 := Person{Name: "MoneyPenny", Age: 27}

	a := []Person{p1, p2}
	fmt.Println(a)
	// below we are casting slice of person to ByAge type
	sort.Sort(ByAge(a))
	fmt.Println(a)
}
