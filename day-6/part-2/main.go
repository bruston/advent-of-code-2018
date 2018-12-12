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

	safe := 0
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			p := point{x: x, y: y}
			sum := 0
			const maxDistance = 10000
			for _, v := range points {
				sum += abs(p.x-v.x) + abs(p.y-v.y)
			}
			if sum < maxDistance {
				safe++
			}
		}
	}
	fmt.Println(safe)
}
