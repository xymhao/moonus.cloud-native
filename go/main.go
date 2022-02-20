package main

import "fmt"

func main() {
	grownSlices := make([]int, 1)
	capCount := cap(grownSlices)
	for i := 0; i < 1000000; i++ {
		grownSlices = append(grownSlices, i)
		if capCount != cap(grownSlices) {
			grownFactory := float64(cap(grownSlices)) / float64(capCount)
			capCount = cap(grownSlices)
			fmt.Println(capCount, grownFactory)
		}
	}
}
