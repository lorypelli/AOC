package main

import (
	"fmt"
	"strconv"
	"aoc.globals.lorypelli"
)

func partialSum(arr []string) int  {
	s := 0
	for i := 0; i < len(arr); i++ {
		firstCharacter := 0
		latestCharacter := 0
		char := ""
		for c := 0; c < len(arr[i]); c++ {
			char = string(arr[i][c])
			character, _ := strconv.ParseFloat(char, 64)
			if character > 0 && firstCharacter == 0 {
				firstCharacter = int(character)
			}
		}
		for c := len(arr[i]) - 1; c >= 0; c-- {
			char = string(arr[i][c])
			character, _ := strconv.ParseFloat(char, 64)
			if character > 0 && latestCharacter == 0 {
				latestCharacter = int(character)
			}
		}
		totalCharacter := fmt.Sprint(firstCharacter) + fmt.Sprint(latestCharacter)
		n, _ := strconv.Atoi(totalCharacter)
		s += n
	}
	return s
}

func finalSum(arr []string) int  {
	s := 0
	for i := 0; i < len(arr); i++ {
		firstCharacter := 0
		latestCharacter := 0
		char := ""
		for c := 0; c < len(arr[i]); c++ {
			if string(arr[i][c]) == "o" && len(arr[i]) > c + 1 && string(arr[i][c + 1]) == "n" && len(arr[i]) > c + 2 && string(arr[i][c + 2]) == "e" {
				char = "1"
			} else if string(arr[i][c]) == "t" && len(arr[i]) > c + 1 && string(arr[i][c + 1]) == "w" && len(arr[i]) > c + 2 && string(arr[i][c + 2]) == "o" {
				char = "2"
			} else if string(arr[i][c]) == "t" && len(arr[i]) > c + 1 && string(arr[i][c + 1]) == "h" && len(arr[i]) > c + 2 && string(arr[i][c + 2]) == "r" && len(arr[i]) > c + 3 && string(arr[i][c + 3]) == "e" && len(arr[i]) > c + 4 && string(arr[i][c + 4]) == "e" {
				char = "3"
			} else if string(arr[i][c]) == "f" && len(arr[i]) > c + 1 && string(arr[i][c + 1]) == "o" && len(arr[i]) > c + 2 && string(arr[i][c + 2]) == "u" && len(arr[i]) > c + 3 && string(arr[i][c + 3]) == "r" {
				char = "4"
			} else if string(arr[i][c]) == "f" && len(arr[i]) > c + 1 && string(arr[i][c + 1]) == "i" && len(arr[i]) > c + 2 && string(arr[i][c + 2]) == "v" && len(arr[i]) > c + 3 && string(arr[i][c + 3]) == "e" {
				char = "5"
			} else if string(arr[i][c]) == "s" && len(arr[i]) > c + 1 && string(arr[i][c + 1]) == "i" && len(arr[i]) > c + 2 && string(arr[i][c + 2]) == "x" {
				char = "6"
			} else if string(arr[i][c]) == "s" && len(arr[i]) > c + 1 && string(arr[i][c + 1]) == "e" && len(arr[i]) > c + 2 && string(arr[i][c + 2]) == "v" && len(arr[i]) > c + 3 && string(arr[i][c + 3]) == "e" && len(arr[i]) > c + 4 && string(arr[i][c + 4]) == "n" {
				char = "7"
			} else if string(arr[i][c]) == "e" && len(arr[i]) > c + 1 && string(arr[i][c + 1]) == "i" && len(arr[i]) > c + 2 && string(arr[i][c + 2]) == "g" && len(arr[i]) > c + 3 && string(arr[i][c + 3]) == "h" && len(arr[i]) > c + 4 && string(arr[i][c + 4]) == "t" {
				char = "8"
			} else if string(arr[i][c]) == "n" && len(arr[i]) > c + 1 && string(arr[i][c + 1]) == "i" && len(arr[i]) > c + 2 && string(arr[i][c + 2]) == "n" && len(arr[i]) > c + 3 && string(arr[i][c + 3]) == "e" {
				char = "9"
			} else {
				char = string(arr[i][c])
			}
			character, _ := strconv.ParseFloat(char, 64)
			if character > 0 && firstCharacter == 0 {
				firstCharacter = int(character)
			}
		}
		for c := len(arr[i]) - 1; c >= 0; c-- {
			if string(arr[i][c]) == "o" && len(arr[i]) > c + 1 && string(arr[i][c + 1]) == "n" && len(arr[i]) > c + 2 && string(arr[i][c + 2]) == "e" {
				char = "1"
			} else if string(arr[i][c]) == "t" && len(arr[i]) > c + 1 && string(arr[i][c + 1]) == "w" && len(arr[i]) > c + 2 && string(arr[i][c + 2]) == "o" {
				char = "2"
			} else if string(arr[i][c]) == "t" && len(arr[i]) > c + 1 && string(arr[i][c + 1]) == "h" && len(arr[i]) > c + 2 && string(arr[i][c + 2]) == "r" && len(arr[i]) > c + 3 && string(arr[i][c + 3]) == "e" && len(arr[i]) > c + 4 && string(arr[i][c + 4]) == "e" {
				char = "3"
			} else if string(arr[i][c]) == "f" && len(arr[i]) > c + 1 && string(arr[i][c + 1]) == "o" && len(arr[i]) > c + 2 && string(arr[i][c + 2]) == "u" && len(arr[i]) > c + 3 && string(arr[i][c + 3]) == "r" {
				char = "4"
			} else if string(arr[i][c]) == "f" && len(arr[i]) > c + 1 && string(arr[i][c + 1]) == "i" && len(arr[i]) > c + 2 && string(arr[i][c + 2]) == "v" && len(arr[i]) > c + 3 && string(arr[i][c + 3]) == "e" {
				char = "5"
			} else if string(arr[i][c]) == "s" && len(arr[i]) > c + 1 && string(arr[i][c + 1]) == "i" && len(arr[i]) > c + 2 && string(arr[i][c + 2]) == "x" {
				char = "6"
			} else if string(arr[i][c]) == "s" && len(arr[i]) > c + 1 && string(arr[i][c + 1]) == "e" && len(arr[i]) > c + 2 && string(arr[i][c + 2]) == "v" && len(arr[i]) > c + 3 && string(arr[i][c + 3]) == "e" && len(arr[i]) > c + 4 && string(arr[i][c + 4]) == "n" {
				char = "7"
			} else if string(arr[i][c]) == "e" && len(arr[i]) > c + 1 && string(arr[i][c + 1]) == "i" && len(arr[i]) > c + 2 && string(arr[i][c + 2]) == "g" && len(arr[i]) > c + 3 && string(arr[i][c + 3]) == "h" && len(arr[i]) > c + 4 && string(arr[i][c + 4]) == "t" {
				char = "8"
			} else if string(arr[i][c]) == "n" && len(arr[i]) > c + 1 && string(arr[i][c + 1]) == "i" && len(arr[i]) > c + 2 && string(arr[i][c + 2]) == "n" && len(arr[i]) > c + 3 && string(arr[i][c + 3]) == "e" {
				char = "9"
			} else {
				char = string(arr[i][c])
			}
			character, _ := strconv.ParseFloat(char, 64)
			if character > 0 && latestCharacter == 0 {
				latestCharacter = int(character)
			}
		}
		totalCharacter := fmt.Sprint(firstCharacter) + fmt.Sprint(latestCharacter)
		n, _ := strconv.Atoi(totalCharacter)
		s += n
	}
	return s
}

func main() {
	arr := globals.SplitFile("input.txt")
	fmt.Printf("The partial sum is %d\n", partialSum(arr))
	fmt.Printf("The final sum is %d", finalSum(arr))
}