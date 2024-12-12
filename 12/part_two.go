package main

import (
	"fmt"
	"os"
	"strings"
)

func searchTwo(r int, c int, v string, m [][]string) (int, int) {
	p := 0
	a := 0

	m[r][c] = strings.ToLower(m[r][c])

	// for _, l := range m {
	// 	fmt.Println(l)
	// }
	// fmt.Println()

	if m[r-1][c] != v && m[r-1][c] != strings.ToLower(v) {
		if (m[r][c-1] == strings.ToLower(v) && strings.ToLower(m[r-1][c-1]) != strings.ToLower(v)) && (m[r][c+1] == strings.ToLower(v) && strings.ToLower(m[r-1][c+1]) != strings.ToLower(v)) {
			fmt.Println("Connection")
			p -= 1
		} else if (m[r][c-1] == strings.ToLower(v) && strings.ToLower(m[r-1][c-1]) != strings.ToLower(v)) || (m[r][c+1] == strings.ToLower(v) && strings.ToLower(m[r-1][c+1]) != strings.ToLower(v)) {
		} else {
			p += 1
		}
	}
	if m[r][c+1] != v && m[r][c+1] != strings.ToLower(v) {
		if (m[r-1][c] == strings.ToLower(v) && strings.ToLower(m[r-1][c+1]) != strings.ToLower(v)) && (m[r+1][c] == strings.ToLower(v) && strings.ToLower(m[r+1][c+1]) != strings.ToLower(v)) {
			fmt.Println("Connection")
			p -= 1
		} else if (m[r-1][c] == strings.ToLower(v) && strings.ToLower(m[r-1][c+1]) != strings.ToLower(v)) || (m[r+1][c] == strings.ToLower(v) && strings.ToLower(m[r+1][c+1]) != strings.ToLower(v)) {
		} else {
			p += 1
		}
	}
	if m[r+1][c] != v && m[r+1][c] != strings.ToLower(v) {
		if (m[r][c-1] == strings.ToLower(v) && strings.ToLower(m[r+1][c-1]) != strings.ToLower(v)) && (m[r][c+1] == strings.ToLower(v) && strings.ToLower(m[r+1][c+1]) != strings.ToLower(v)) {
			fmt.Println("Connection")
			p -= 1
		} else if (m[r][c-1] == strings.ToLower(v) && strings.ToLower(m[r+1][c-1]) != strings.ToLower(v)) || (m[r][c+1] == strings.ToLower(v) && strings.ToLower(m[r+1][c+1]) != strings.ToLower(v)) {
		} else {
			p += 1
		}
	}
	if m[r][c-1] != v && m[r][c-1] != strings.ToLower(v) {
		if (m[r-1][c] == strings.ToLower(v) && strings.ToLower(m[r-1][c-1]) != strings.ToLower(v)) && (m[r+1][c] == strings.ToLower(v) && strings.ToLower(m[r+1][c-1]) != strings.ToLower(v)) {
			fmt.Println("Connection")
			p -= 1
		} else if (m[r-1][c] == strings.ToLower(v) && strings.ToLower(m[r-1][c-1]) != strings.ToLower(v)) || (m[r+1][c] == strings.ToLower(v) && strings.ToLower(m[r+1][c-1]) != strings.ToLower(v)) {
		} else {
			p += 1
		}
	}

	if m[r-1][c] == v {
		np, na := searchTwo(r-1, c, v, m)
		p += np
		a += na
	}
	if m[r][c+1] == v {
		np, na := searchTwo(r, c+1, v, m)
		p += np
		a += na
	}
	if m[r+1][c] == v {
		np, na := searchTwo(r+1, c, v, m)
		p += np
		a += na
	}
	if m[r][c-1] == v {
		np, na := searchTwo(r, c-1, v, m)
		p += np
		a += na
	}

	return p, a + 1
}

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	in := strings.Split(string(f), "\n")

	m := [][]string{}
	m = append(m, strings.Split(strings.Repeat(".", len(in[0])+2), ""))
	for _, l := range in {
		m = append(m, strings.Split("."+l+".", ""))
	}
	m = append(m, strings.Split(strings.Repeat(".", len(in[0])+2), ""))

	s := 0
	for r := 1; r < len(m)-1; r++ {
		for c := 1; c < len(m[0])-1; c++ {
			v := m[r][c]
			if strings.ToUpper(v) == v {
				p, a := searchTwo(r, c, v, m)
				s += p * a
			}
		}
	}
	fmt.Println(s)

	// fmt.Println(searchTwo(1, 1, "A", m))
}
