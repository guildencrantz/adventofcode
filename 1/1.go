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
	for _, b := range input {
		switch b {
		case open:
			floor++
		case close:
			floor--
		}
	}

	fmt.Println("Final floor: ", floor)

}
