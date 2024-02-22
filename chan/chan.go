package main

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	// go(fmt.Println, "goroutine")
	fmt.Println("main")

	for i := 0; i < 5; i++ {
		// Fix 3: Use go 1.22+
		// Fix 2: Loop local variable
		i := i
		go func() {
			fmt.Println(i)
		}()
		/* Fix 1: Pass a parameter
		go func(n int) {
			fmt.Println(n)
		}(i)
		*/
		/*
			go func() {
				fmt.Println(i) // Go < 1.22 BUG: All goroutines see the same i (closure)
			}()
		*/
	}

	time.Sleep(100 * time.Millisecond) // Never do that IRL

	// var ch chan string // ch is nil
	ch := make(chan string)
	go func() {
		ch <- "hi" // send
	}()
	msg := <-ch // receive
	fmt.Println("got:", msg)

	go func() { // producer/sender
		for i := range 4 {
			msg := fmt.Sprintf("msg #%d", i)
			ch <- msg
		}
		close(ch) // signal no more messages coming in
	}()

	/* What the "for range" does
	for {
		msg, ok := <- ch
		if !ok {
			break
		}
		fmt.Println("got:", msg)
	}
	*/

	for msg := range ch {
		fmt.Println("got:", msg)
	}

	msg = <-ch // ch is closed
	fmt.Printf("got (closed): %q\n", msg)
	// Problem: Did someone send "" or is the channel closed? (zero vs missing)
	msg, ok := <-ch // ok=false means the channel is closed
	fmt.Printf("got (closed): %q (ok=%v)\n", msg, ok)
	// ch <- "hi" // send to a closed channel (panic)

	// patterns
	/*
		fmt.Println("spinning future")
		fut := Future(func() int {
			time.Sleep(10 * time.Millisecond) // simulate work
			return 42
		})
		fmt.Println("doing other work")
		fmt.Println("getting result")
		val := <-fut
		fmt.Println("got:", val)

			fanOut()
			limitedFanOut()
			pooling()
	*/

	ch1, ch2 := make(chan int, 1), make(chan int, 1)
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- 1
	}()
	go func() {
		time.Sleep(20 * time.Millisecond)
		ch2 <- 2
	}()
	/*
		done := make(chan struct{})
		go func() {
			time.Sleep(time.Millisecond)
			close(done)
		}()
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel() // Common memory leak: forget to cancel a context

	select {
	case v := <-ch1:
		fmt.Println("ch1:", v)
	case v := <-ch2:
		fmt.Println("ch2:", v)
	case <-ctx.Done():
		fmt.Println("context expired:", ctx.Err())
		/*
			case <-time.After(5 * time.Millisecond):
				fmt.Println("timeout")
			case <-done:
				fmt.Println("cancelled")
		*/
	}

	drop()
	limit()
}

/* Patterns
- future/delayed/deferred/wait for result
	Start a function, read result later, block if not ready
- fan out
	Spin worker, collect results
	Great for I/O bound
	Need results back
- limitedFanOut
	Limited number of workers
	Good for CPU bound
	Need results back
- pooling
	Limited number of workers
	Don't care about results
- drop
	Don't want to flood resource (DB ...)
	OK to drop some messages
- limit
	Limit number of goroutines entering a section (semaphore)
*/

/*
  - limit
    Limit number of goroutines entering a section (semaphore)
*/
func limit() {
	ch := make(chan bool, 3) // semaphore
	count := int64(0)

	for i := range 10 {
		go func() {
			ch <- true
			atomic.AddInt64(&count, 1)
			defer func() {
				atomic.AddInt64(&count, -1)
				<-ch
			}()

			fmt.Printf("%d in (count=%d)\n", i, count)
		}()
	}

	time.Sleep(time.Second)
}

/*
  - drop
    Don't want to flood resource (DB ...)
    OK to drop some messages
*/
func drop() { // rate limit, OK to drop some
	const size = 10 // TODO: What is the optimal size?
	ch := make(chan int, size)

	// worker
	go func() {
		for v := range ch {
			fmt.Printf("worker got: %d\n", v)
		}
	}()

	for i := range 200 {
		fmt.Printf("len: %d, cap: %d\n", len(ch), cap(ch))
		select {
		case ch <- i:
			fmt.Printf("sent %d\n", i)
		default:
			fmt.Printf("dropped %d\n", i)
			// TODO: Spin a "janitor" process that will monitor the log and backfill the database
		}
	}
	close(ch)
	time.Sleep(time.Second)
}

/*
  - pooling
    Limited number of workers
    Don't care about results
*/
func pooling() { // CPU bound, don't care about result
	ch := make(chan int)
	n := 4
	var wg sync.WaitGroup

	wg.Add(n)

	fmt.Printf("spinning %d goroutines\n", n)
	for i := range n { // workers
		go func(id int) {
			defer wg.Done()
			for v := range ch {
				fmt.Printf("%d: got %d\n", id, v)
				time.Sleep(time.Microsecond)
			}
			fmt.Printf("%d shutting down\n", id)
		}(i)
	}

	fmt.Println("sending work")
	for i := range 17 {
		ch <- i
	}
	fmt.Println("shutting down")
	close(ch)
	wg.Wait()
}

/*
  - limitedFanOut
    Limited number of workers
    Good for CPU bound
    Need results back
*/
func limitedFanOut() { // CPU bound work loads
	words := strings.Fields("a little copying is better than a little dependency")
	in := make(chan string)
	out := make(chan string)

	// Spin goroutines as number of CPUs
	n := runtime.GOMAXPROCS(0)
	for range n {
		go func() {
			for w := range in {
				out <- strings.ToUpper(w)
			}
		}()
	}

	go func() {
		for _, w := range words {
			in <- w
		}
		close(in)
	}()

	for range words {
		w := <-out
		fmt.Println("limited:", w)
	}
}

/*
  - fan out
    Spin worker, collect results
    Great for I/O bound
    Need results back
*/
func fanOut() {
	words := strings.Fields("a little copying is better than a little dependency")
	ch := make(chan string)

	// fan out
	for _, w := range words {
		w := w // Go < 1.22
		go func() {
			ch <- strings.ToUpper(w)
		}()
	}

	// collect results
	for range words {
		w := <-ch
		fmt.Println(w)
	}
}

/*
	 future/delayed/deferred/wait for result
		Start a function, read result later, block if not ready
*/
func Future[T any](fn func() T) <-chan T {
	// goroutine leak: a goroutine is waiting for something that will never happen
	ch := make(chan T, 1) // buffered channel, to avoid goroutine leak

	go func() {
		ch <- fn()
	}()

	return ch
}

/* Channel sematics
- send/receive on a channel will block until opposite action
	- guarantee of delivery
	- buffered channel with cap "n" has "n" non blocking sends
- receive from a closed channel will return zero value without blocking
- send/close to/a closed channel will panic
	- channel ownership (usually the producer/sender)
- send/receive to/from a nil channel will block forever
*/

/*
Context switch: Moving a thread off the CPU, ~1us = 10k lost instructions
Two types of work loads:
- CPU bound: image processing, compression, hash (sha1) calculation, ...
	- Easy to know how many units of work (number of CPUs)
- IO bound: Reading a file, network access ...
	- Does not utilize the CPU
	- Hard to know how many units of work

Go scheduler: Turn IO bound to CPU bound

Two problems as a developer:
- Coordination: Wait for job to be done, wait for a signal ... (waitgroup, channels)
- Synchronization: Access to shared resource (mutex ...)
*/
