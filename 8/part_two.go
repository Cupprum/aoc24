package main

import (
	"fmt"
	"os"
	"strings"
)

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	ls := strings.Split(string(f), "\n")

	fs := map[rune][]F{}

	for r, l := range ls {
		for c, f := range l {
			if f == '.' {
				continue
			}
			if _, ok := fs[f]; !ok {
				fs[f] = []F{}
			}
			fs[f] = append(fs[f], F{r: r, c: c})
		}
	}

	m := make([]string, len(ls))
	for i := range m {
		m[i] = strings.Repeat(".", len(ls[0]))
	}

	for _, f := range fs {
		for _, v1 := range f {
			m[v1.r] = m[v1.r][:v1.c] + "#" + m[v1.r][v1.c+1:]
			for _, v2 := range f {
				if v1 == v2 {
					continue
				}

				r := v2.r + (v2.r - v1.r)
				c := v2.c + (v2.c - v1.c)
				for r >= 0 && r < len(ls) && c >= 0 && c < len(ls[0]) {
					m[r] = m[r][:c] + "#" + m[r][c+1:]
					r = r + (v2.r - v1.r)
					c = c + (v2.c - v1.c)
				}
			}
		}
	}

	c := 0
	for _, l := range m {
		c += strings.Count(l, "#")
	}
	fmt.Println(c)
}
