package main

import (
    "fmt"
    "os"
	"strings"
	"strconv"
)

type point struct {
	r, c int
}

var redTiles []point // vertices of the polygon

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func loadRedTiles(lines []string) {
	for _, line := range lines {
		elems := strings.Split(line, ",")
		if len(elems) != 2 {
			fmt.Println("invalid input:", line)
			continue
		}
		c, errC := strconv.Atoi(elems[0])
		r, errR := strconv.Atoi(elems[1])
		if errC != nil || errR != nil {
			fmt.Println("invalid coordinates:", line)
			continue
		}
		redTiles = append(redTiles, point{r, c})
	}
}

func solvePart1() int {
	maxSize := 0
	rectCount := 0
	for i := 0; i < len(redTiles); i++ {
		for j := i + 1; j < len(redTiles); j++ {
			rectCount++
			size := (abs(redTiles[i].r - redTiles[j].r) + 1) * (abs(redTiles[i].c - redTiles[j].c) + 1)
			if size > maxSize {
				maxSize = size
			}
		}
	}

	return maxSize
}

// cast a ray to the right, count intersections with polygon edges
// if odd -> inside, even -> outside
func isPointInPolygon(p point) bool {
	intersections := 0
	n := len(redTiles)
	
	for i := 0; i < n; i++ {
		v1 := redTiles[i]
		v2 := redTiles[(i+1)%n]

		// horizontal edges of the polygon
		if v1.r == v2.r {
			// check if point is on the edge
			if p.r == v1.r &&
			p.c >= min(v1.c, v2.c) && p.c <= max(v1.c, v2.c) {
				return true
			}
			continue
		}

		// vertical edges of the polygon
		if v1.r > v2.r {
			v1, v2 = v2, v1
		}

		if p.r >= v1.r && p.r < v2.r {
			xIntersection := v1.c
			if xIntersection > p.c {
				intersections++
			}
		}
	}
	return intersections % 2 == 1
}

func isHorizontal(r1, c1, r2, c2 int) bool {
    return r1 == r2
}

func isVertical(r1, c1, r2, c2 int) bool {
    return c1 == c2
}

func rectEdgeIntersectsPolygonEdge(r1, c1, r2, c2 int) bool {
    rectEdges := [][4]int{
        {r1, c1, r1, c2}, // top
        {r2, c1, r2, c2}, // bottom
        {r1, c1, r2, c1}, // left
        {r1, c2, r2, c2}, // right
    }

    for _, re := range rectEdges {
        rr1, rc1, rr2, rc2 := re[0], re[1], re[2], re[3]

        for i := 0; i < len(redTiles); i++ {
            p1 := redTiles[i]
            p2 := redTiles[(i+1)%len(redTiles)]

            pr1, pc1 := p1.r, p1.c
            pr2, pc2 := p2.r, p2.c

            // horizontal–vertical
            if isHorizontal(rr1, rc1, rr2, rc2) && isVertical(pr1, pc1, pr2, pc2) {
                if pc1 > min(rc1, rc2) && pc1 < max(rc1, rc2) &&
                   rr1 > min(pr1, pr2) && rr1 < max(pr1, pr2) {
                    return true
                }
            }

            // vertical–horizontal
            if isVertical(rr1, rc1, rr2, rc2) && isHorizontal(pr1, pc1, pr2, pc2) {
                if rc1 > min(pc1, pc2) && rc1 < max(pc1, pc2) &&
                   pr1 > min(rr1, rr2) && pr1 < max(rr1, rr2) {
                    return true
                }
            }
        }
    }
    return false
}

// the rectangle is valid if
// 1. it's 4 corners are all inside the polygon, and
// 2. it's 4 edges do not intersect with any polygon edge
func isValidRect(i, j int) bool {
	r1 := redTiles[i].r
	c1 := redTiles[i].c
	r2 := redTiles[j].r
	c2 := redTiles[j].c

	// normalize coordinates
    if r1 > r2 {
		r1, r2 = r2, r1
	}
	if c1 > c2 {
		c1, c2 = c2, c1
	}

	corners := []point{
		{r1, c1},
		{r1, c2},
		{r2, c1},
		{r2, c2},
	}
	for _, corner := range corners {
		if !isPointInPolygon(corner) {
			return false
		}
	}
	if rectEdgeIntersectsPolygonEdge(r1, c1, r2, c2) {
		return false
	}
	return true
}

func solvePart2() int {
	maxValidSize := 0
	for i := 0; i < len(redTiles); i++ {
		for j := i + 1; j < len(redTiles); j++ {
			size := (abs(redTiles[i].r - redTiles[j].r) + 1) * (abs(redTiles[i].c - redTiles[j].c) + 1)
			if size <= maxValidSize {
				continue
			}
			if isValidRect(i, j) {
				maxValidSize = size
			}
		}
	}
	return maxValidSize
}

func main() {
    data, err := os.ReadFile("inputs/day09.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	lines := strings.Split(string(data), "\n")

	redTiles = make([]point, 0, len(lines))
	loadRedTiles(lines)

	fmt.Printf("part1: %d\n", solvePart1())
	fmt.Printf("part2: %d\n", solvePart2())
}
