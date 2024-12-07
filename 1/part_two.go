package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(f), "\n")

	// // Each line consists of five characters and we dont want the last line.
	l := len(f)/5 - 1

	l1 := make([]int, l)
	l2 := make([]int, l)

	for i, l := range lines {
		ns := strings.Split(l, "   ")
		if len(ns) != 2 {
			panic("this should not happen")
		}
		l1[i], err = strconv.Atoi(ns[0])
		if err != nil {
			panic(err)
		}
		l2[i], err = strconv.Atoi(ns[1])
		if err != nil {
			panic(err)
		}
	}

	// sort.Ints(l1)
	// sort.Ints(l2)

	m := make(map[int]int)

	for _, v := range l2 {
		m[v]++
	}

	s := 0
	for _, v := range l1 {
		s += v * m[v]
	}
	fmt.Println(s)
}
