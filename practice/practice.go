package practice

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
)

type Person struct {
	Name string
	Age  int
}

func Ex1() {

	p1 := Person{Name: "James", Age: 42}
	p2 := Person{Name: "MoneyPenny", Age: 27}

	people := []Person{p1, p2}

	bs, err := json.Marshal(people)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(bs))
}

func Ex2() []Person {
	ps := `
	[{"Name":"James","Age":42},{"Name":"MoneyPenny","Age":27}]
	`
	pepole := []Person{}

	err := json.Unmarshal([]byte(ps), &pepole)

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(pepole)

	return pepole
}

func Ex3() {
	people := Ex2()

	err := json.NewEncoder(os.Stdout).Encode(people)

	if err != nil {
		fmt.Println(err.Error())
	}
}

type Circle struct {
	radius int64
}

type Square struct {
	side int64
}

type Shape interface {
	area() float64
}

func (c Circle) area() float64 {
	return math.Pow(float64(c.radius), 2) * math.Pi
}

func (s *Square) area() float64 {
	return float64(s.side) * float64(s.side)
}

func info(s Shape) {
	fmt.Printf("The area is %v of %T \n", s.area(), s)
}

func Ex5() {
	// circle reciever is a type sow we can accpet
	// a pointer type or a  type
	c := Circle{radius: 2}
	// square reciever is a pointer so we only accept
	// pointer type
	s := &Square{side: 4}

	info(c)

	info(s)
}

func Ex6() {
	// print os cpu and system info
	fmt.Println("Arch", runtime.GOARCH)
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("CPU cores:", runtime.NumCPU())
	fmt.Println("-----")
}
