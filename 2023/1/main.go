package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "input.txt"
	file, _ := os.Open(fileName)
	stats, _ := os.Stat(fileName)
	size := stats.Size()
	fileByte := make([]byte, size)
	fileSlice := fileByte[:]
	file.Read(fileSlice)
	arr := strings.Split(string(fileSlice), "\n")
	s := 0
	for i := 0; i < len(arr); i++ {
		firstCharacter := 0
		latestCharacter := 0
		for c := 0; c < len(arr[i]); c++ {
			char := string(arr[i][c])
			character, _ := strconv.ParseFloat(char, 64)
			if character > 0 && firstCharacter == 0 {
				firstCharacter = int(character)
			}
		}
		for c := len(arr[i]) - 1; c >= 0; c-- {
			char := string(arr[i][c])
			character, _ := strconv.ParseFloat(char, 64)
			if character > 0 && latestCharacter == 0 {
				latestCharacter = int(character)
			}
		}
		totalCharacter := fmt.Sprint(firstCharacter) + fmt.Sprint(latestCharacter)
		n, _ := strconv.Atoi(totalCharacter)
		s += n
	}
	fmt.Printf("The final sum is %d", s)
}