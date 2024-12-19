package main

import (
	"fmt"
	"os"
	"strings"
)

func verifyTwo(p string, cmbs []string, hist map[string]int) int {
	if p == "" {
		return 1
	}

	s := 0
	for _, c := range cmbs {
		if strings.HasPrefix(p, c) {
			tp := strings.TrimPrefix(p, c)
			if v, ok := hist[tp]; ok {
				s += v
				continue
			}
			t := verifyTwo(tp, cmbs, hist)
			if t > 0 {
				hist[tp] = t
				s += t
			}
		}
	}

	return s
}

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	pts := strings.Split(string(f), "\n\n")

	cmbs := strings.Split(pts[0], ", ")
	ptrns := strings.Split(pts[1], "\n")

	s := 0

	hist := make(map[string]int)

	for _, p := range ptrns {
		s += verifyTwo(p, cmbs, hist)
	}

	fmt.Println(s)
}
