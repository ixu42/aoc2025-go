package main

import (
    "fmt"
    "os"
	"strings"
)

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func solvePart1(lines []string) int {
	res := 0
	return res
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

	splitCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'S' ||  grid[i][j] == '|' {
				if i + 1 < len(grid) && grid[i+1][j] == '.' {
					grid[i+1][j] = '|'
				}
			} else if grid[i][j] == '^' && i - 1 >= 0 && grid[i-1][j] == '|' {
				splitCount++
				if j - 1 >= 0 && grid[i][j-1] == '.' {
					grid[i][j-1] = '|'
					if i + 1 < len(grid) && grid[i+1][j-1] == '.' {
						grid[i+1][j-1] = '|'
					}
				}
				if j + 1 < len(grid[i]) && grid[i][j+1] == '.' {
					grid[i][j+1] = '|'
				}
			}
		}
	}

	printGrid(grid)

	fmt.Println("part1:", splitCount)
}
