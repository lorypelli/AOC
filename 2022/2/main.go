package main

import (
	"fmt"

	"github.com/lorypelli/globals"
)

func score(arr []string) int {
	s := 0
	choice := 0
	score := 0
	for i := 0; i < len(arr); i++ {
		switch string(arr[i][0]) {
		case "A":
			{
				switch string(arr[i][2]) {
				case "X":
					{
						choice = 1
						score = choice + 3
						s += score
						break
					}
				case "Y":
					{
						choice = 2
						score = choice + 6
						s += score
						break
					}
				case "Z":
					{
						choice = 3
						score = choice + 0
						s += score
						break
					}
				}
				break
			}
		case "B":
			{
				switch string(arr[i][2]) {
				case "X":
					{
						choice = 1
						score = choice + 0
						s += score
						break
					}
				case "Y":
					{
						choice = 2
						score = choice + 3
						s += score
						break
					}
				case "Z":
					{
						choice = 3
						score = choice + 6
						s += score
						break
					}
				}
				break
			}
		case "C":
			{
				switch string(arr[i][2]) {
				case "X":
					{
						choice = 1
						score = choice + 6
						s += score
						break
					}
				case "Y":
					{
						choice = 2
						score = choice + 0
						s += score
						break
					}
				case "Z":
					{
						choice = 3
						score = choice + 3
						s += score
						break
					}
				}
				break
			}
		}
	}
	return s
}

func secondScore(arr []string) int {
	s := 0
	choice := 0
	score := 0
	for i := 0; i < len(arr); i++ {
		switch string(arr[i][0]) {
		case "A":
			{
				switch string(arr[i][2]) {
				case "X":
					{
						choice = 3
						score = choice + 0
						s += score
						break
					}
				case "Y":
					{
						choice = 1
						score = choice + 3
						s += score
						break
					}
				case "Z":
					{
						choice = 2
						score = choice + 6
						s += score
						break
					}
				}
				break
			}
		case "B":
			{
				switch string(arr[i][2]) {
				case "X":
					{
						choice = 1
						score = choice + 0
						s += score
						break
					}
				case "Y":
					{
						choice = 2
						score = choice + 3
						s += score
						break
					}
				case "Z":
					{
						choice = 3
						score = choice + 6
						s += score
						break
					}
				}
				break
			}
		case "C":
			{
				switch string(arr[i][2]) {
				case "X":
					{
						choice = 2
						score = choice + 0
						s += score
						break
					}
				case "Y":
					{
						choice = 3
						score = choice + 3
						s += score
						break
					}
				case "Z":
					{
						choice = 1
						score = choice + 6
						s += score
						break
					}
				}
				break
			}
		}
	}
	return s
}

func main() {
	arr := globals.SplitFile("input.txt")
	fmt.Printf("The total score is %d\n", score(arr))
	fmt.Printf("The total score with the new strategy is %d", secondScore(arr))
}
