package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stepTwo(ns []int, d int, c int, p int) bool {
	if p >= len(ns) {
		return c == d
	}

	f := false
	if stepTwo(ns, d, c+ns[p], p+1) {
		f = true
	}
	if stepTwo(ns, d, c*ns[p], p+1) {
		f = true
	}
	rcmb := strconv.Itoa(c) + strconv.Itoa(ns[p])
	cmb, err := strconv.Atoi(rcmb)
	if err != nil {
		panic(err)
	}
	if stepTwo(ns, d, cmb, p+1) {
		f = true
	}

	return f
}

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	ls := strings.Split(string(f), "\n")

	c := 0
	for _, l := range ls {
		ps := strings.Split(l, ": ")
		d, err := strconv.Atoi(ps[0])
		if err != nil {
			panic(err)
		}
		rns := strings.Split(ps[1], " ")
		ns := make([]int, len(rns))
		for i, rn := range rns {
			n, err := strconv.Atoi(rn)
			if err != nil {
				panic(err)
			}
			ns[i] = n
		}

		if stepTwo(ns, d, ns[0], 1) {
			c += d
		}
	}
	fmt.Println(c)
}