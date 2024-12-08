package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func findGuard(mp []string) (int, int) {
	for r, v := range mp {
		for _, p := range []string{"V", "<", ">", "^"} {
			if c := strings.Index(v, p); c != -1 {
				return r, c
			}
		}
	}
	return -1, -1
}

func countX(mp []string) int {
	s := 1

	for _, v := range mp {
		s += strings.Count(v, "X")
	}
	return s
}

type P struct {
	r int
	c int
}

func rot(p P) P {
	return P{p.c, -p.r}
}

func partOne() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	mp := strings.Split(string(f), "\n")

	r, c := findGuard(mp)

	p := P{-1, 0}
	d := P{r, c}
	h := []P{}

	for within(mp, d.r, d.c) {
		mp[d.r] = mp[d.r][:d.c] + "X" + mp[d.r][d.c+1:]
		dest := P{d.r + p.r, d.c + p.c}
		if mp[dest.r][dest.c] == '#' {
			p = rot(p)
		} else {
			d = dest
		}
		if slices.Contains(h, p) {
			c++
			break
		}
		h = append(h, d)
	}
	fmt.Println(countX(mp))
}
