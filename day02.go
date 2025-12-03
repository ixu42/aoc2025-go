package main

import (
    "fmt"
    "os"
	"strings"
	"strconv"
)

func isRepeatedTwice(n int) bool {
	str := strconv.Itoa(n)

	if len(str) % 2 != 0 {
		return false
	}

	mid := len(str) / 2
	return str[:mid] == str[mid:]
}

func main() {
    data, err := os.ReadFile("inputs/day02.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	pairs := strings.Split(string(data), ",")

	res := 0

	for _, pair := range pairs {
		idPair := strings.Split(pair, "-")
		if len(idPair) != 2 {
			fmt.Println("Invalid id pair:", pair)
			return
		}

		firstId, err1 := strconv.Atoi(idPair[0])
		lastId, err2 := strconv.Atoi(idPair[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting ids:", err1, err2)
			return
		}

		for i := firstId; i <= lastId; i++ {
			if isRepeatedTwice(i) {
				res += i
			}
		}
	}

	fmt.Println("part1:", res)
}
