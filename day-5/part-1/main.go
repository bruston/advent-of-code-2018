package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	b, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("error reading input file: %v", err)
	}
	b = bytes.TrimSuffix(b, []byte{'\n'})
	var pos int
	for pos != len(b)-1 {
		if b[pos+1] == b[pos]+32 || b[pos+1] == b[pos]-32 {
			b = append(b[:pos], b[pos+2:]...)
			if pos > 0 {
				pos--
			}
			continue
		}
		pos++
	}
	fmt.Println(len(b))
}
