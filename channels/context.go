package channels

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// https://go.dev/blog/context

func Ex11() {
	ctx := context.Background()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("-------")

	ctx, _ = context.WithCancel(ctx)

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("-------")

	ctx2, cancel := context.WithCancel(ctx)
	cancel()
	fmt.Println("context:\t", ctx2)
	fmt.Println("context err:\t", ctx2.Err())
	fmt.Println("context cancel:\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("-------")
}

func Ex12() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num goroutines 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num goroutines 2", runtime.NumGoroutine())
	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("cancelled")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num goroutines 3", runtime.NumGoroutine())
}
