package main

import (
    "fmt"
    "os"
	"strings"
)

var dirs = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func copyGrid(grid [][]rune) [][]rune {
	newGrid := make([][]rune, len(grid))
	for i := range grid {
		newGrid[i] = make([]rune, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	return newGrid
}

func solvePart1(grid [][]rune) (int, [][]rune) {
	removeCount := 0
	newGrid := copyGrid(grid)
	// printGrid(grid)
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] != '@' {
				continue
			}
			count := 0
			for _, d := range dirs {
				nr := r + d[0]
				nc := c + d[1]

				if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) {
					neighbor := grid[nr][nc]
					if neighbor == '@' {
						count++
					}
				}
			}
			if count < 4 {
				removeCount += 1
				newGrid[r][c] = 'x'
			}
		}
	}
	// printGrid(newGrid)
	return removeCount, newGrid
}

func solvePart2(grid [][]rune) int {
	removeTotal := 0
	for {
		removeCount, newGrid := solvePart1(grid)
		if removeCount == 0 {
			break
		}
		removeTotal += removeCount
		grid = newGrid
	}
	return removeTotal
}

func main() {
    data, err := os.ReadFile("inputs/day04.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	lines := strings.Split(string(data), "\n")

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = make([]rune, len(line))
		grid[i] = []rune(line)
	}

	part1Res, _ := solvePart1(grid)
	fmt.Println("part1:", part1Res)
	fmt.Println("part2:", solvePart2(grid))
}
