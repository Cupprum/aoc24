package main

import (
	"fmt"
	"os"
	"strings"
)

func partOne() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	pts := strings.Split(string(f), "\n\n")

	kys := [][]int{}
	lks := [][]int{}

	for _, pt := range pts {
		lns := strings.Split(pt, "\n")
		if lns[0] == "#####" {
			// Key
			ky := []int{}
			for c := 0; c < len(lns[0]); c++ {
				cnt := 0
				for r := 1; r < len(lns); r++ {
					if lns[r][c] == '#' {
						cnt += 1
					} else {
						break
					}
				}
				ky = append(ky, cnt)
			}
			kys = append(kys, ky)
		} else if lns[len(lns)-1] == "#####" {
			// Lock
			lk := []int{}
			for c := 0; c < len(lns[0]); c++ {
				cnt := 0
				for r := len(lns) - 2; r >= 0; r-- {
					if lns[r][c] == '#' {
						cnt += 1
					} else {
						break
					}
				}
				lk = append(lk, cnt)
			}
			lks = append(lks, lk)
		}
	}

	s := 0

	for i := 0; i < len(kys); i++ {
		for j := 0; j < len(lks); j++ {
			fl := true
			for c := 0; c < len(kys[0]); c++ {
				if kys[i][c]+lks[j][c] > 5 {
					fl = false
					break
				}
			}
			if fl {
				s += 1
			}
		}
	}

	fmt.Println(s)
}
