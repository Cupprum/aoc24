package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func partTwo() {
	i := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	re := regexp.MustCompile(`(mul\(\d+,\d+\)|don't\(\)|do\(\))`)
	ms := re.FindAllString(i, -1)

	s := 0

	do := true

	for _, m := range ms {
		if strings.Contains(m, "don't") {
			do = false
		} else if strings.Contains(m, "do") {
			do = true
		} else if strings.Contains(m, "mul") && do {
			m = strings.TrimPrefix(m, "mul(")
			m = strings.TrimSuffix(m, ")")
			vs := strings.Split(m, ",")
			l, err := strconv.Atoi(vs[0])
			if err != nil {
				panic(err)
			}
			r, err := strconv.Atoi(vs[1])
			if err != nil {
				panic(err)
			}
			s += l * r
		}
	}

	fmt.Println(s)
}
