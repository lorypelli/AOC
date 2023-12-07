package main

import (
	"aoc.globals.lorypelli"
	"fmt"
	"strconv"
)

func largerThan(arr []string) int {
	s := 0
	for i := 0; i < len(arr); i++ {
		n, _ := strconv.Atoi(arr[i])
		previous := 0
		if i-1 >= 0 {
			previous, _ = strconv.Atoi(arr[i-1])
		}
		if previous != 0 && n > previous {
			s += 1
		}
	}
	return s
}

func main() {
	arr := globals.SplitFile("input.txt")
	fmt.Printf("The number of measurements that are larger than the previous one is %d", largerThan(arr))
}
