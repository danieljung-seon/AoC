package main

import (
	"2024/utils"
	"fmt"
)

func solve_00_2() {
	input, err := utils.ReadInput("00/00.0.input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(input)
}
