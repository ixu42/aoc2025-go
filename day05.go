package main

import (
    "fmt"
    "os"
	"strings"
	"strconv"
	"sort"
)

type Range struct {
	start, end int
}

func isFresh(id int, ranges []Range) bool {
	for _, r := range ranges {
		if id >= r.start && id <= r.end {
			return true
		}
	}
	return false
}

func main() {
    input, err := os.ReadFile("inputs/day05.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	var ranges []Range
	var ids []int

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			start, err1 := strconv.Atoi(parts[0])
			end, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				fmt.Println("Error converting to integer:", err1, err2)
				return
			}
			ranges = append(ranges, Range{start: start, end: end})
		} else {
			val, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				return
			}
			ids = append(ids, val)
		}
	}
	// fmt.Println("Ranges:", ranges)
	// fmt.Println("Ints:", ids)

	// part1
	// res := 0
	// for _, id := range ids {
	// 	if isFresh(id, ranges) {
	// 		res++
	// 	}
	// }
	// fmt.Println("part1:", res)

	// part2
	resultRanges := []Range{}
	resultRanges = append(resultRanges, ranges[0])

	for i := range ranges[1:] {
		r := &ranges[i+1]
		toAppend := true
		toDelete := []int{}

		for j := range resultRanges {
			rr := &resultRanges[j]

			// r fully covers rr
			if r.start < rr.start && r.end > rr.end {
				toDelete = append(toDelete, j)
				continue
			}
			// left overlap
			if r.start < rr.start && r.end >= rr.start && r.end <= rr.end {
				r.end = rr.end
				toDelete = append(toDelete, j)
				continue
			}
			// right overlap
			if r.start >= rr.start && r.start <= rr.end && r.end > rr.end {
				r.start = rr.start
				toDelete = append(toDelete, j)
				continue
			}
			// r is inside rr -> no need to append
			if r.start >= rr.start && r.end <= rr.end {
				toAppend = false
				break
			}
		}

		// delete in descending order to avoid messing up indices
		sort.Slice(toDelete, func(i, j int) bool {
			return toDelete[i] > toDelete[j]
		})
		for _, idx := range toDelete {
			resultRanges = append(resultRanges[:idx], resultRanges[idx+1:]...)
		}

		// append r to resultRanges
		if toAppend {
			resultRanges = append(resultRanges, *r)
		}
	}

	res := 0
	for _, rr := range resultRanges {
		res += rr.end - rr.start + 1
	}
	fmt.Println("part2:", res)
}
