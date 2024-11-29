package main

import (
	"2015/utils"
	"fmt"
)

func solve_2() {
	input, err := utils.ReadInput("01/01.0.input.txt")
	if err != nil {
		panic(err)
	}

	floor := 0
	for i, char := range input {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}

		if floor == -1 {
			fmt.Println(i + 1)
			break
		}
	}
}
