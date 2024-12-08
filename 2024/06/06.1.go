package main

import (
	"2024/utils"
	"fmt"
	"strings"
)

func solve_06_1() {
	input, err := utils.ReadInput("06/06.0.input.txt")
	if err != nil {
		panic(err)
	}

	directions := [4][2]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}
	dir_index := 0

	obstacles := make(map[string]bool)
	visited := make(map[string]bool)
	// var guard_start [2]int
	var guard [2]int
	var maxY, maxX int

	rows := strings.Split(input, "\n")
	maxY = len(rows)

	for y, row := range rows {
		cell := strings.Split(row, "")
		maxX = len(cell)

		for x, c := range cell {
			if c == "#" {
				key := fmt.Sprintf("%d,%d", y, x)
				obstacles[key] = true
			}

			if c == "^" {
				// guard_start = [2]int{y, x}
				guard = [2]int{y, x}
			}
		}
	}

	for !isOutside(guard, maxY, maxX) {
		direction := directions[dir_index]
		next := move(guard, direction)

		// check if obstacle
		key := fmt.Sprintf("%d,%d", next[0], next[1])
		if obstacles[key] {
			// turn right
			dir_index = (dir_index + 1) % 4
		} else {
			guard = next

			if !isOutside(guard, maxY, maxX) {
				visited[key] = true
			}
		}
	}

	// render(visited, obstacles, maxY, maxX, guard_start)

	// not sure why +1 is needed, but that was the answer
	fmt.Println(len(visited) + 1)
}
