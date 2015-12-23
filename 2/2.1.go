package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Package []int

func NewPackage(s string) Package {
	p := Package(make([]int, 3))
	for i, d := range strings.Split(s, "x") {
		p[i], _ = strconv.Atoi(d)
	}
	sort.IntSlice(p).Sort()
	return p
}

func PaperNeeded(p Package) int {
	return p[0]*p[1] + // Smallest side bonus paper
		p[0]*p[1]*2 +
		p[0]*p[2]*2 +
		p[1]*p[2]*2
}

func main() {
	input, err := ioutil.ReadFile("./input")
	if err != nil {
		panic("Unable to read input")
	}

	paperNeeded := 0
	for _, s := range strings.Split(string(input), "\n") {
		p := NewPackage(s)
		paperNeeded += PaperNeeded(p)
	}

	fmt.Printf("Paper needed: %d sqft", paperNeeded)
}
