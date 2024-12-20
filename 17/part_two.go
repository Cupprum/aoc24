package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func search(pgm []int, p int, rra int) bool {
	if p < 0 {
		fmt.Println(rra)
		return true
	}
	for d := 0; d < 8; d++ {
		ra := rra<<3 | d
		rb := 0
		rc := 0

		ic := 0
		var out int

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
				out = co % 8
				break
			} else if opc == 6 {
				rb = ra >> co
			} else if opc == 7 {
				rc = ra >> co
			}

			ic += 2
		}
		if out == pgm[p] && search(pgm, p-1, rra<<3|d) {
			return true
		}

	}
	return false
}

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	pts := strings.Split(string(f), "\n\n")

	pgms := strings.Split(strings.TrimPrefix(pts[1], "Program: "), ",")
	pgm := []int{}
	for _, p := range pgms {
		pp, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}
		pgm = append(pgm, pp)
	}

	search(pgm, len(pgm)-1, 0)
}
