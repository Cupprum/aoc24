package main

import (
	"fmt"
	"strconv"
	"strings"
)

type M struct {
	i int
	s int
}

func searchTwo(i int, s int, mm map[M]int) int {
	if i == 75 {
		return 1
	}

	if v, ok := mm[M{i, s}]; ok {
		return v
	}

	sum := 0
	if s == 0 {
		sum += searchTwo(i+1, 1, mm)
	} else if rs := strconv.Itoa(s); len(rs)%2 == 0 {
		n1, _ := strconv.Atoi(rs[:len(rs)/2])
		n2, _ := strconv.Atoi(rs[len(rs)/2:])
		sum += searchTwo(i+1, n1, mm)
		sum += searchTwo(i+1, n2, mm)
	} else {
		sum += searchTwo(i+1, s*2024, mm)
	}
	mm[M{i, s}] = sum
	return sum
}

func partTwo() {
	in := strings.Split("965842 9159 3372473 311 0 6 86213 48", " ")

	ss := []int{}
	for _, rs := range in {
		s, err := strconv.Atoi(rs)
		if err != nil {
			panic(err)
		}
		ss = append(ss, s)
	}

	ss = searchOne(44, ss)

	sum := 0

	mm := make(map[M]int)
	for _, s := range ss {
		sum += searchTwo(44, s, mm)
	}

	fmt.Println(sum)
}
