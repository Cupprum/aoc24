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
	sscts := strings.Split(string(f), "\n")
	scts := make([]int, len(sscts))

	for i, sct := range sscts {
		scts[i], err = strconv.Atoi(sct)
		if err != nil {
			panic(err)
		}
	}

	s := 0
	for _, sct := range scts {
		for i := 0; i < 2000; i++ {
			t := sct * 64
			sct = sct ^ t
			sct = sct % 16777216

			t = sct / 32
			sct = sct ^ t
			sct = sct % 16777216

			t = sct * 2048
			sct = sct ^ t
			sct = sct % 16777216
		}
		s += sct
	}
	fmt.Println(s)
}
