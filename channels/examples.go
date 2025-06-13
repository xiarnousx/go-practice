package channels

import "fmt"

func Ex19() {
	c := make(chan int)
	// below blocks
	// c <-1
	// below unbloack
	go func() { c <- 1 }()
	fmt.Println(<-c)
}

func Ex20() {
	c := factorial(4)
	for n := range c {
		fmt.Print(n, " ")
	}
	fmt.Println("")
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
