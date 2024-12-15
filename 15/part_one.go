package main

import (
	"fmt"
	"os"
	"strings"
)

func pushUpO(m [][]string, r, c int) bool {
	if m[r-1][c] == "#" {
		return false
	} else if m[r-1][c] == "." {
		m[r-1][c] = m[r][c]
		m[r][c] = "."
		return true
	} else if m[r-1][c] == "O" {
		b := pushUpO(m, r-1, c)
		if b {
			m[r-1][c] = m[r][c]
			m[r][c] = "."
		}
		return b
	}
	panic("pushUp")
}

func pushRightO(m [][]string, r, c int) bool {
	if m[r][c+1] == "#" {
		return false
	} else if m[r][c+1] == "." {
		m[r][c+1] = m[r][c]
		m[r][c] = "."
		return true
	} else if m[r][c+1] == "O" {
		b := pushRightO(m, r, c+1)
		if b {
			m[r][c+1] = m[r][c]
			m[r][c] = "."
		}
		return b
	}
	panic("pushRight")
}

func pushDownO(m [][]string, r, c int) bool {
	if m[r+1][c] == "#" {
		return false
	} else if m[r+1][c] == "." {
		m[r+1][c] = m[r][c]
		m[r][c] = "."
		return true
	} else if m[r+1][c] == "O" {
		b := pushDownO(m, r+1, c)
		if b {
			m[r+1][c] = m[r][c]
			m[r][c] = "."
		}
		return b
	}
	panic("pushDown")
}

func pushLeftO(m [][]string, r, c int) bool {
	if m[r][c-1] == "#" {
		return false
	} else if m[r][c-1] == "." {
		m[r][c-1] = m[r][c]
		m[r][c] = "."
		return true
	} else if m[r][c-1] == "O" {
		b := pushLeftO(m, r, c-1)
		if b {
			m[r][c-1] = m[r][c]
			m[r][c] = "."
		}
		return b
	}
	panic("pushLeft")
}

func partOne() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	pts := strings.Split(string(f), "\n\n")

	m := [][]string{}
	for _, l := range strings.Split(pts[0], "\n") {
		rs := strings.Split(l, "")
		m = append(m, rs)
	}

	stps := strings.Split(strings.Replace(pts[1], "\n", "", -1), "")

	var rc, rr int
	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m[0]); c++ {
			if m[r][c] == "@" {
				rc, rr = c, r
				break
			}
		}
	}

	for _, s := range stps {
		switch s {
		case "^":
			if pushUpO(m, rr, rc) {
				rr--
			}
		case ">":
			if pushRightO(m, rr, rc) {
				rc++
			}
		case "v":
			if pushDownO(m, rr, rc) {
				rr++
			}
		case "<":
			if pushLeftO(m, rr, rc) {
				rc--
			}
		}
	}

	sum := 0

	for r, l := range m {
		for c, v := range l {
			if v == "O" {
				sum += 100*r + c
			}
		}
	}

	fmt.Println(sum)
}
