package main

import (
    "fmt"
    "os"
	"strings"
	"strconv"
)

func main() {
    data, err := os.ReadFile("inputs/day01.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	lines := strings.Split(string(data), "\n")

	// part1
	// pos := 50
	// countAtZero := 0

	// for _, line := range lines {
	// 	if line == "" {
	// 		continue
	// 	}

	// 	direction := line[0]
	// 	distance, err := strconv.Atoi(line[1:])
	// 	if err != nil {
	// 		fmt.Println("Error parsing distance:", err)
	// 		return
	// 	}

	// 	if direction == 'L' {
    //         pos -= distance
    //     } else if direction == 'R' {
    //         pos += distance
    //     }

	// 	// ensure pos stays within the range 0-99
	// 	pos = (pos + 100) % 100

	// 	if pos == 0 {
	// 		countAtZero++
	// 	}
	// }

	// fmt.Println("part1:", countAtZero)

	// part2
	pos := 50
	countAtZero := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Error parsing distance:", err)
			return
		}

		if direction == 'L' {
            for i := 0; i < distance; i++ {
                pos--
                if pos < 0 {
                    pos = 99
                }
                if pos == 0 {
                    countAtZero++
                }
            }
        } else if direction == 'R' {
            for i := 0; i < distance; i++ {
                pos++
                if pos > 99 {
                    pos = 0
					countAtZero++
                }
            }
        }
	}
	fmt.Println("part2:", countAtZero)
}
