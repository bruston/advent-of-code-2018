package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type node struct {
	children []node
	meta     []int
}

func parseNode(r io.Reader) (node, error) {
	var children, metas int
	if _, err := fmt.Fscan(r, &children, &metas); err != nil {
		return node{}, fmt.Errorf("error parsing node header: %v", err)
	}
	parent := node{
		children: make([]node, 0, children),
		meta:     make([]int, 0, metas),
	}
	for i := 0; i < children; i++ {
		child, err := parseNode(r)
		if err != nil {
			return node{}, fmt.Errorf("error parsing child: %v", err)
		}
		parent.children = append(parent.children, child)
	}
	for i := 0; i < metas; i++ {
		n := 0
		if _, err := fmt.Fscan(r, &n); err != nil {
			return node{}, fmt.Errorf("error parsing metadata: %v", err)
		}
		parent.meta = append(parent.meta, n)
	}
	return parent, nil
}

func (nd node) value() int {
	sum := 0
	if len(nd.children) == 0 {
		for _, v := range nd.meta {
			sum += v
		}
		return sum
	}
	for _, v := range nd.meta {
		if v > len(nd.children) {
			continue
		}
		sum += nd.children[v-1].value()
	}
	return sum
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	defer f.Close()
	root, err := parseNode(f)
	if err != nil {
		log.Fatalf("error while parsing nodes: %v", err)
	}
	fmt.Println(root.value())
}
