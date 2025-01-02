package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type T struct {
	op string
	l  string
	r  string
}

func calculate(vls map[string]int, m map[string]T, s string) int {
	t := m[s]

	if _, ok := vls[t.l]; !ok {
		vls[t.l] = calculate(vls, m, t.l)
	}
	if _, ok := vls[t.r]; !ok {
		vls[t.r] = calculate(vls, m, t.r)
	}

	switch t.op {
	case "AND":
		return vls[t.l] & vls[t.r]
	case "OR":
		return vls[t.l] | vls[t.r]
	case "XOR":
		return vls[t.l] ^ vls[t.r]
	}
	panic("invalid operator")
}

func partOne() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	pts := strings.Split(string(f), "\n\n")

	ins := strings.Split(pts[0], "\n")
	swrs := strings.Split(pts[1], "\n")

	vls := map[string]int{}

	for _, in := range ins {
		t := strings.Split(in, ": ")
		vls[t[0]], err = strconv.Atoi(t[1])
		if err != nil {
			panic(err)
		}
	}

	mz := 0
	m := map[string]T{}

	for _, swr := range swrs {
		t := strings.Split(swr, " -> ")

		if t[1][0] == 'z' {
			o, err := strconv.Atoi(t[1][1:])
			if err != nil {
				panic(err)
			}

			if o > mz {
				mz = o
			}
		}

		eq := strings.Split(t[0], " ")
		m[t[1]] = T{eq[1], eq[0], eq[2]}
	}

	r := ""
	for i := 0; i <= mz; i++ {
		s := strconv.Itoa(i)
		if len(s) == 1 {
			s = "0" + s
		}
		s = "z" + s

		r = strconv.Itoa(calculate(vls, m, s)) + r
	}

	num, err := strconv.ParseInt(r, 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
}
