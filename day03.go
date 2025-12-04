package main

import (
    "fmt"
    "os"
	"strings"
	"strconv"
)

func getMax2Digits (s string) int {
	digits := []byte(s)
	maxNum := -1

	for i := 0; i < len(digits) - 1; i++ {
		d1 := int(digits[i] - '0')
		for j := i + 1; j < len(digits); j++ {
			d2 := int(digits[j] - '0')
			val := d1 * 10 + d2
			if val > maxNum {
				maxNum = val
			}
		}
	}

	return maxNum
}

func getMax12Digits(s string) int {
	need := 12
	stack := make([]byte, 0, need)
	totalDigits := len(s)
	canDrop := totalDigits - need

	for i := 0; i < totalDigits; i++ {
		digit := s[i]

		// pop smaller digits from stack while we can still fill 12 digits
		for len(stack) > 0 && canDrop > 0 && stack[len(stack)-1] < digit {
			stack = stack[:len(stack)-1] // pop the last element
			canDrop--
		}

		// add digit if we need more, otherwise, drop it
		if len(stack) < need {
			stack = append(stack, digit)
		} else {
			canDrop--
		}
	}
	maxNum, err := strconv.Atoi(string(stack))
	if err != nil {
		fmt.Println("Error converting to int:", err)
		return 0
	}

	return maxNum
}

func main() {
    data, err := os.ReadFile("inputs/day03.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	lines := strings.Split(string(data), "\n")

	// part1
	res := 0
	for _, line := range lines {
		res += getMax2Digits(line)
	}
	fmt.Println("part1:", res)

	// part2
	res = 0
	for _, line := range lines {
		res += getMax12Digits(line)
	}
	fmt.Println("part2:", res)
}
