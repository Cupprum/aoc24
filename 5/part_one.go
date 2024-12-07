package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkRule(rls [][]int, l int, r int) bool {
	for _, rl := range rls {
		// I need to compare the oposite, to prove that its not broken.
		if rl[0] == r && rl[1] == l {
			return false
		}
	}

	return true
}

func verifyIn(rls [][]int, in []int) bool {
	for i, l := range in {
		if i == len(in) {
			break
		}
		// fmt.Println("l: ", l)
		for _, r := range in[i+1:] {
			// fmt.Println("r: ", r)

			if !checkRule(rls, l, r) {
				return false
			}
		}
	}

	return true
}

func partOne() {
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
		if b {
			s += in[len(in)/2]
		}
	}

	fmt.Println("Sum: ", s)
}
