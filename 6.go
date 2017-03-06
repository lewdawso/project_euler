package main

import (
	"log"
)

//sum=n(n+1)/2
func sum(n int) int {
	return n * (n + 1) / 2
}

func main() {
	s := sum(100)
	var a int
	for i := 1; i < 101; i++ {
		a += i * i
	}
	log.Println(s*s - a)
}
