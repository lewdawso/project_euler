package main

import (
	"log"
)

const max = 1000

//check if pythagorean triplet
func triplet(a, b, c int) bool {
	if a*a+b*b == c*c {
		return true
	}
	return false
}

func sum(a, b, c int) bool {
	if a+b+c == 1000 {
		return true
	}
	return false
}

func main() {

	//only check combinations for which a < b < c
	for c := 3; c < max; c++ {
		for b := 2; b < c; b++ {
			for a := 1; a < b; a++ {
				if triplet(a, b, c) && sum(a, b, c) {
					log.Printf("a= %v b= %v c= %v", a, b, c)
					log.Printf("product abc = %v", a*b*c)
					return
				}
			}
		}
	}
}
