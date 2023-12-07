package main

import (
	"aoc.globals.lorypelli"
	"fmt"
	"strconv"
)

func sum(arr []string) int {
	s := 0
	for i := 0; i < len(arr); i++ {
		n, _ := strconv.Atoi(arr[i])
		s += n
	}
	return s
}

func main() {
	arr := globals.SplitFile("input.txt")
	fmt.Printf("The sum of all frequencies is %d", sum(arr))
}
