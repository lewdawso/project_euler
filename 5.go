package main

import (
	"log"
)

var out chan int
var list = []int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 9, 8}

//start from 20 and increment by 20
func gen(ch chan int) {
	for i := 20; ; i += 20 {
		ch <- i
	}
}

func filter(in chan int, out chan int, num int) {
	for {
		try := <-in
		if try%num == 0 {
			out <- try
		}
	}
}

func main() {
	in := make(chan int)
	go gen(in)
	for _, v := range list {
		out = make(chan int)
		go filter(in, out, v)
		in = out
	}
	log.Printf("solution is: %v", <-out)
}
