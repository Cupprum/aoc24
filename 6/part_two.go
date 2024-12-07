package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var count int = 0

func findLoop(mp []string, r int, c int, h [][]int, br int, bc int) bool {
	if r == 0 || r == len(mp)-1 || c == 0 || c == len(mp[0])-1 {
		return false
	}

	if len(h) > 4 && slices.Compare(h[0], h[4]) == 0 {
		count++
		return true
	}

	p := string(mp[r][c])

	if p == ">" {
		if mp[r][c+1] == '#' || (r == br && c+1 == bc) {
			mp = turn(mp, r, c)
			return findLoop(mp, r, c, slices.Concat([][]int{{r, c + 1}}, h), br, bc)
		} else {
			mp[r] = mp[r][:c] + "." + ">" + mp[r][c+2:]
			return findLoop(mp, r, c+1, h, br, bc)
		}
	} else if p == "V" {
		if mp[r+1][c] == '#' || (r+1 == br && c == bc) {
			mp = turn(mp, r, c)
			return findLoop(mp, r, c, slices.Concat([][]int{{r + 1, c}}, h), br, bc)
		} else {
			mp[r] = mp[r][:c] + "." + mp[r][c+1:]
			mp[r+1] = mp[r+1][:c] + "V" + mp[r+1][c+1:]
			return findLoop(mp, r+1, c, h, br, bc)
		}
	} else if p == "<" {
		if mp[r][c-1] == '#' || (r == br && c-1 == bc) {
			mp = turn(mp, r, c)
			return findLoop(mp, r, c, slices.Concat([][]int{{r, c - 1}}, h), br, bc)
		} else {
			mp[r] = mp[r][:c-1] + "<" + "." + mp[r][c+1:]
			return findLoop(mp, r, c-1, h, br, bc)
		}
	} else if p == "^" {
		if mp[r-1][c] == '#' || (r-1 == br && c == bc) {
			mp = turn(mp, r, c)
			return findLoop(mp, r, c, slices.Concat([][]int{{r - 1, c}}, h), br, bc)
		} else {
			mp[r] = mp[r][:c] + "." + mp[r][c+1:]
			mp[r-1] = mp[r-1][:c] + "^" + mp[r-1][c+1:]
			return findLoop(mp, r-1, c, h, br, bc)
		}
	}
	panic("Shouldnt happen")
}

func stepTwo(mp []string, r int, c int, h [][]int) []string {
	if r == 0 || r == len(mp)-1 || c == 0 || c == len(mp[0])-1 {
		mp[r] = mp[r][:c] + "X" + mp[r][c+1:]
		return mp
	}

	p := string(mp[r][c])

	if p == ">" {
		if mp[r][c+1] == '#' {
			mp = turn(mp, r, c)
			stepTwo(mp, r, c, h)
		} else {
			if !slices.Contains(h, []int{r, c + 1}) {
				lmp := make([]string, len(mp))
				copy(lmp, mp)
				lmp = turn(lmp, r, c)
				findLoop(lmp, r, c, [][]int{}, r, c+1)
			}
			mp[r] = mp[r][:c] + "." + ">" + mp[r][c+2:]
			stepTwo(mp, r, c+1, h)
		}
	} else if p == "V" {
		if mp[r+1][c] == '#' {
			mp = turn(mp, r, c)
			stepTwo(mp, r, c, h)
		} else {
			lmp := make([]string, len(mp))
			copy(lmp, mp)
			lmp = turn(lmp, r, c)
			findLoop(lmp, r, c, [][]int{}, r+1, c)
			mp[r] = mp[r][:c] + "." + mp[r][c+1:]
			mp[r+1] = mp[r+1][:c] + "V" + mp[r+1][c+1:]
			stepTwo(mp, r+1, c, h)
		}
	} else if p == "<" {

		if mp[r][c-1] == '#' {
			mp = turn(mp, r, c)
			stepTwo(mp, r, c, h)
		} else {
			lmp := make([]string, len(mp))
			copy(lmp, mp)
			lmp = turn(lmp, r, c)
			findLoop(lmp, r, c, [][]int{}, r, c-1)
			mp[r] = mp[r][:c-1] + "<" + "." + mp[r][c+1:]
			stepTwo(mp, r, c-1, h)
		}
	} else if p == "^" {
		if mp[r-1][c] == '#' {
			mp = turn(mp, r, c)
			stepTwo(mp, r, c, h)
		} else {
			lmp := make([]string, len(mp))
			copy(lmp, mp)
			lmp = turn(lmp, r, c)
			findLoop(lmp, r, c, [][]int{}, r-1, c)
			mp[r] = mp[r][:c] + "." + mp[r][c+1:]
			mp[r-1] = mp[r-1][:c] + "^" + mp[r-1][c+1:]
			stepTwo(mp, r-1, c, h)
		}
	}

	return mp
}

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	mp := strings.Split(string(f), "\n")

	r, c := findGuard(mp)

	h := [][]int{{r, c}}
	mp = stepTwo(mp, r, c, h)

	fmt.Println(strings.Join(mp, "\n"))
	// fmt.Println(countX(mp))
	fmt.Println(count)
}
