package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func partOne() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	in := string(f)

	r := []int{}
	c := 0
	for i := 0; i < len(in); i++ {
		n, err := strconv.Atoi(string(in[i]))
		if err != nil {
			panic(err)
		}
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				r = append(r, c)
			}
			c++
		} else if i%2 == 1 {
			for j := 0; j < n; j++ {
				r = append(r, -1)
			}
		} else {
			panic("this should never happen")
		}
	}

	i := len(r) - 1
	for slices.Contains(r, -1) {
		if r[i] != -1 {
			l := slices.Index(r, -1)
			r[l] = r[i]
		}
		r = r[:i]
		i--
	}

	s := 0
	for i, n := range r {
		s += i * n
	}

	fmt.Println(s)
}
