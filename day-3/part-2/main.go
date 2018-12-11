package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type square struct{ x, y int }

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	claims := make(map[square]int)
	squares := make(map[int][]*square)
	for s.Scan() {
		var id, xOffset, yOffset, width, height int
		if _, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &id, &xOffset, &yOffset, &width, &height); err != nil {
			log.Fatalf("error while scanning input file: %v", err)
		}
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				sq := square{
					x: xOffset + x,
					y: yOffset + y,
				}
				claims[sq]++
				squares[id] = append(squares[id], &sq)
			}
		}
	}
	if err := s.Err(); err != nil {
		log.Fatalf("error reading input file: %v", err)
	}
	for id, s := range squares {
		conflict := false
		for _, v := range s {
			if n := claims[*v]; n > 1 {
				conflict = true
				break
			}
		}
		if !conflict {
			fmt.Println(id)
		}
	}
}
