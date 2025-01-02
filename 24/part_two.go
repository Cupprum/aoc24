package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Te struct {
	op  string
	l   string
	r   string
	out string
}

func partTwo() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	pts := strings.Split(string(f), "\n\n")

	var gts []Te
	for _, line := range strings.Split(pts[1], "\n") {
		if line == "" {
			continue
		}
		t := strings.Split(line, " -> ")
		eq := strings.Split(t[0], " ")
		gts = append(gts, Te{
			op:  eq[1],
			l:   eq[0],
			r:   eq[2],
			out: t[1],
		})
	}

	ww := make(map[string]bool)

	// Find highest z-wire for reference
	mz := "z00"
	for _, gate := range gts {
		if strings.HasPrefix(gate.out, "z") {
			if gate.out > mz {
				mz = gate.out
			}
		}
	}

	// First pass: check direct patterns
	for _, gt := range gts {
		// Pattern 1: z-wires that use non-XOR operations
		if strings.HasPrefix(gt.out, "z") && gt.op != "XOR" && gt.out != mz {
			ww[gt.out] = true
		}

		// Pattern 2: XOR gates with non-x/y/z inputs/outputs
		if gt.op == "XOR" &&
			!strings.HasPrefix(gt.out, "x") && !strings.HasPrefix(gt.out, "y") && !strings.HasPrefix(gt.out, "z") &&
			!strings.HasPrefix(gt.l, "x") && !strings.HasPrefix(gt.l, "y") && !strings.HasPrefix(gt.l, "z") &&
			!strings.HasPrefix(gt.r, "x") && !strings.HasPrefix(gt.r, "y") && !strings.HasPrefix(gt.r, "z") {
			ww[gt.out] = true
		}
	}

	// Second pass: check relationships between gates
	for _, gt := range gts {
		// Pattern 3: AND gates not using x00
		if gt.op == "AND" && gt.l != "x00" && gt.r != "x00" {
			for _, otherGate := range gts {
				if (otherGate.l == gt.out || otherGate.r == gt.out) && otherGate.op != "OR" {
					ww[gt.out] = true
				}
			}
		}

		// Pattern 4: XOR gates connected to OR gates
		if gt.op == "XOR" {
			for _, otherGate := range gts {
				if (otherGate.l == gt.out || otherGate.r == gt.out) && otherGate.op == "OR" {
					ww[gt.out] = true
				}
			}
		}
	}

	aww := make([]string, 0, len(ww))
	for w := range ww {
		aww = append(aww, w)
	}
	sort.Strings(aww)

	fmt.Println(strings.Join(aww, ","))
}
