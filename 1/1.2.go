package main

import (
	"fmt"
	"io/ioutil"
)

const (
	open  = 0x28
	close = 0x29
)

func main() {
	input, err := ioutil.ReadFile("./input")
	if err != nil {
		panic("Unable to read input file")
	}

	floor := 0
	for i, b := range input {
		switch b {
		case open:
			floor++
		case close:
			floor--
		}

		if floor == -1 {
			fmt.Println("First character to lead to floor -1 is ", i+1)
			break
		}
	}
}
