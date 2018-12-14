package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	b, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}
	sum := 0
	seen := make(map[int]struct{})
	for {
		s := bufio.NewScanner(bytes.NewReader(b))
		for s.Scan() {
			n := 0
			if _, err := fmt.Sscan(s.Text(), &n); err != nil {
				log.Fatalf("error while scanning input: %s", err)
			}
			sum += n
			if _, ok := seen[sum]; ok {
				fmt.Println(sum)
				return
			}
			seen[sum] = struct{}{}
		}
		if err := s.Err(); err != nil {
			log.Fatal("error reading from input: %s", err)
		}
	}
}
