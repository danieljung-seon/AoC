package main

import (
	"2024/utils"
	"fmt"
	"strings"
)

func count_x_mas(letters map[int]map[int]string, start [2]int) int {
	if letters[start[0]][start[1]] != "A" {
		return 0
	}

	// Define the 4 possible patterns
	patterns := [][4][2]string{
		{{"M", "M"}, {"S", "S"}}, // M M, S S
		{{"S", "S"}, {"M", "M"}}, // S S, M M
		{{"S", "M"}, {"S", "M"}}, // S M, S M
		{{"M", "S"}, {"M", "S"}}, // M S, M S
	}

	count := 0
	for _, pattern := range patterns {
		if letters[start[0]-1][start[1]-1] == pattern[0][0] &&
			letters[start[0]-1][start[1]+1] == pattern[0][1] &&
			letters[start[0]+1][start[1]-1] == pattern[1][0] &&
			letters[start[0]+1][start[1]+1] == pattern[1][1] {
			count++
		}
	}

	return count
}

func solve_04_2() {
	input, err := utils.ReadInput("04/04.0.input.txt")
	if err != nil {
		panic(err)
	}

	// Parse input into grid and find all A positions
	letters := make(map[int]map[int]string)
	a_positions := make([][2]int, 0)

	for y, line := range strings.Split(input, "\n") {
		letters[y] = make(map[int]string)
		for x, char := range line {
			letters[y][x] = string(char)
			if string(char) == "A" {
				a_positions = append(a_positions, [2]int{y, x})
			}
		}
	}

	// Count valid XMAS patterns around each A
	sum := 0
	for _, pos := range a_positions {
		sum += count_x_mas(letters, pos)
	}

	fmt.Println(sum)
}
