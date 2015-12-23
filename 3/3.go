package main

import (
	"fmt"
	"io/ioutil"
)

const (
	north = 0x5E
	south = 0x76
	east  = 0x3C
	west  = 0x3E
)

type House struct {
	lat  int
	long int
}

func (h House) Move(d byte) House {
	var n House
	switch d {
	case north:
		n = House{h.lat + 1, h.long}
	case south:
		n = House{h.lat - 1, h.long}
	case east:
		n = House{h.lat, h.long + 1}
	case west:
		n = House{h.lat, h.long - 1}
	}

	return n
}

func main() {
	input, err := ioutil.ReadFile("./input")
	if err != nil {
		panic("Unable to read input file")
	}

	deliveries := make(map[House]bool)
	l := House{0, 0}
	deliveries[l] = true
	for _, d := range input {
		l = l.Move(d)
		deliveries[l] = true
	}

	fmt.Printf("%d houses got presents", len(deliveries))
}
