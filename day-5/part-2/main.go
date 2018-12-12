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
	best := len(b)
	for i := byte(65); i <= 90; i++ {
		c := bytes.Replace(b, []byte{i}, []byte(""), -1)
		c = bytes.Replace(c, []byte{i + 32}, []byte(""), -1)
		var pos int
		for pos != len(c)-1 {
			if c[pos+1] == c[pos]+32 || c[pos+1] == c[pos]-32 {
				c = append(c[:pos], c[pos+2:]...)
				if pos > 0 {
					pos--
				}
				continue
			}
			pos++
		}
		if len(c) < best {
			best = len(c)
		}
	}
	fmt.Println(best)
}
