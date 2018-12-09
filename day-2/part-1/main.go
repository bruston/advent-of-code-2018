package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	var two, three int
	for s.Scan() {
		seen := make(map[rune]int)
		for _, v := range s.Text() {
			seen[v]++
		}
		var seenTwo, seenThree bool
		for _, v := range seen {
			if v == 2 && !seenTwo {
				two++
				seenTwo = true
			}
			if v == 3 && !seenThree {
				three++
				seenThree = true
			}
		}
	}
	if err := s.Err(); err != nil {
		log.Fatalf("error reading from input file: %v", err)
	}
	fmt.Println(two * three)
}
