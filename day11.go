package main

import (
    "fmt"
    "os"
	"strings"
	"log"
)

var graph map[string][]string
var memo map[string]int // key:current-containsFFT-containsDAC, value:pathCount

func countPathPart1(current string) int {
	if current == "out" {
		return 1
	}

	totalPaths := 0
	for _, next := range graph[current] {
		totalPaths += countPathPart1(next)
	}
	return totalPaths
}

// count paths from current to "out" that contain both "fft" and "dac"
func countPathPart2(current string, containsFFT, containsDAC bool) int {
	key := fmt.Sprintf("%s-%t-%t", current, containsFFT, containsDAC)
	if pathCount, found := memo[key]; found {
		return pathCount
	}

	if current == "out" {
		if containsFFT && containsDAC {
			return 1
		}
		return 0
	}

	totalPaths := 0
	for _, next := range graph[current] {
		newFFT := containsFFT
		newDAC := containsDAC
		if next == "fft" {
			newFFT = true
		}
		if next == "dac" {
			newDAC = true
		}
		totalPaths += countPathPart2(next, newFFT, newDAC)
	}

	memo[key] = totalPaths
	return totalPaths
}

func main() {
    data, err := os.ReadFile("inputs/day11.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	lines := strings.Split(string(data), "\n")

	graph = make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			log.Fatalf("Invalid line format: %s", line)
		}
		device := strings.TrimSpace(parts[0])
		outputs := strings.TrimSpace(parts[1])
		outputList := strings.Fields(outputs)
		graph[device] = outputList
	}

	// for device, outs := range graph {
	// 	fmt.Println(device, "->", outs)
	// }

	// fmt.Println("part1:", countPathPart1("you"))
	memo = make(map[string]int)
	fmt.Println("part2:", countPathPart2("svr", false, false))
}
