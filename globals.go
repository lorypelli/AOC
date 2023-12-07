package globals

import (
	"os"
	"strings"
)

func SplitFile(fileName string) []string {
	file, _ := os.Open(fileName)
	stats, _ := os.Stat(fileName)
	size := stats.Size()
	fileByte := make([]byte, size)
	fileSlice := fileByte[:]
	file.Read(fileSlice)
	arr := strings.Split(string(fileSlice), "\n")
	return arr
}
