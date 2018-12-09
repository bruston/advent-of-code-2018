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
		log.Fatalf("unable to open input file: %v", err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	sum := 0
	for s.Scan() {
		n := 0
		if _, err := fmt.Sscanf(s.Text(), "%d", &n); err != nil {
			log.Fatalf("error while scanning input file: %s", err)
		}
		sum += n
	}
	if err := s.Err(); err != nil {
		log.Fatal("error reading from input file: %s", err)
	}
	fmt.Println(sum)
}
