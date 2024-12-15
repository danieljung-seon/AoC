package main

import (
	"2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func go_10_1(start Point, parsed [][]int, ends *map[Point]bool) {
	next_value := start.v + 1

	// up
	up := Point{start.x, start.y - 1, next_value}
	if get_value(parsed, up) == next_value {
		if next_value == 9 {
			(*ends)[up] = true
		} else {
			go_10_1(up, parsed, ends)
		}
	}

	// down
	down := Point{start.x, start.y + 1, next_value}
	if get_value(parsed, down) == next_value {
		if next_value == 9 {
			(*ends)[down] = true
		} else {
			go_10_1(down, parsed, ends)
		}
	}

	// left
	left := Point{start.x - 1, start.y, next_value}
	if get_value(parsed, left) == next_value {
		if next_value == 9 {
			(*ends)[left] = true
		} else {
			go_10_1(left, parsed, ends)
		}
	}

	// right
	right := Point{start.x + 1, start.y, next_value}
	if get_value(parsed, right) == next_value {
		if next_value == 9 {
			(*ends)[right] = true
		} else {
			go_10_1(right, parsed, ends)
		}
	}
}

func trail_10_1(start Point, parsed [][]int) int {
	ends := make(map[Point]bool)
	go_10_1(start, parsed, &ends)
	return len(ends)
}

func solve_10_1() {
	input, err := utils.ReadInput("10/10.0.input.txt")
	if err != nil {
		panic(err)
	}

	parsed := make([][]int, 0)
	zeros := make([]Point, 0)
	lines := strings.Split(input, "\n")

	for y, rawline := range lines {
		line := strings.Split(rawline, "")

		for x, char := range line {
			if char == "0" {
				zeros = append(zeros, Point{x, y, 0})
			}

			if len(parsed) <= y {
				parsed = append(parsed, make([]int, 0))
			}

			if len(parsed[y]) <= x {
				parsed[y] = append(parsed[y], 0)
			}

			parsed[y][x], _ = strconv.Atoi(char)
		}
	}

	sum := 0
	for _, zero := range zeros {
		add := trail_10_1(zero, parsed)
		sum += add
	}

	fmt.Println(sum)
}
