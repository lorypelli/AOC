package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	stats, _ := os.Stat("input.txt")
	size := stats.Size()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fileByte := make([]byte, size)
	file.Read(fileByte[:])
	arr := strings.Split(string(fileByte[:]), "\n")
	s := 0
	for i := 0; i < len(arr); i++ {
		firstCharacter := 0
		latestCharacter := 0
		for c := 0; c < len(arr[i]); c++ {
			character, _ := strconv.ParseFloat(string(arr[i][c]), 64)
			if (character > 0 && firstCharacter == 0) {
				firstCharacter = int(character)
			}
		}
		for c := len(arr[i]) - 1; c >= 0; c-- {
			character, _ := strconv.ParseFloat(string(arr[i][c]), 64)
			if (character > 0 && latestCharacter == 0) {
				latestCharacter = int(character)
			}
		}
		totalCharacter := fmt.Sprint(firstCharacter) + fmt.Sprint(latestCharacter)
		n, _ := strconv.Atoi(totalCharacter)
		s += n
	}
	fmt.Printf("The final sum is %d", s)
}