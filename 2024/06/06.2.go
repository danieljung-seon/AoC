package main

import (
	"2024/utils"
	"fmt"
	"strings"
)

func solve_06_2() {
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

	obstacles := make(map[string]bool)
	// var guard_start [2]int
	var start [2]int
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
				start = [2]int{y, x}
			}
		}
	}

	loop_found_cnt := 0

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			// fmt.Printf("%d,%d\n", y, x)
			visited := make(map[string]bool)
			dir_index := 0
			guard := start
			loop_found := false

			for !isOutside(guard, maxY, maxX) && !loop_found {
				direction := directions[dir_index]
				next := move(guard, direction)

				key_visited := fmt.Sprintf("%d,%d,%d", next[0], next[1], dir_index)
				if visited[key_visited] {
					loop_found_cnt++
					loop_found = true
					// fmt.Println(y, x)
					break
				}

				// check if obstacle
				key := fmt.Sprintf("%d,%d", next[0], next[1])
				if obstacles[key] || (next[0] == y && next[1] == x) {
					// turn right
					dir_index = (dir_index + 1) % 4
				} else {
					guard = next

					if !isOutside(guard, maxY, maxX) {
						key_visited := fmt.Sprintf("%d,%d,%d", guard[0], guard[1], dir_index)
						visited[key_visited] = true
					}
				}
			}
		}
	}

	// render(visited, obstacles, maxY, maxX, guard_start)

	fmt.Println(loop_found_cnt)
}
