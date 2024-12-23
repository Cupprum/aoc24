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
	sscts := strings.Split(string(f), "\n")
	scts := make([]int, len(sscts))

	for i, sct := range sscts {
		scts[i], err = strconv.Atoi(sct)
		if err != nil {
			panic(err)
		}
	}

	hst := map[string]int{}
	for _, sct := range scts {
		dfs := []int{}

		lhst := map[string]bool{}

		lst := sct % 10
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

			dfs = append(dfs, (sct%10)-lst)
			lst = sct % 10

			if len(dfs) >= 4 {
				ll := dfs[len(dfs)-4:]
				n := ""
				for _, l := range ll {
					n += strconv.Itoa(l) + ","
				}
				if _, ok := lhst[n]; !ok {
					lhst[n] = true
					hst[n] += lst
				}
			}
		}
	}

	mv := 0
	for _, v := range hst {
		if v > mv {
			mv = v
		}
	}
	fmt.Println(mv)
}
