package main

import (
	"aoc.globals.lorypelli"
	"fmt"
	"strconv"
)

func twoNumbers(arr []string) int {
	for i := 0; i < len(arr); i++ {
		n, _ := strconv.Atoi(arr[i])
		for c := 0; c < len(arr); c++ {
			n2, _ := strconv.Atoi(arr[c])
			if n+n2 == 2020 {
				return n * n2
			}
		}
	}
	return 0
}

func threeNumbers(arr []string) int {
	for i := 0; i < len(arr); i++ {
		n, _ := strconv.Atoi(arr[i])
		for c := 0; c < len(arr); c++ {
			n2, _ := strconv.Atoi(arr[c])
			for j := 0; j < len(arr); j++ {
				n3, _ := strconv.Atoi(arr[j])
				if n+n2+n3 == 2020 {
					return n * n2 * n3
				}
			}
		}
	}
	return 0
}

func main() {
	arr := globals.SplitFile("input.txt")
	fmt.Printf("The two numbers whose sum is 2020 multiplied together give %d\n", twoNumbers(arr))
	fmt.Printf("The three numbers whose sum is 2020 multiplied together give %d", threeNumbers(arr))
}
