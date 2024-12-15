package main

import (
	"fmt"
	"os"
	"strings"
)

func pushUpT(m [][]string, r, c int, p bool) bool {
	if m[r-1][c] == "#" {
		return false
	} else if m[r-1][c] == "." {
		if p {
			m[r-1][c] = m[r][c]
			m[r][c] = "."
		}
		return true
	} else if m[r-1][c] == "[" {
		if pushUpT(m, r-1, c, false) && pushUpT(m, r-1, c+1, false) {
			if p {
				pushUpT(m, r-1, c, true)
				pushUpT(m, r-1, c+1, true)
				m[r-1][c] = m[r][c]
				m[r][c] = "."
			}
			return true
		}
		return false
	} else if m[r-1][c] == "]" {
		if pushUpT(m, r-1, c-1, false) && pushUpT(m, r-1, c, false) {
			if p {
				pushUpT(m, r-1, c-1, true)
				pushUpT(m, r-1, c, true)
				m[r-1][c] = m[r][c]
				m[r][c] = "."
			}
			return true
		}
		return false
	}
	panic("pushUp")
}

func pushDownT(m [][]string, r, c int, p bool) bool {
	if m[r+1][c] == "#" {
		return false
	} else if m[r+1][c] == "." {
		if p {
			m[r+1][c] = m[r][c]
			m[r][c] = "."
		}
		return true
	} else if m[r+1][c] == "[" {
		if pushDownT(m, r+1, c, false) && pushDownT(m, r+1, c+1, false) {
			if p {
				pushDownT(m, r+1, c, true)
				pushDownT(m, r+1, c+1, true)
				m[r+1][c] = m[r][c]
				m[r][c] = "."
			}
			return true
		}
		return false
	} else if m[r+1][c] == "]" {
		if pushDownT(m, r+1, c-1, false) && pushDownT(m, r+1, c, false) {
			if p {
				pushDownT(m, r+1, c-1, true)
				pushDownT(m, r+1, c, true)
				m[r+1][c] = m[r][c]
				m[r][c] = "."
			}
			return true
		}
		return false
	}
	panic("pushDown")
}

func pushRightT(m [][]string, r, c int) bool {
	if m[r][c+1] == "#" {
		return false
	} else if m[r][c+1] == "." {
		m[r][c+1] = m[r][c]
		m[r][c] = "."
		return true
	} else if m[r][c+1] == "[" {
		if pushRightT(m, r, c+2) {
			m[r][c+2] = m[r][c+1]
			m[r][c+1] = m[r][c]
			m[r][c] = "."
			return true
		}
		return false
	}
	panic("pushRight")
}

func pushLeftT(m [][]string, r, c int) bool {
	if m[r][c-1] == "#" {
		return false
	} else if m[r][c-1] == "." {
		m[r][c-1] = m[r][c]
		m[r][c] = "."
		return true
	} else if m[r][c-1] == "]" {
		if pushLeftT(m, r, c-2) {
			m[r][c-2] = m[r][c-1]
			m[r][c-1] = m[r][c]
			m[r][c] = "."
			return true
		}
		return false
	}
	panic("pushLeft")
}

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	pts := strings.Split(string(f), "\n\n")

	m := [][]string{}
	for _, l := range strings.Split(pts[0], "\n") {

		rs := []string{}
		rrs := strings.Split(l, "")
		for _, r := range rrs {
			switch r {
			case "#":
				rs = append(rs, "#", "#")
			case "O":
				rs = append(rs, "[", "]")
			case ".":
				rs = append(rs, ".", ".")
			case "@":
				rs = append(rs, "@", ".")
			}
		}
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
			if pushUpT(m, rr, rc, true) {
				rr--
			}
		case ">":
			if pushRightT(m, rr, rc) {
				rc++
			}
		case "v":
			if pushDownT(m, rr, rc, true) {
				rr++
			}
		case "<":
			if pushLeftT(m, rr, rc) {
				rc--
			}
		}
	}

	sum := 0

	for r, l := range m {
		for c, v := range l {
			if v == "[" {
				sum += 100*r + c
			}
		}
	}

	fmt.Println(sum)
}
