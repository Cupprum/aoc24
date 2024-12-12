package main

import (
	"fmt"
	"os"
	"strings"
)

func searchOne(r int, c int, v string, m [][]string) (int, int) {
	p := 0
	a := 0

	m[r][c] = strings.ToLower(m[r][c])
	if r > 0 {
		if m[r-1][c] == v {
			np, na := searchOne(r-1, c, v, m)
			p += np
			a += na
		} else if m[r-1][c] != strings.ToLower(v) {
			p += 1
		}
	} else {
		p += 1
	}
	if c < len(m)-1 {
		if m[r][c+1] == v {
			np, na := searchOne(r, c+1, v, m)
			p += np
			a += na
		} else if m[r][c+1] != strings.ToLower(v) {
			p += 1
		}
	} else {
		p += 1
	}
	if r < len(m[0])-1 {
		if m[r+1][c] == v {
			np, na := searchOne(r+1, c, v, m)
			p += np
			a += na
		} else if m[r+1][c] != strings.ToLower(v) {
			p += 1
		}
	} else {
		p += 1
	}
	if c > 0 {
		if m[r][c-1] == v {
			np, na := searchOne(r, c-1, v, m)
			p += np
			a += na
		} else if m[r][c-1] != strings.ToLower(v) {
			p += 1
		}
	} else {
		p += 1
	}

	return p, a + 1
}

func partOne() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	in := strings.Split(string(f), "\n")

	m := [][]string{}
	for _, l := range in {
		m = append(m, strings.Split(l, ""))
	}

	s := 0
	for r, rr := range m {
		for c, v := range rr {
			if strings.ToUpper(v) == v {
				p, a := searchOne(r, c, v, m)
				s += p * a
			}
		}
	}

	fmt.Println(s)
}
