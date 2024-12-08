package main

import (
	"fmt"
	"os"
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

func turn(mp []string, r int, c int) []string {
	p := string(mp[r][c])

	if p == ">" {
		mp[r] = strings.Replace(mp[r], ">", "V", 1)
	} else if p == "V" {
		mp[r] = strings.Replace(mp[r], "V", "<", 1)
	} else if p == "<" {
		mp[r] = strings.Replace(mp[r], "<", "^", 1)
	} else if p == "^" {
		mp[r] = strings.Replace(mp[r], "^", ">", 1)
	}
	return mp
}

func countX(mp []string) int {
	s := 1

	for _, v := range mp {
		s += strings.Count(v, "X")
	}
	return s
}

func partOne() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	mp := strings.Split(string(f), "\n")

	r, c := findGuard(mp)

	for r > 0 && r < len(mp)-1 && c > 0 && c < len(mp[0])-1 {
		if mp[r][c] == '>' {
			if mp[r][c+1] == '#' {
				turn(mp, r, c)
			} else {
				mp[r] = mp[r][:c] + "X" + ">" + mp[r][c+2:]
				c = c + 1
			}
		} else if mp[r][c] == 'V' {
			if mp[r+1][c] == '#' {
				turn(mp, r, c)
			} else {
				mp[r] = mp[r][:c] + "X" + mp[r][c+1:]
				mp[r+1] = mp[r+1][:c] + "V" + mp[r+1][c+1:]
				r = r + 1
			}
		} else if mp[r][c] == '<' {
			if mp[r][c-1] == '#' {
				turn(mp, r, c)
			} else {
				mp[r] = mp[r][:c-1] + "<" + "X" + mp[r][c+1:]
				c = c - 1
			}
		} else if mp[r][c] == '^' {
			if mp[r-1][c] == '#' {
				turn(mp, r, c)
			} else {
				mp[r-1] = mp[r-1][:c] + "^" + mp[r-1][c+1:]
				mp[r] = mp[r][:c] + "X" + mp[r][c+1:]
				r = r - 1
			}
		} else {
			fmt.Println(string(mp[r][c]))
			panic("Invalid position")
		}
	}

	fmt.Println(countX(mp))
}
