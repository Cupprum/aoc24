package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type P struct {
	r, c int
}

func add(o []P, a []P) []P {
	for _, p := range a {
		if !slices.Contains(o, p) {
			o = append(o, p)
		}
	}
	return o
}

func search(m [][]int, r, c, f int) []P {
	if f == 9 {
		return []P{{r, c}}
	}

	ps := []P{}
	// Up
	if r > 0 && m[r-1][c] == f+1 {
		pps := search(m, r-1, c, f+1)
		ps = add(ps, pps)
	}
	// Down
	if r < len(m)-1 && m[r+1][c] == f+1 {
		pps := search(m, r+1, c, f+1)
		ps = add(ps, pps)
	}
	// Left
	if c > 0 && m[r][c-1] == f+1 {
		pps := search(m, r, c-1, f+1)
		ps = add(ps, pps)
	}
	// Right
	if c < len(m[0])-1 && m[r][c+1] == f+1 {
		pps := search(m, r, c+1, f+1)
		ps = add(ps, pps)
	}

	return ps
}

func partOne() {
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
				s += len(search(m, r, c, 0))
			}
		}
	}

	fmt.Println(s)
}
