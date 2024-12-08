package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type H struct {
	r int
	c int
	o byte
}

func subSearch(mp []string, r int, c int, dr int, dc int) bool {
	h := []H{}

	f := false
	for r > 0 && r < len(mp)-1 && c > 0 && c < len(mp[0])-1 {
		if len(h) > 3 && slices.Contains(h, H{r, c, mp[r][c]}) {
			fmt.Println(h)
			f = true
			break
		}
		if mp[r][c] == '>' {
			if mp[r][c+1] == '#' || (r == dr && c+1 == dc) {
				h = append(h, H{r, c, '>'})
				turn(mp, r, c)
			} else {
				mp[r] = mp[r][:c] + "X" + ">" + mp[r][c+2:]
				c = c + 1
			}
		} else if mp[r][c] == 'V' {
			if mp[r+1][c] == '#' || (r+1 == dr && c == dc) {
				h = append(h, H{r, c, 'V'})
				turn(mp, r, c)
			} else {
				mp[r] = mp[r][:c] + "X" + mp[r][c+1:]
				mp[r+1] = mp[r+1][:c] + "V" + mp[r+1][c+1:]
				r = r + 1
			}
		} else if mp[r][c] == '<' {
			if mp[r][c-1] == '#' || (r == dr && c-1 == dc) {
				h = append(h, H{r, c, '<'})
				turn(mp, r, c)
			} else {
				mp[r] = mp[r][:c-1] + "<" + "X" + mp[r][c+1:]
				c = c - 1
			}
		} else if mp[r][c] == '^' {
			if mp[r-1][c] == '#' || (r-1 == dr && c == dc) {
				h = append(h, H{r, c, '^'})
				turn(mp, r, c)
			} else {
				mp[r-1] = mp[r-1][:c] + "^" + mp[r-1][c+1:]
				mp[r] = mp[r][:c] + "X" + mp[r][c+1:]
				r = r - 1
			}
		} else {
			panic("Invalid position")
		}
	}

	return f
}

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	mp := strings.Split(string(f), "\n")

	r, c := findGuard(mp)

	count := 0

	for r > 0 && r < len(mp)-1 && c > 0 && c < len(mp[0])-1 {
		if mp[r][c] == '>' {
			if mp[r][c+1] == '#' {
				turn(mp, r, c)
			} else {
				nmp := make([]string, len(mp))
				copy(nmp, mp)
				turn(nmp, r, c)
				if subSearch(nmp, r, c, r, c+1) {
					count++
				}
				mp[r] = mp[r][:c] + "X" + ">" + mp[r][c+2:]
				c = c + 1
			}
		} else if mp[r][c] == 'V' {
			if mp[r+1][c] == '#' {
				turn(mp, r, c)
			} else {
				nmp := make([]string, len(mp))
				copy(nmp, mp)
				turn(nmp, r, c)
				if subSearch(nmp, r, c, r+1, c) {
					count++
				}
				mp[r] = mp[r][:c] + "X" + mp[r][c+1:]
				mp[r+1] = mp[r+1][:c] + "V" + mp[r+1][c+1:]
				r = r + 1
			}
		} else if mp[r][c] == '<' {
			if mp[r][c-1] == '#' {
				turn(mp, r, c)
			} else {
				nmp := make([]string, len(mp))
				copy(nmp, mp)
				turn(nmp, r, c)
				if subSearch(nmp, r, c, r, c-1) {
					count++
				}
				mp[r] = mp[r][:c-1] + "<" + "X" + mp[r][c+1:]
				c = c - 1
			}
		} else if mp[r][c] == '^' {
			if mp[r-1][c] == '#' {
				turn(mp, r, c)
			} else {
				nmp := make([]string, len(mp))
				copy(nmp, mp)
				turn(nmp, r, c)
				if subSearch(nmp, r, c, r-1, c) {
					count++
				}
				mp[r-1] = mp[r-1][:c] + "^" + mp[r-1][c+1:]
				mp[r] = mp[r][:c] + "X" + mp[r][c+1:]
				r = r - 1
			}
		} else {
			panic("Invalid position")
		}
	}

	fmt.Println(strings.Join(mp, "\n"))
	fmt.Println(count)
}
