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

	deliveries := Santa(input)
	fmt.Printf("%d houses got presents\n", len(deliveries))

	deliveries = RoboSanta(input)
	fmt.Printf("%d houses got presents with robosanta helping\n", len(deliveries))
}

func Santa(input []byte) map[House]bool {
	deliveries := make(map[House]bool)
	l := House{0, 0}
	deliveries[l] = true
	for _, d := range input {
		l = l.Move(d)
		deliveries[l] = true
	}

	return deliveries
}

func RoboSanta(input []byte) map[House]bool {
	deliveries := make(map[House]bool)

	s := House{0, 0}
	deliveries[s] = true

	r := s

	for i, d := range input {
		if i%2 == 0 {
			s = s.Move(d)
			deliveries[s] = true
		} else {
			r = r.Move(d)
			deliveries[r] = true
		}
	}
	return deliveries
}
