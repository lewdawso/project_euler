package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const file = "data"

type result struct {
	product int
	ints    []int
}

func main() {

	var r result

	digits, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	//stringify bytes and split into slice
	s := string(digits)
	ar := strings.Split(s, "")

	var nums [1000]int

	for i := 0; i < len(ar)-1; i++ { //trim trailing byte
		nums[i], _ = strconv.Atoi(ar[i])
	}

	for i := 0; i < len(nums)-13; i++ {
		var zero = false
		tmp := nums[i : i+13]
		for _, v := range tmp {
			if v == 0 {
				zero = true
				break
			}
		}
		if zero {
			continue
		}

		//have thirteen ints in tmp
		var p int = 1
		for _, v := range tmp {
			p *= v
		}

		if p > r.product {
			r.product = p
			r.ints = tmp
		}
	}
	log.Printf("largest product is: %v", r.product)
	log.Printf("thirteen digits are: %v", r.ints)
}
