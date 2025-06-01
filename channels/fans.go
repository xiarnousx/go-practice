package channels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func fanout(e, o chan<- int) {
	for i := range 10 {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

func r(e, o <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range e {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			fanin <- v * 2
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}

func FanIn() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go fanout(even, odd)

	go r(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}
}

func FanOut() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")

}

func populate(c chan int) {
	for i := range 10 {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup

	for v := range c1 {
		wg.Add(1)
		go func(v int) {
			c2 <- timeConsumingWork(v)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(v int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return v + rand.Intn(1000)
}
