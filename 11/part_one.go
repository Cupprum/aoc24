package main

import (
	"fmt"
	"strconv"
	"strings"
)

func searchOne(l int, ss []int) []int {
	for i := 0; i < l; i++ {
		ns := []int{}
		for _, s := range ss {
			if s == 0 {
				ns = append(ns, 1)
			} else if rs := strconv.Itoa(s); len(rs)%2 == 0 {
				n1, err := strconv.Atoi(rs[:len(rs)/2])
				if err != nil {
					panic(err)
				}
				n2, err := strconv.Atoi(rs[len(rs)/2:])
				if err != nil {
					panic(err)
				}
				ns = append(ns, n1)
				ns = append(ns, n2)
			} else {
				ns = append(ns, s*2024)
			}
		}
		ss = ns
	}
	return ss
}

func partOne() {
	in := strings.Split("0 1 10 99 999", " ")

	ss := []int{}
	for _, rs := range in {
		s, err := strconv.Atoi(rs)
		if err != nil {
			panic(err)
		}
		ss = append(ss, s)
	}

	ss = searchOne(25, ss)
	fmt.Println(len(ss))
}
