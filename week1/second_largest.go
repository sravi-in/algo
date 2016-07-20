package main

import (
	"fmt"
)

func main() {
	fmt.Println(secondLargest([]int{5, 7, 3, 8, 9}))
	fmt.Println(secondLargest([]int{2, 1}))
}

// Assumes arr has atleast 2 elems
func secondLargest(arr []int) int {
	_, losers := largest(arr)
	max := losers[0]
	for _, i := range losers[1:] {
		if i > max {
			max = i
		}
	}
	return max
}

func largest(arr []int) (int, []int) {
	n := len(arr)
	if n == 1 {
		return arr[0], nil
	}
	left, llosers := largest(arr[:n/2])
	right, rlosers := largest(arr[n/2:])

	if left > right {
		return left, append(llosers, right)
	}
	return right, append(rlosers, left)
}
