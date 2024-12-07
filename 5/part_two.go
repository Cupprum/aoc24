package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func order(rls [][]int, in []int) []int {
	slices.SortFunc(in, func(a, b int) int {
		for _, rl := range rls {
			if rl[0] == a && rl[1] == b {
				return -1
			}
		}
		return 1
	})

	fmt.Println("Order: ", in)
	return in
}

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	scs := strings.Split(string(f), "\n\n")

	var rls [][]int
	for _, rl := range strings.Split(scs[0], "\n") {
		lar := strings.Split(string(rl), "|")
		l, err := strconv.Atoi(lar[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(lar[1])
		if err != nil {
			panic(err)
		}

		rls = append(rls, []int{l, r})
	}

	ins := [][]int{}
	for _, in := range strings.Split(scs[1], "\n") {
		s := strings.Split(in, ",")

		inn := make([]int, len(s))
		for i, v := range s {
			inn[i], err = strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
		}

		ins = append(ins, inn)
	}

	s := 0
	for _, in := range ins {
		b := verifyIn(rls, in)
		if !b {
			o := order(rls, in)
			s += o[len(o)/2]
		}
	}

	fmt.Println("Sum: ", s)
}
