package main

import (
	"2015/utils"
	"fmt"
	"strconv"
	"strings"
)

func calculate_ribbon_needed(line string) int {
	dimensions := strings.Split(line, "x")
	l, _ := strconv.Atoi(dimensions[0])
	w, _ := strconv.Atoi(dimensions[1])
	h, _ := strconv.Atoi(dimensions[2])

	smallest := min(min(l+w, w+h), h+l)

	return 2*smallest + l*w*h
}

func solve_02_2() {
	input, err := utils.ReadInput("02/02.0.input.txt")
	if err != nil {
		panic(err)
	}

	total := 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		total += calculate_ribbon_needed(line)
	}

	fmt.Println(total)
}