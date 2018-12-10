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
	ids := []string{}
	for s.Scan() {
		ids = append(ids, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatalf("error reading from input file: %v", err)
	}
	correct := []rune{}
outer:
	for _, id := range ids {
		id1 := []rune(id)
		for _, v1 := range ids {
			id2 := []rune(v1)
			diff := 0
			for i, _ := range id1 {
				if id1[i] != id2[i] {
					diff++
				} else {
					correct = append(correct, id1[i])
				}
			}
			if diff == 1 {
				break outer
			}
			correct = []rune{}
		}
	}
	fmt.Printf("%s\n", string(correct))
}
