package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
*
Go lang a major language built to take native advantage of multi core cpus
2006 Intel released first daul core processor
2012 Google introduced Go lang to open source community
concurrency does not gurantees code run in parallel
when have multi core cpu then concurrent code runs in parallel

# Do not communicate by sharing memory, share memory by communicating

go routines needs channels to signal completion
if a go routine function returns a value, it's return is discarded
*/
var wg sync.WaitGroup

func foo() {
	for i := range 10 {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := range 10 {
		fmt.Println("bar:", i)
	}
	wg.Done()
}

func Ex1() {
	wg.Add(2)
	go foo()
	go bar()

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())
	wg.Wait()
}

func Ex2() {
	doSomething := func(x int) int {
		return x * 5
	}

	ch := make(chan int)

	go func() {
		ch <- doSomething(5)
	}()

	fmt.Println(<-ch)

}

func Ex3() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	counter := 0
	var wg sync.WaitGroup

	const gs = 100

	wg.Add(gs)

	var mu sync.Mutex

	for range gs {
		go func() {
			mu.Lock()
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("-----")
	fmt.Println(counter)
	fmt.Println("-----")
	fmt.Println("Goroutines:", runtime.NumGoroutine())
}

func Ex4() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	var counter int64
	var wg sync.WaitGroup

	const gs = 100

	wg.Add(gs)

	for range gs {
		go func() {
			//time.Sleep(time.Second)
			atomic.AddInt64(&counter, 1)
			fmt.Println("->\t", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("-----")
	fmt.Println(counter)
	fmt.Println("-----")
	fmt.Println("Goroutines:", runtime.NumGoroutine())
}
