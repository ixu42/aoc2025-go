package main

import (
	"os"
	"log"
	"bufio"
	"io"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"github.com/draffensperger/golp"
)

type machine struct {
	target	[]int
	buttons	[][]int
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

func parseButtons (m *machine, line string) {
	buttonMatches := buttonRe.FindAllStringSubmatch(line, -1)
	m.buttons = make([][]int, 0, len(buttonMatches))
	for _, wm := range buttonMatches {
		nums := parseInts(wm[1])
		m.buttons = append(m.buttons, nums)
	}
}

func parseTargetsAndButtonsPart1(line string) *machine {
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

	parseButtons(m, line)

	return m
}

func parseTargetsAndButtonsPart2(line string) *machine {
	m := &machine{}

	if match := joltageRe.FindStringSubmatch(line); match != nil {
		m.target = parseInts(match[1])
	} else {
		log.Fatal("invalid or missing joltage requirements")
	}

	parseButtons(m, line)

	return m
}

func pressButtonToToggle(state []int, button []int) {
	for _, idx := range button {
		state[idx] ^= 1 // toggle 0->1 or 1->0
	}
}

func solvePart1(file *os.File) int {
	res := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		machine := parseTargetsAndButtonsPart1(line)
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
					pressButtonToToggle(state, machine.buttons[i])
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

	return res
}

// integer linear programming
func minPressesILP(machine *machine) int {
	nButtons := len(machine.buttons)
    nCounters := len(machine.target)

    // constraints/rows = counters, columns = buttons
    lp := golp.NewLP(0, nButtons)

    // build constraints
    for i := 0; i < nCounters; i++ {
        row := make([]float64, nButtons)
        for j := 0; j < nButtons; j++ {
            for _, idx := range machine.buttons[j] {
                if idx == i {
                    row[j] = 1 // set coefficient
                }
            }
        }
        lp.AddConstraint(row, golp.EQ, float64(machine.target[i]))
    }

    // objective: minimize total button presses
    obj := make([]float64, nButtons)
    for j := 0; j < nButtons; j++ {
        obj[j] = 1
        // mark variable j as integer
        lp.SetInt(j, true)
    }
    lp.SetObjFn(obj)
    // default is minimize; no need to call SetMaximize

    // solve
    sol := lp.Solve()
    if sol != golp.OPTIMAL {
        return -1
    }

    // sum rounded results
    vars := lp.Variables()
    total := 0
    for _, v := range vars {
        total += int(v + 0.5) // round to nearest int
    }
    return total
}

func solvePart2(file *os.File) int {
	file.Seek(0, io.SeekStart) // reset file pointer to beginning
	scanner := bufio.NewScanner(file)
	res := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		machine := parseTargetsAndButtonsPart2(line)
		// fmt.Printf("%+v\n", machine)

		minPresses := minPressesILP(machine)

		res += minPresses
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("input error:", err)
	}

	return res
}

func main() {
	file, err := os.Open("inputs/day10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// part1
	fmt.Println("part1:", solvePart1(file))

	// part2
	res := solvePart2(file)
	if res < 0 {
		log.Fatal("no solution found for part2")
	}
	fmt.Println("part2:", res)
}
