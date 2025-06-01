package channels

import (
	"fmt"
)

func Ex1() {
	/**
	This code blocks
	channels blocks
	c <- 42 blocks and does not further carry on
	*/
	c := make(chan int)

	// put into channel 42
	c <- 42

	// get out of channel
	fmt.Println(<-c)

}

func Ex2() {
	c := make(chan int)

	// fired off
	// and main thread continues
	go func() {
		// blocks in this goroutine
		c <- 42
	}()

	// pulled off and the previous threads released
	fmt.Println(<-c)
}

func Ex3() {
	// this one works
	// because we specified the buffer
	c := make(chan int, 1)
	c <- 42
	fmt.Println(<-c)
}

func Ex4() {
	// unsuccessful buffer
	c := make(chan int, 1)
	c <- 42
	c <- 43 // buffer blocked 1 roam still occupied

	fmt.Println(<-c)
}

func Ex5() {
	// successful buffer
	c := make(chan int, 2)
	c <- 42
	c <- 43

	// first in first out
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func Ex6() {
	fmt.Println("-----")
	// channels directions
	// put direction
	c := make(chan<- int)
	// recieve direction
	d := make(<-chan int)
	fmt.Printf("%T\n", c)
	go func() { c <- 42 }()
	fmt.Println("-----")
	fmt.Println(<-d)
}

func foo(c chan<- int) {
	c <- 42
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}

func Ex7() {
	c := make(chan int)
	go foo(c)
	bar(c)
	fmt.Println("about exit")
}

func Ex8() {

	c := make(chan int)

	go func() {
		for i := range 5 {
			c <- i
		}
		close(c)
	}()

	// range over channels
	for v := range c {
		fmt.Println(v)
	}
}

func send(e, o, q chan<- int) {
	for i := range 100 {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
	close(q)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case y := <-e:
			fmt.Println("From even channel", y)
		case x := <-o:
			fmt.Println("From odd channel", x)
		case i, ok := <-q:
			if !ok {
				fmt.Println("From Quit Channel OK", i, ok)
			} else {
				fmt.Println("From commma ok", ok)
			}
			return
		}
	}
}

func Ex9() {
	// tools to work with channels
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go send(even, odd, quit)
	receive(even, odd, quit)
}

func Ex10() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()
	v, ok := <-c
	fmt.Println(v, ok)
	fmt.Println("---")
	v, ok = <-c
	fmt.Println(v, ok)
}
