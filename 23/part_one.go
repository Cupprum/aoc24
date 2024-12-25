package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func partOne() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	prs := strings.Split(string(f), "\n")

	m := map[string][]string{}
	for _, pr := range prs {
		cs := strings.Split(pr, "-")
		c1 := cs[0]
		c2 := cs[1]

		if _, ok := m[c1]; !ok {
			m[c1] = []string{c2}
		} else {
			m[c1] = append(m[c1], c2)
		}

		if _, ok := m[c2]; !ok {
			m[c2] = []string{c1}
		} else {
			m[c2] = append(m[c2], c1)
		}
	}

	trps := []string{}
	for _, pr := range prs {
		cs := strings.Split(pr, "-")
		c1 := cs[0]
		c2 := cs[1]

		for _, i1 := range m[c1] {
			for _, i2 := range m[c2] {
				if i1 == i2 {
					trp := []string{c1, c2, i1}
					sort.Strings(trp)

					pv := strings.Join(trp, "-")

					if (c1[0] == 't' || c2[0] == 't' || i1[0] == 't') && !slices.Contains(trps, pv) {
						trps = append(trps, pv)
					}
				}
			}
		}
	}

	fmt.Println(len(trps))
}
