package main

import (
	"aoc.globals.lorypelli"
	"fmt"
	"strconv"
	"strings"
)

func validPasswords(arr []string) int {
	s := 0
	for i := 0; i < len(arr); i++ {
		j := 0
		password := strings.Split(arr[i], ":")[1][1:]
		criterion := strings.Split(arr[i], ":")[0]
		min, _ := strconv.Atoi(string(strings.Split(criterion, "-")[0]))
		trimEnd := strings.Split(criterion, " ")[0]
		max, _ := strconv.Atoi(string(strings.Split(trimEnd, "-")[1]))
		letterLength := len(strings.Split(arr[i], ":")[0]) - 1
		letter := string(strings.Split(arr[i], ":")[0][letterLength])
		for c := 0; c < len(password); c++ {
			if string(password[c]) == letter {
				j += 1
			}
		}
		if j >= min && j <= max {
			s += 1
		}
	}
	return s
}

func secondValidPasswords(arr []string) int {
	s := 0
	for i := 0; i < len(arr); i++ {
		j := 0
		password := strings.Split(arr[i], ":")[1][1:]
		criterion := strings.Split(arr[i], ":")[0]
		firstLetter, _ := strconv.Atoi(string(strings.Split(criterion, "-")[0]))
		trimEnd := strings.Split(criterion, " ")[0]
		secondLetter, _ := strconv.Atoi(string(strings.Split(trimEnd, "-")[1]))
		letterLength := len(strings.Split(arr[i], ":")[0]) - 1
		letter := string(strings.Split(arr[i], ":")[0][letterLength])
		if (string(password[firstLetter-1]) == letter && len(password) > secondLetter-1 && string(password[secondLetter-1]) != letter) || (string(password[firstLetter-1]) != letter && len(password) > secondLetter-1 && string(password[secondLetter-1]) == letter) {
			j += 1
		}
		if j > 0 {
			s += 1
		}
	}
	return s
}

func main() {
	arr := globals.SplitFile("input.txt")
	fmt.Printf("The number of valid passwords in the array is %d\n", validPasswords(arr))
	fmt.Printf("The number of valid passwords in the array according to the new rules is %d", secondValidPasswords(arr))
}
