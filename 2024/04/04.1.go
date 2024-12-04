package main

import (
	"2024/utils"
	"fmt"
	"strings"
)

// Check if sequence matches XMAS starting from given position and direction
func check_xmas(letters map[int]map[int]string, start [2]int, dy, dx int) bool {
	pattern := []string{"X", "M", "A", "S"}
	y, x := start[0], start[1]

	for i := 0; i < 4; i++ {
		if letters[y][x] != pattern[i] {
			return false
		}
		y += dy
		x += dx
	}
	return true
}

func solve_04_1() {
	input, err := utils.ReadInput("04/04.0.input.txt")
	if err != nil {
		panic(err)
	}

	// Parse input into map
	letters := make(map[int]map[int]string)
	letters_x := make([][2]int, 0)

	for y, line := range strings.Split(input, "\n") {
		letters[y] = make(map[int]string)
		for x, c := range line {
			letters[y][x] = string(c)
			if string(c) == "X" {
				letters_x = append(letters_x, [2]int{y, x})
			}
		}
	}

	// Directions to check: right, down, diagonal down-right, diagonal down-left
	directions := [][2]int{{0, 1}, {1, 0}, {1, 1}, {1, -1}}

	sum := 0
	for _, pos := range letters_x {
		// Check forward and backward from each X
		for _, dir := range directions {
			if check_xmas(letters, pos, dir[0], dir[1]) {
				sum++
			}
			// Check reverse direction
			if check_xmas(letters, pos, -dir[0], -dir[1]) {
				sum++
			}
		}
	}

	fmt.Println(sum)
}
