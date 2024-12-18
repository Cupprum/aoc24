package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partTwo() {
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

	for i := 0; i < mc*mr; i++ {
		e := bts[i]
		m[e.r][e.c] = "#"

		if i > 2048 {
			vstd := map[P]int{}
			w := visit(m, vstd, P{r: 1, c: 1}, 0)
			if w == -1 {
				fmt.Print(e.c - 1)
				fmt.Print(",")
				fmt.Println(e.r - 1)
				break
			}
		}
	}
}
