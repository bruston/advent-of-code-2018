package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func parseNode(r io.Reader) (int, error) {
	var children, metas int
	if _, err := fmt.Fscan(r, &children, &metas); err != nil {
		return 0, err
	}
	sum := 0
	for i := 0; i < children; i++ {
		n, err := parseNode(r)
		if err != nil {
			return 0, err
		}
		sum += n
	}
	for i := 0; i < metas; i++ {
		var n int
		if _, err := fmt.Fscan(r, &n); err != nil {
			return 0, err
		}
		sum += n
	}
	return sum, nil
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer f.Close()
	sum, err := parseNode(f)
	if err != nil {
		log.Fatalf("error parsing nodes: %v", err)
	}
	fmt.Println(sum)
}
