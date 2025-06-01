package channels

import (
	"fmt"
	"sync"
)

func Ex13() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func Ex14() {
	c := make(chan int, 1)
	c <- 42
	fmt.Println(<-c)
}

func Ex15() {
	gen := func() <-chan int {
		c := make(chan int)
		go func() {
			for i := range 5 {
				c <- i
			}
			close(c)
		}()

		return c
	}

	receive := func(c <-chan int) {
		for v := range c {
			fmt.Println(v)
		}
	}

	c := gen()
	receive(c)
}

func Ex16() {
	c := make(chan int, 3)
	c <- 42
	c <- 43
	c <- 44
	for {
		select {
		case v := <-c:
			fmt.Println("hello", v)
			if v == 44 {
				close(c)
				return
			}
		}
	}
}

func Ex17() {
	c := make(chan int, 1)
	c <- 42
	if v, ok := <-c; ok {
		fmt.Println(v)
	}
}

func Ex18() {
	c := make(chan int)
	gen := func() <-chan int {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			for i := range 10 {
				c <- i
			}
			close(c)
		}()
		wg.Done()
		return c
	}

	sumup := func(c <-chan int) <-chan int {
		n := 0
		for i := range c {
			n += i
		}
		sum := make(chan int, 1)
		sum <- n
		return sum
	}
	v := <-sumup(gen())
	fmt.Println(v)
}
