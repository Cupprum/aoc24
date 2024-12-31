package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cacheTwo = make(map[CK]CV)

func solveTwo(s string, l int) int {
	key := CK{s, l}
	if val, ok := cacheTwo[key]; ok {
		return val.length
	}

	if l > 25 {
		return len(s)
	}

	grid := dp
	if l == 0 {
		grid = np
	}

	sum := 0

	pts := strings.Split("A"+s, "")
	for i, c := range pts {
		if i == 0 {
			continue
		}
		ll := pts[i-1]

		nextPath := findPaths(grid, ll, c)
		sum += solveTwo(nextPath, l+1)
	}

	cacheTwo[key] = CV{path: s, length: sum}
	return sum
}

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	cds := strings.Split(string(f), "\n")

	s := 0
	for _, cd := range cds {
		n, err := strconv.Atoi(strings.TrimSuffix(cd, "A"))
		if err != nil {
			panic(err)
		}

		s += solveTwo(cd, 0) * n
	}
	fmt.Println(s)
}
