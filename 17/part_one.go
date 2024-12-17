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
	pts := strings.Split(string(f), "\n\n")

	rgs := strings.Split(pts[0], "\n")
	ra, err := strconv.Atoi(strings.TrimPrefix(rgs[0], "Register A: "))
	if err != nil {
		panic(err)
	}
	rb, err := strconv.Atoi(strings.TrimPrefix(rgs[1], "Register B: "))
	if err != nil {
		panic(err)
	}
	rc, err := strconv.Atoi(strings.TrimPrefix(rgs[2], "Register C: "))
	if err != nil {
		panic(err)
	}

	pgms := strings.Split(strings.TrimPrefix(pts[1], "Program: "), ",")
	pgm := []int{}
	for _, p := range pgms {
		pp, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}
		pgm = append(pgm, pp)
	}

	outs := []int{}

	ic := 0

	for ic < len(pgm) {
		opc := pgm[ic]
		opr := pgm[ic+1]

		co := 0
		if opr >= 0 && opr <= 3 {
			co = opr
		} else if opr == 4 {
			co = ra
		} else if opr == 5 {
			co = rb
		} else if opr == 6 {
			co = rc
		}

		if opc == 0 {
			ra = ra >> co
		} else if opc == 1 {
			rb = rb ^ opr
		} else if opc == 2 {
			rb = co % 8
		} else if opc == 3 && ra != 0 {
			ic = opr
			continue
		} else if opc == 4 {
			rb = rb ^ rc
		} else if opc == 5 {
			outs = append(outs, co%8)
		} else if opc == 6 {
			rb = ra >> co
		} else if opc == 7 {
			rc = ra >> co
		}

		ic += 2
	}

	s := ""
	for _, o := range outs {
		s += strconv.Itoa(o) + ","
	}
	s = s[:len(s)-1]
	fmt.Println(s)
}
