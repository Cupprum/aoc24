package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type P struct {
	x int
	y int
}

func searchOne(p P, a P, b P) int {
	c := P{0, 0}

	ca := 0
	cb := 0

	for {
		if c.x+a.x <= p.x && c.y+a.y <= p.y && ca < 100 {
			c.x += a.x
			c.y += a.y
			ca++
		} else {
			for c.x+b.x <= p.x && c.y+b.y <= p.y && cb < 100 {
				c.x += b.x
				c.y += b.y
				cb++
			}
			break
		}
	}

	tkns := []int{}

	for ; ca >= 0; ca-- {
		if c.x == p.x && c.y == p.y {
			tkns = append(tkns, ca*3+cb)
		}

		c.x -= a.x
		c.y -= a.y

		for c.x+b.x <= p.x && c.y+b.y <= p.y && cb < 100 {
			c.x += b.x
			c.y += b.y
			cb++
		}
	}

	sort.Ints(tkns)

	if len(tkns) == 0 {
		return -1
	}
	return tkns[0]
}

func partOne() {
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
		pp := P{px, py}

		if n := searchOne(pp, ap, bp); n != -1 {
			c += n
		}
	}

	fmt.Println(c)
}
