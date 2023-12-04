package main

import (
	"fmt"
	"strconv"
	"aoc.globals.lorypelli"
)

func greatestNumber(arr []string) int {
	s := 0
	var sums []int
	for i := 0; i < len(arr); i++ {
		n, _ := strconv.Atoi(arr[i])
		if n != 0 {
			s += n
		} else {
			sums = append(sums, s)
			s = 0
		}
	}
	for i := 0; i < len(sums); i++ {
		if sums[0] < sums[i] {
			sums[0] = sums[i]
		}
	}
	return sums[0]
}

func topNumbers(arr []string) int {
	s := 0
	var sums []int
	for i := 0; i < len(arr); i++ {
		n, _ := strconv.Atoi(arr[i])
		if n != 0 {
			s += n
		} else {
			sums = append(sums, s)
			s = 0
		}
	}
	for i := 0; i < len(sums); i++ {
		if sums[0] < sums[i] {
			sums[0] = sums[i]
		}
		if sums[1] < sums[i] && sums[0] > sums[i] {
			sums[1] = sums[i]
		}
		if sums[2] < sums[i] && sums[1] > sums[i] {
			sums[2] = sums[i]
		}
	}
	return sums[0] + sums[1] + sums[2]
}

func main() {
	arr := globals.SplitFile("input.txt")
	fmt.Printf("The greatest number in the array is %d\n", greatestNumber(arr))
	fmt.Printf("The sum of the top 3 numbers in the array is %d", topNumbers(arr))
}