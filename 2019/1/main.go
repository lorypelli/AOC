package main

import (
	"fmt"
	"strconv"
	"math"
	"aoc.globals.lorypelli"
)

func moduleSum(arr []string) int {
	s := 0
	for i := 0; i < len(arr); i++ {
		n, _ := strconv.Atoi(arr[i])
		n = int(math.Round(float64(n / 3))) - 2
		s += n
	}
	return s
}

func loopModuleSum(arr []string) int {
	s := 0
	for i := 0; i < len(arr); i++ {
		n, _ := strconv.Atoi(arr[i])
		n = int(math.Round(float64(n / 3))) - 2
		s += n
		for n > 0 {
			n = int(math.Round(float64(n / 3))) - 2
			if n < 0 {
				n = 0
			}
			s += n
		}
	}
	return s
}

func main() {
	arr := globals.SplitFile("input.txt")
	fmt.Printf("The sum of all modules is %d\n", moduleSum(arr))
	fmt.Printf("The sum of all modules looped until zero is %d", loopModuleSum(arr))
}