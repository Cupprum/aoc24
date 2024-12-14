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
	pts := strings.Split(string(f), "\n\n")

	c := 0
	for _, pt := range pts {
		ls := strings.Split(pt, "\n")
		a := strings.TrimPrefix(ls[0], "Button A: ")
		b := strings.TrimPrefix(ls[1], "Button B: ")
		p := strings.TrimPrefix(ls[2], "Prize: ")

		axy := strings.Split(a, ", ")
		ax, err := strconv.Atoi(strings.TrimPrefix(axy[0], "X+"))
		if err != nil {
			panic(err)
		}
		ay, err := strconv.Atoi(strings.TrimPrefix(axy[1], "Y+"))
		if err != nil {
			panic(err)
		}
		ap := P{ax, ay}
		bxy := strings.Split(b, ", ")
		bx, err := strconv.Atoi(strings.TrimPrefix(bxy[0], "X+"))
		if err != nil {
			panic(err)
		}
		by, err := strconv.Atoi(strings.TrimPrefix(bxy[1], "Y+"))
		if err != nil {
			panic(err)
		}
		bp := P{bx, by}
		pxy := strings.Split(p, ", ")
		px, err := strconv.Atoi(strings.TrimPrefix(pxy[0], "X="))
		if err != nil {
			panic(err)
		}
		py, err := strconv.Atoi(strings.TrimPrefix(pxy[1], "Y="))
		if err != nil {
			panic(err)
		}
		pp := P{px + 10000000000000, py + 10000000000000}

		bb := float64(ap.x*pp.y-ap.y*pp.x) / float64(bp.y*ap.x-bp.x*ap.y)
		aa := float64(pp.x-int(bb)*bp.x) / float64(ap.x)

		if aa == float64(int(aa)) && bb == float64(int(bb)) {
			c += int(aa)*3 + int(bb)
		}
	}

	fmt.Println(c)
}
