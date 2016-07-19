package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var integers []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		integers = append(integers, num)
	}
	//fmt.Println(inversions(integers[:]))
	fmt.Println(countInversions(integers[:]))
}

func countInversions(arr []int) (inversions int) {
	n := len(arr)
	if n == 1 {
		return
	}

	return countInversions(arr[:n/2]) +
		countInversions(arr[n/2:]) +
		countSplitInversions(arr)
}

func countSplitInversions(arr []int) (inversions int) {
	n := len(arr)
	out := make([]int, 0, len(arr))

	left := arr[:n/2]
	right := arr[n/2:]

	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			out = append(out, left[0])
			left = left[1:]
		} else {
			out = append(out, right[0])
			right = right[1:]
			inversions += len(left)
		}
	}
	out = append(out, left...)
	// right slice need not be appended as it is
	// already at the end of arr
	copy(arr, out)
	return
}

// inversions is a brute-force implementation
func inversions(arr []int) (inversions int) {
	for i, a := range arr {
		for _, b := range arr[i:] {
			if a > b {
				inversions++
			}
		}
	}
	return
}
