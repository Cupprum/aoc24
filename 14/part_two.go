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
	ls := strings.Split(string(f), "\n")

	rs := make([]R, len(ls))
	for i, l := range ls {
		vs := strings.Split(l, " ")
		pxy := strings.Split(strings.TrimPrefix(vs[0], "p="), ",")
		vxy := strings.Split(strings.TrimPrefix(vs[1], "v="), ",")
		rs[i] = R{p: P{x: atoi(pxy[0]), y: atoi(pxy[1])}, v: P{x: atoi(vxy[0]), y: atoi(vxy[1])}}
	}

	mr := 103
	mc := 101

	var grid [][]int

	// Hopefully 100000 is enough
	for s := 0; s < 100000; s++ {

		for i, r := range rs {
			cx := (r.p.x + r.v.x) % mc
			if cx < 0 {
				rs[i].p.x = mc + cx
			} else {
				rs[i].p.x = cx
			}
			cy := (r.p.y + r.v.y) % mr
			if cy < 0 {
				rs[i].p.y = mr + cy
			} else {
				rs[i].p.y = cy
			}
			continue
		}

		grid = make([][]int, mr)
		for i := range grid {
			grid[i] = make([]int, mc)
		}

		q1 := 0
		q2 := 0
		q3 := 0
		q4 := 0
		for _, r := range rs {
			grid[r.p.y][r.p.x] += 1
			if r.p.y < mr/2 && r.p.x < mc/2 {
				q1 += 1
			}
			if r.p.y < mr/2 && r.p.x > mc/2 {
				q2 += 1
			}
			if r.p.y > mr/2 && r.p.x < mc/2 {
				q3 += 1
			}
			if r.p.y > mr/2 && r.p.x > mc/2 {
				q4 += 1
			}
		}

		for _, i := range grid {
			l := ""
			for _, j := range i {
				if j > 0 {
					l += strconv.Itoa(j)
				} else {
					l += "."
				}
			}
			if strings.Contains(l, "111111111") {
				fmt.Println(s + 1)
				return
			}
		}
	}
}
