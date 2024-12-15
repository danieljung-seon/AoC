package main

import (
	"2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func go_10_2(start Point, parsed [][]int, sum *int) {
	next_value := start.v + 1

	// up
	up := Point{start.x, start.y - 1, next_value}
	if get_value(parsed, up) == next_value {
		if next_value == 9 {
			*sum += 1
		} else {
			go_10_2(up, parsed, sum)
		}
	}

	// down
	down := Point{start.x, start.y + 1, next_value}
	if get_value(parsed, down) == next_value {
		if next_value == 9 {
			*sum += 1
		} else {
			go_10_2(down, parsed, sum)
		}
	}

	// left
	left := Point{start.x - 1, start.y, next_value}
	if get_value(parsed, left) == next_value {
		if next_value == 9 {
			*sum += 1
		} else {
			go_10_2(left, parsed, sum)
		}
	}

	// right
	right := Point{start.x + 1, start.y, next_value}
	if get_value(parsed, right) == next_value {
		if next_value == 9 {
			*sum += 1
		} else {
			go_10_2(right, parsed, sum)
		}
	}
}

func trail_10_2(start Point, parsed [][]int) int {
	sum := 0
	go_10_2(start, parsed, &sum)
	return sum
}

func solve_10_2() {
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
		add := trail_10_2(zero, parsed)
		sum += add
	}

	fmt.Println(sum)
}
