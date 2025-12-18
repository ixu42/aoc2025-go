package main

import (
    "fmt"
    "os"
	"strings"
	"strconv"
	"math"
	"sort"
)

var parent []int
var size []int

// box position in 3D space
type box struct{
	x, y, z float64
}

// connection between 2 boxes, with distance
type connection struct {
	a, b int
	dist float64
}

func getDistance(a, b box) float64 {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// find root / representative of the connected boxes
func find(x int) int {
	if parent[x] != x {
		parent[x] = find(parent[x])
	}
	return parent[x]
}

func union(x, y int) {
	rootX := find(x)
	rootY := find(y)

	// if in the same tree, do nothing
	if rootX == rootY {
		return
	}

	// attach smaller tree to larger tree
	if size[rootX] < size[rootY] {
		rootX, rootY = rootY, rootX
	}

	parent[rootY] = rootX
	size[rootX] += size[rootY]
}

func ifAllConnected() bool {
	root := find(0)
	for i := 1; i < len(parent); i++ {
		if find(i) != root {
			return false
		}
	}
	return true
}

func main() {
    data, err := os.ReadFile("inputs/day08.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	lines := strings.Split(string(data), "\n")

	// get a list of positions of all junction boxes
	boxes := make([]box, 0)
	for _, line := range lines {
		vals := strings.Split(line, ",")
		if len(vals) != 3 {
			fmt.Println("Invalid line format:", line)
			return
		}
		xStr, yStr, zStr := vals[0], vals[1], vals[2]
		x, errX := strconv.ParseFloat(xStr, 64)
		y, errY := strconv.ParseFloat(yStr, 64)
		z, errZ := strconv.ParseFloat(zStr, 64)
		if errX != nil || errY != nil || errZ != nil {
			fmt.Println("Error converting string to int:", errX, errY, errZ)
			return
		}
		boxes = append(boxes, box{x, y, z})
	}

	// get connections between each pair of boxes
	conns := make([]connection, 0)
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			dist := getDistance(boxes[i], boxes[j])
			// fmt.Printf("Distance between pos %d and pos %d: %.2f\n", i, j, dist)
			conns = append(conns, connection{i, j, dist})
		}
	}

	// from the closest to the furthest
	sort.Slice(conns, func(a, b int) bool {
		return conns[a].dist < conns[b].dist
	})

	parent = make([]int, len(boxes))
	size = make([]int, len(boxes))
	for i := 0; i < len(boxes); i++ {
		parent[i] = i // each box starts in its own circuit
		size[i] = 1 // each circuit starts with size 1
	}

	// limit := 1000
	// if limit > len(conns) {
	// 	limit = len(conns)
	// }
	// for k := 0; k < limit; k++ {
	lastPair := make([]int, 2)
	for k := 0; k < len(conns); k++ {
		conn := conns[k]
		// fmt.Printf("box %d and box %d: distance %.2f\n", conn.a, conn.b, conn.dist)
		// if not already connected, connect them and if all connected, stop
		if find(conn.a) != find(conn.b) && !ifAllConnected() {
			union(conn.a, conn.b)
			lastPair[0] = conn.a
			lastPair[1] = conn.b
			// fmt.Printf("Connected box %d and box %d\n", conn.a, conn.b)
		}
	}

	// count circuit sizes
	circuitCounts := make(map[int]int)
	for i := 0; i < len(boxes); i++ {
		root := find(i)
		circuitCounts[root]++
	}

	// for root, count := range circuitCounts {
	// 	fmt.Printf("circuit root %d has size %d\n", root, count)
	// }

	// collect and sort sizes
	sizes := make([]int, 0, len(circuitCounts))
	for _, s := range circuitCounts {
		sizes = append(sizes, s)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	res := 1
	for i := 0; i < 3 && i < len(sizes); i++ {
		res *= sizes[i]
	}

	// fmt.Printf("part1: %d\n", res)

	fmt.Printf("part2: %d\n", int(boxes[lastPair[0]].x*boxes[lastPair[1]].x))

}
