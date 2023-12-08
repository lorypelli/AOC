package main

import (
	"aoc.globals.lorypelli"
	"fmt"
	"strconv"
	"strings"
)

func steps(arr []string) int {
	h := 0
	d := 0
	for i := 0; i < len(arr); i++ {
		move := strings.Split(arr[i], " ")[0]
		step, _ := strconv.Atoi(strings.Split(arr[i], " ")[1])
		switch move {
		case "forward":
			{
				h += step
				break
			}
		case "down":
			{
				d += step
				break
			}
		case "up":
			{
				d -= step
				break
			}
		}
	}
	return h * d
}

func newSteps(arr []string) int {
	h := 0
	d := 0
	a := 0
	for i := 0; i < len(arr); i++ {
		move := strings.Split(arr[i], " ")[0]
		step, _ := strconv.Atoi(strings.Split(arr[i], " ")[1])
		switch move {
		case "forward":
			{
				h += step
				d += a * step
				break
			}
		case "down":
			{
				a += step
				break
			}
		case "up":
			{
				a -= step
				break
			}
		}
	}
	return h * d
}

func main() {
	arr := globals.SplitFile("input.txt")
	fmt.Printf("Your final horizontal position and your final depth position multiplied together give %d\n", steps(arr))
	fmt.Printf("Your final horizontal position and your final depth position multiplied together based on new counting process give %d", newSteps(arr))
}
