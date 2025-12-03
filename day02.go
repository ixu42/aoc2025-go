package main

import (
    "fmt"
    "os"
	"strings"
	"strconv"
)

func getIdsFromPair(pair string) (int, int, error) {
	idPair := strings.Split(pair, "-")

	if len(idPair) != 2 {
		return 0, 0, fmt.Errorf("invalid id pair: %s", pair)
	}

	firstId, err1 := strconv.Atoi(idPair[0])
	lastId, err2 := strconv.Atoi(idPair[1])
	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("error converting ids: %v, %v", err1, err2)
	}

	return firstId, lastId, nil
}

func isRepeatedTwice(n int) bool {
	str := strconv.Itoa(n)

	if len(str) % 2 != 0 {
		return false
	}

	mid := len(str) / 2
	return str[:mid] == str[mid:]
}

func isRepeatedAtLeastTwice(n int) bool {
	str := strconv.Itoa(n)
	maxRepeat := len(str)

	for i := 2; i <= maxRepeat; i++ {
		if len(str) % i != 0 {
			continue
		}

		chunkSize := len(str) / i
		pattern := str[:chunkSize]
		matched := true

		for j := 1; j < i; j++ {
			if str[j*chunkSize:(j+1)*chunkSize] != pattern {
				matched = false
				break
			}
		}

		if matched {
			return true
		}
	}
	return false
}

func solvePart1(pairs []string) int {
	res := 0

	for _, pair := range pairs {
		firstId, lastId, err := getIdsFromPair(pair)
		if err != nil {
			fmt.Println(err)
			return 0
		}

		for i := firstId; i <= lastId; i++ {
			if isRepeatedTwice(i) {
				res += i
			}
		}
	}

	return res
}

func solvePart2(pairs []string) int {
	res := 0

	for _, pair := range pairs {
		firstId, lastId, err := getIdsFromPair(pair)
		if err != nil {
			fmt.Println(err)
			return 0
		}

		for i := firstId; i <= lastId; i++ {
			if isRepeatedAtLeastTwice(i) {
				res += i
			}
		}
	}

	return res
}

func main() {
    data, err := os.ReadFile("inputs/day02.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	pairs := strings.Split(string(data), ",")

	fmt.Println("part1:", solvePart1(pairs))
	fmt.Println("part2:", solvePart2(pairs))
}
