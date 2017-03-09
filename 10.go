package main

import (
	"log"
)

const max = 2000000

var nums [max + 1]bool
var sum int

func sieve() {
	for i, _ := range nums {
		if nums[i] == false {
			continue
		}
		for j := i + i; j < max+1; j += i {
			nums[j] = false
		}
	}
}

func main() {

	for i, _ := range nums {
		nums[i] = true
	}

	nums[0] = false
	nums[1] = false

	sieve()

	for i, v := range nums {
		if v {
			sum += i
		}
	}
	log.Println(sum)
}
