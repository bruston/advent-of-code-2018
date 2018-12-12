package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type point struct{ x, y int }

type distance struct{ pointIndex, n int }

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	points := []point{}
	for s.Scan() {
		p := point{}
		if _, err := fmt.Sscanf(s.Text(), "%d, %d", &p.x, &p.y); err != nil {
			log.Fatalf("error scanning input file: %v", err)
		}
		points = append(points, p)
	}
	if err := s.Err(); err != nil {
		log.Fatalf("error reading input file: %v", err)
	}

	sort.Slice(points, func(i, j int) bool { return points[i].x < points[j].x })
	xMin := points[0].x
	xMax := points[len(points)-1].x
	sort.Slice(points, func(i, j int) bool { return points[i].y < points[j].y })
	yMin := points[0].y
	yMax := points[len(points)-1].y

	grid := make([][]int, xMax-xMin)
	for i := range grid {
		grid[i] = make([]int, yMax-yMin)
	}

	for x := range grid {
		for y := range grid[x] {
			p := point{x: x + xMin, y: y + yMin}
			distances := []distance{}
			for i, v := range points {
				d := abs(p.x-v.x) + abs(p.y-v.y)
				distances = append(distances, distance{pointIndex: i, n: d})
			}

			sort.Slice(distances, func(i, j int) bool { return distances[i].n < distances[j].n })
			if distances[0].n == distances[1].n {
				grid[x][y] = -1
				continue
			}
			grid[x][y] = distances[0].pointIndex
		}
	}

	count := make(map[int]int)
	infinite := make(map[int]struct{})
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] > 0 {
				count[grid[x][y]]++
			}
			if x+xMin == xMin || x+xMin == xMax || y+yMin == yMin || y+yMin == yMax {
				infinite[grid[x][y]] = struct{}{}
			}
		}
	}

	max := 0
	for k, v := range count {
		if v > max {
			if _, ok := infinite[k]; ok {
				continue
			}
			max = v
		}
	}
	fmt.Println(max)
}
