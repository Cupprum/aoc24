package main

import (
	"fmt"
	"os"
	"strings"
)

type P struct {
	r int
	c int
}

func findChar(m [][]string, ch string) P {
	for r, row := range m {
		for c, col := range row {
			if col == ch {
				return P{r, c}
			}
		}
	}
	return P{-1, -1}
}

func search(d int, cp P, ep P, m [][]string, hist map[P]int) int {
	if cp == ep {
		hist[cp] = d
		return 0
	}

	if v, ok := hist[cp]; ok && v < d {
		return -1
	}
	hist[cp] = d

	mv := -1

	dir := func(p P) {
		if v, ok := hist[p]; !ok || v > d+1 {
			tmv := search(d+1, p, ep, m, hist)
			if tmv != -1 && (mv == -1 || tmv+1 < mv) {
				mv = tmv + 1
			}
		}
	}

	if m[cp.r-1][cp.c] != "#" {
		dir(P{cp.r - 1, cp.c})
	}
	if m[cp.r][cp.c+1] != "#" {
		dir(P{cp.r, cp.c + 1})
	}
	if m[cp.r+1][cp.c] != "#" {
		dir(P{cp.r + 1, cp.c})
	}
	if m[cp.r][cp.c-1] != "#" {
		dir(P{cp.r, cp.c - 1})
	}

	return mv
}

func searchHistOne(d int, cp P, ep P, m [][]string, hist map[P]int, cm map[int]int) {
	if cp == ep {
		return
	}

	if m[cp.r-1][cp.c] == "#" && cp.r-2 >= 0 {
		if v, ok := hist[P{cp.r - 2, cp.c}]; ok && v > d {
			cm[v-d-2]++
		}
	}
	if m[cp.r][cp.c+1] == "#" && cp.c+2 < len(m[0]) {
		if v, ok := hist[P{cp.r, cp.c + 2}]; ok && v > d {
			cm[v-d-2]++
		}
	}
	if m[cp.r+1][cp.c] == "#" && cp.r+2 < len(m) {
		if v, ok := hist[P{cp.r + 2, cp.c}]; ok && v > d {
			cm[v-d-2]++
		}
	}
	if m[cp.r][cp.c-1] == "#" && cp.c-2 >= 0 {
		if v, ok := hist[P{cp.r, cp.c - 2}]; ok && v > d {
			cm[v-d-2]++
		}
	}

	if v, ok := hist[P{cp.r - 1, cp.c}]; ok && v == d+1 {
		searchHistOne(d+1, P{cp.r - 1, cp.c}, ep, m, hist, cm)
	} else if v, ok := hist[P{cp.r, cp.c + 1}]; ok && v == d+1 {
		searchHistOne(d+1, P{cp.r, cp.c + 1}, ep, m, hist, cm)
	} else if v, ok := hist[P{cp.r + 1, cp.c}]; ok && v == d+1 {
		searchHistOne(d+1, P{cp.r + 1, cp.c}, ep, m, hist, cm)
	} else if v, ok := hist[P{cp.r, cp.c - 1}]; ok && v == d+1 {
		searchHistOne(d+1, P{cp.r, cp.c - 1}, ep, m, hist, cm)
	}
}

func partOne() {
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

	search(0, sp, ep, m, hist)

	searchHistOne(0, sp, ep, m, hist, cm)

	s := 0
	for k, v := range cm {
		if k >= 37 {
			s += v
		}
	}
	fmt.Println(s)
}
