package main

import (
	"fmt"
	"os"
	"strings"
)

func verifyOne(p string, cmbs []string) bool {
	if p == "" {
		return true
	}

	for _, c := range cmbs {
		if strings.HasPrefix(p, c) {
			tp := strings.TrimPrefix(p, c)
			if verifyOne(tp, cmbs) {
				return true
			}
		}
	}

	return false
}

func partOne() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	pts := strings.Split(string(f), "\n\n")

	cmbs := strings.Split(pts[0], ", ")
	ptrns := strings.Split(pts[1], "\n")

	s := 0

	for _, p := range ptrns {
		if verifyOne(p, cmbs) {
			s++
		}
	}

	fmt.Println(s)
}
