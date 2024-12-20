package main

import (
	"fmt"
	"os"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func searchHistTwo(d int, cp P, ep P, m [][]string, hist map[P]int, cm map[int]int) {
	if cp == ep {
		return
	}

	lmt := 20

	for r := -lmt; r <= lmt; r++ {
		for c := -lmt + abs(r); c <= lmt-abs(r); c++ {
			if v, ok := hist[P{cp.r + r, cp.c + c}]; ok && v > d {
				rl := abs(r) + abs(c)
				cm[v-d-rl]++
			}
		}
	}

	if v, ok := hist[P{cp.r - 1, cp.c}]; ok && v == d+1 {
		searchHistTwo(d+1, P{cp.r - 1, cp.c}, ep, m, hist, cm)
	} else if v, ok := hist[P{cp.r, cp.c + 1}]; ok && v == d+1 {
		searchHistTwo(d+1, P{cp.r, cp.c + 1}, ep, m, hist, cm)
	} else if v, ok := hist[P{cp.r + 1, cp.c}]; ok && v == d+1 {
		searchHistTwo(d+1, P{cp.r + 1, cp.c}, ep, m, hist, cm)
	} else if v, ok := hist[P{cp.r, cp.c - 1}]; ok && v == d+1 {
		searchHistTwo(d+1, P{cp.r, cp.c - 1}, ep, m, hist, cm)
	}
}

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lns := strings.Split(string(f), "\n")

	m := [][]string{}
	for _, ln := range lns {
		m = append(m, strings.Split(ln, ""))
	}

	sp := findChar(m, "S")
	ep := findChar(m, "E")

	hist := map[P]int{}

	cm := map[int]int{}

	// Fill history
	search(0, sp, ep, m, hist)

	searchHistTwo(0, sp, ep, m, hist, cm)

	s := 0
	for k, v := range cm {
		if k >= 100 {
			s += v
		}
	}
	fmt.Println(s)
}
