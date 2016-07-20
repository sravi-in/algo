package main

import (
	"fmt"
)

func main() {
	fmt.Println(secondLargest([]int{5, 7, 3, 8, 9}))
	fmt.Println(secondLargest([]int{2, 1}))
	fmt.Println(secondLargest([]int{7, 3, 7, 7, 7, 7, 7, 7}))
	fmt.Println(secondLargest([]int{7, 3, 1, 5, 6, 8, 4, 2}))
	fmt.Println(secondLargest([]int{7, 3, 1, 7, 6, 7, 4, 2}))
}

// Assumes arr has atleast 2 elems
func secondLargest(arr []int) int {
	_, losers := largest(arr)
	fmt.Println(losers)
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
	} else if right > left {
		return right, append(rlosers, left)
	}

	// Both left and right are equal, return any but losers
	// could be in both
	return left, append(llosers, rlosers...)
}
