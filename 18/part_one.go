package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const mr int = 70
const mc int = 70

type P struct {
	r int
	c int
}

func visit(m [][]string, vstd map[P]int, p P, d int) int {
	if p.r == mr+1 && p.c == mc+1 {
		return d
	}

	if v, ok := vstd[p]; ok && v < d {
		return -1
	}
	vstd[p] = d

	mw := -1
	if m[p.r-1][p.c] != "#" {
		if v, ok := vstd[P{r: p.r - 1, c: p.c}]; !ok || v > d+1 {
			w := visit(m, vstd, P{r: p.r - 1, c: p.c}, d+1)
			if w != -1 && (mw == -1 || w < mw) {
				mw = w
			}
		}
	}
	if m[p.r][p.c+1] != "#" {
		if v, ok := vstd[P{r: p.r, c: p.c + 1}]; !ok || v > d+1 {
			w := visit(m, vstd, P{r: p.r, c: p.c + 1}, d+1)
			if w != -1 && (mw == -1 || w < mw) {
				mw = w
			}
		}
	}
	if m[p.r+1][p.c] != "#" {
		if v, ok := vstd[P{r: p.r + 1, c: p.c}]; !ok || v > d+1 {
			w := visit(m, vstd, P{r: p.r + 1, c: p.c}, d+1)
			if w != -1 && (mw == -1 || w < mw) {
				mw = w
			}
		}
	}
	if m[p.r][p.c-1] != "#" {
		if v, ok := vstd[P{r: p.r, c: p.c - 1}]; !ok || v > d+1 {
			w := visit(m, vstd, P{r: p.r, c: p.c - 1}, d+1)
			if w != -1 && (mw == -1 || w < mw) {
				mw = w
			}
		}
	}

	return mw
}

func partOne() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lns := strings.Split(string(f), "\n")

	bts := []P{}

	for _, ln := range lns {
		pts := strings.Split(ln, ",")
		x, err := strconv.Atoi(pts[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(pts[1])
		if err != nil {
			panic(err)
		}
		bts = append(bts, P{r: y + 1, c: x + 1})
	}

	m := [][]string{}
	for r := 0; r <= mr+2; r++ {
		m = append(m, make([]string, mc+3))
	}
	for c := 0; c <= mc+2; c++ {
		m[0][c] = "#"
		m[mr+2][c] = "#"
	}
	for r := 1; r <= mr+1; r++ {
		m[r][0] = "#"
		for c := 1; c <= mc+2; c++ {
			m[r][c] = "."
		}
		m[r][mc+2] = "#"
	}

	for i := 0; i < 1024; i++ {
		e := bts[i]
		m[e.r][e.c] = "#"
	}

	vstd := map[P]int{}
	fmt.Println(visit(m, vstd, P{r: 1, c: 1}, 0))
}
