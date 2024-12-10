package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func searchTwo(m [][]int, r, c, f int) int {
	if f == 9 {
		return 1
	}

	s := 0
	// Up
	if r > 0 && m[r-1][c] == f+1 {
		s += searchTwo(m, r-1, c, f+1)
	}
	// Down
	if r < len(m)-1 && m[r+1][c] == f+1 {
		s += searchTwo(m, r+1, c, f+1)
	}
	// Left
	if c > 0 && m[r][c-1] == f+1 {
		s += searchTwo(m, r, c-1, f+1)
	}
	// Right
	if c < len(m[0])-1 && m[r][c+1] == f+1 {
		s += searchTwo(m, r, c+1, f+1)
	}

	return s
}

func partTwo() {

	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	in := strings.Split(string(f), "\n")

	m := [][]int{}
	for _, l := range in {
		r := []int{}
		for _, c := range l {
			if c == '.' {
				r = append(r, -1)
				continue
			}
			v, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			r = append(r, v)
		}
		m = append(m, r)
	}

	s := 0
	for r, l := range m {
		for c, v := range l {
			if v == 0 {
				s += searchTwo(m, r, c, 0)
			}
		}
	}

	fmt.Println(s)
}
