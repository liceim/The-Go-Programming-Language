//Exercise 9.5: Write a program with two goroutines that send messages back and forth over two unbuffered channels in ping-pong fashion. How many communications per second can the program sustain?


package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	q := make(chan int)
	var i int64
	start := time.Now()
	go func() {
		q <- 1
		for {
			i++
			q <- <-q
		}
	}()
	go func() {
		for {
			q <- <-q
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println(float64(i)/float64(time.Since(start))*1e9, "round trips per second")
}

/*
func main() {
	//ping := make(chan int)
	//pong := make(chan int)
	q := make(chan int)
	var count int64

	t := time.NewTimer(1 * time.Minute)
	start := time.Now()
	shutdown := make(chan struct{})
	done := make(chan struct{})

	go func() {
	loop:
		for {
			select {
			case <-shutdown:
				break loop
			case v := <-q:
				q <- v
				count++
			}
		}
		done <- struct{}{}
	}()

	go func() {
	loop:
		for {
			select {
			case <-shutdown:
				break loop
			case v := <-q:
				q <- v
				count++
			}
		}
		done <- struct{}{}
	}()

	q <- 1

	<-t.C
	close(shutdown)

	<-q

	<-done
	<-done
	t.Stop()
	fmt.Println(count)
	fmt.Println(float64(count)/float64(time.Since(start))*1e9, "rounds per second")
}
*/
