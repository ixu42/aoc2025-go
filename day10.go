package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type machine struct {
	target	[]int
	buttons	[][]int
	joltage	[]int
}

var (
	targetRe	= regexp.MustCompile(`\[(.*?)\]`)
	buttonRe    = regexp.MustCompile(`\((.*?)\)`)
	joltageRe   = regexp.MustCompile(`\{(.*?)\}`)
)

func parseInts(s string) []int {
	if strings.TrimSpace(s) == "" {
		log.Fatal("empty integer list")
	}
	parts := strings.Split(s, ",")
	nums := make([]int, 0, len(parts))
	for _, p := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, n)
	}
	return nums
}

func parseLine(line string) *machine {
	m := &machine{}

	if match := targetRe.FindStringSubmatch(line); match != nil {
		targetStr := match[1]
		m.target = make([]int, len(targetStr))
		for i, c := range targetStr {
			if c == '#' {
				m.target[i] = 1
			} else {
				m.target[i] = 0
			}
		}
	} else {
		log.Fatal("invalid or missing target state")
	}

	buttonMatches := buttonRe.FindAllStringSubmatch(line, -1)
	for _, wm := range buttonMatches {
		nums := parseInts(wm[1])
		m.buttons = append(m.buttons, nums)
	}

	if match := joltageRe.FindStringSubmatch(line); match != nil {
		nums := parseInts(match[1])
		m.joltage = nums
	} else {
		log.Fatal("invalid or missing joltage requirements")
	}

	return m
}

func pressButton(state []int, button []int) {
	for _, idx := range button {
		state[idx] ^= 1 // toggle 0->1 or 1->0
	}
}

func main() {
	file, err := os.Open("inputs/day10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// part1
	res := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		machine := parseLine(line)

		// fmt.Printf("%+v\n", machine)

		nButtons := len(machine.buttons)
		nLights := len(machine.target)
		minPresses := nButtons + 1 // max + 1 possibilities

		// brute force all button press combinations
		// binary mask to present button press combination
		// 00 -> 0 -> no buttons pressed
		// 01 -> 1 -> button 0 pressed
		// 10 -> 2 -> button 1 pressed
		// 11 -> 3 -> buttons 0 and 1 pressed
		for mask := 0; mask < (1 << nButtons); mask++ {
			state := make([]int, nLights)
			pressCount := 0
			for i := 0; i < nButtons; i++ {
				// check if button i is pressed
				if (mask >> i) & 1 == 1 {
					pressButton(state, machine.buttons[i])
					pressCount++
				}
			}

			match := true
			for j := 0; j < nLights; j++ {
				if state[j] != machine.target[j] {
					match = false
					break
				}
			}
			if match && pressCount < minPresses {
				minPresses = pressCount
			}
		}
		// fmt.Println("minimum button presses:", minPresses)
		res += minPresses
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("input error:", err)
	}

	fmt.Println("part1:", res)
}
