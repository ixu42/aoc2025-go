package main

import (
    "fmt"
    "os"
	"strings"
)

var memo = make(map[[2]int]int)

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func countSplit(grid [][]rune) int {
	splitCount := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 'S' ||  grid[r][c] == '|' {
				if r + 1 < len(grid) && grid[r+1][c] == '.' {
					grid[r+1][c] = '|'
				}
			} else if grid[r][c] == '^' && r - 1 >= 0 && grid[r-1][c] == '|' {
				splitCount++
				if c - 1 >= 0 && grid[r][c-1] == '.' {
					grid[r][c-1] = '|'
					if r + 1 < len(grid) && grid[r+1][c-1] == '.' {
						grid[r+1][c-1] = '|'
					}
				}
				if r + 1 < len(grid[r]) && grid[r][c+1] == '.' {
					grid[r][c+1] = '|'
				}
			}
		}
	}
	// printGrid(grid)
	return splitCount
}

func getStartPos(grid [][]rune) (int, int) {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 'S' {
				return r, c
			}
		}
	}
	fmt.Println("Start position not found")
	return 0, 0
}

func countTimeline(r, c int, grid [][]rune) int {
	key := [2]int{r, c}
	val, exists := memo[key]
	if exists {
		return val
	}

	if r == len(grid) - 1 {
		return 1
	}

	total := 0

	if grid[r][c] == '^' {
		if c - 1 >= 0 {
			total += countTimeline(r, c - 1, grid)
		}
		if c + 1 < len(grid[r]) {
			total += countTimeline(r, c + 1, grid)
		}
	} else {
		total += countTimeline(r + 1, c, grid)
	}

	memo[key] = total
	return total
}

func main() {
    data, err := os.ReadFile("inputs/day07.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	lines := strings.Split(string(data), "\n")

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	// part1
	fmt.Println("part1:", countSplit(grid))

	// part2
	r, c := getStartPos(grid)
	fmt.Println("part2:", countTimeline(r, c, grid))
}
