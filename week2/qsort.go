package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Initial partition fn, doesn't do anything so
// the pivot is picked as first element
var choosePartition = func(arr []int) {
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	integers := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		integers = append(integers, i)
	}

	fmt.Println("# of comparisons with first pivot:\t", qsort(append([]int(nil), integers...)))
	choosePartition = chooseLast
	fmt.Println("# of comparisons with last pivot:\t", qsort(append([]int(nil), integers...)))
	choosePartition = chooseRdm
	fmt.Println("# of comparisons with random pivot:\t", qsort(append([]int(nil), integers...)))
	choosePartition = chooseMedn
	fmt.Println("# of comparisons with median pivot:\t", qsort(integers))
}

func qsort(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}

	choosePartition(arr)
	i := partition(arr)
	return qsort(arr[:i]) + qsort(arr[i+1:]) + len(arr) - 1
}

func chooseLast(arr []int) {
	n := len(arr) - 1
	arr[0], arr[n] = arr[n], arr[0]
}

func chooseMedn(arr []int) {
	n := len(arr) - 1
	m := n / 2

	switch {
	case arr[0] < arr[m] && arr[m] < arr[n], arr[n] < arr[m] && arr[m] < arr[0]:
		arr[0], arr[m] = arr[m], arr[0]
	case arr[0] < arr[n] && arr[n] < arr[m], arr[m] < arr[n] && arr[n] < arr[0]:
		arr[0], arr[n] = arr[n], arr[0]
	}
}

func chooseRdm(arr []int) {
	r := rand.Intn(len(arr))
	arr[0], arr[r] = arr[r], arr[0]
}

func partition(arr []int) int {
	pivot := arr[0]
	i := 1
	for j := 1; j < len(arr); j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[0], arr[i-1] = arr[i-1], arr[0]
	return i - 1
}

func readFile(file string) (integers []int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		integers = append(integers, i)
	}
	return
}
