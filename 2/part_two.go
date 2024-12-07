package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func checkLevel(l []int, f bool) bool {
	for i := 0; i < len(l)-1; i++ {
		if l[0] < l[1] { // Ascending
			if l[i+1]-l[i] >= 1 && l[i+1]-l[i] <= 3 {
			} else {
				return false
			}
		} else {
			if l[i]-l[i+1] >= 1 && l[i]-l[i+1] <= 3 {
			} else {
				return false
			}
		}
	}
	return f
}

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	rls := strings.Split(string(f), "\n")

	ls := [][]int{}

	for _, rrl := range rls {
		rl := strings.Split(rrl, " ")
		l := []int{}
		for _, rr := range rl {
			r, err := strconv.Atoi(rr)
			if err != nil {
				panic(err)
			}
			l = append(l, r)
		}
		ls = append(ls, l)
	}

	s := 0

	for _, l := range ls {

		f := true

		f = checkLevel(l, f)
		if f {
			s++
		} else {
			for i := 0; i < len(l); i++ {
				f = true
				f = checkLevel(slices.Concat(l[:i], l[i+1:]), f)
				if f {
					s++
					break
				}
			}
		}
	}

	fmt.Println(s)
}
