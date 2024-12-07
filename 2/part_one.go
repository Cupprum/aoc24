package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partOne() {
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
		a := l[0] < l[1]
		f := true

		for i := 0; i < len(l)-1; i++ {
			if a {
				if l[i+1]-l[i] >= 1 && l[i+1]-l[i] <= 3 {
				} else {
					f = false
					break
				}
			} else {
				if l[i]-l[i+1] >= 1 && l[i]-l[i+1] <= 3 {
				} else {
					f = false
					break
				}
			}
		}
		if f {
			s++
		}
	}

	fmt.Println(s)

}
