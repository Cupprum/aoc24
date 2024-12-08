package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type H struct {
	p P
	d P
}

func within(mp []string, r, c int) bool {
	return r > 0 && r < len(mp)-1 && c > 0 && c < len(mp[0])-1
}

func subSearch(mp []string, p P, d P, b P) bool {
	h := []H{}

	for within(mp, d.r, d.c) {
		dest := P{d.r + p.r, d.c + p.c}
		if mp[dest.r][dest.c] == '#' || (dest.r == b.r && dest.c == b.c) {
			p = rot(p)
			if slices.Contains(h, H{p, d}) {
				return true
			}
			h = append(h, H{p, d})
		} else {
			d = dest
		}
	}

	return false
}

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	mp := strings.Split(string(f), "\n")

	r, c := findGuard(mp)

	counter := 0

	p := P{-1, 0}
	d := P{r, c}

	h := []P{}

	for within(mp, d.r, d.c) {
		dest := P{d.r + p.r, d.c + p.c}
		if mp[dest.r][dest.c] == '#' {
			p = rot(p)
		} else {
			if !slices.Contains(h, dest) {
				nmp := make([]string, len(mp))
				copy(nmp, mp)
				if subSearch(nmp, rot(p), d, dest) {
					counter++
				}
				h = append(h, dest)
			}
			d = dest
		}
	}

	fmt.Println(counter)
}
