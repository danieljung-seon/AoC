package main

import (
	"2015/utils"
	"fmt"
)

func solve_1() {
	input, err := utils.ReadInput("01/01.0.input.txt")
	if err != nil {
		panic(err)
	}

	floor := 0
	for _, char := range input {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
	}

	fmt.Println(floor)
}
