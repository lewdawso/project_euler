package main

//10,001st prime number

import (
	"log"
)

const limit = 10001

var prime int

func gen(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in chan int, out chan int, prime int) {
	for {
		if a := <-in; a%prime != 0 {
			out <- a
		}
	}
}

func main() {
	in := make(chan int)
	go gen(in)
	for i := 1; i < limit+1; i++ {
		prime = <-in
		out := make(chan int)
		go filter(in, out, prime)
		in = out
	}
	log.Printf("result is: %v", prime)
}
