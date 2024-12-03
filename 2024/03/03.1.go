package main

import (
	"2024/utils"
	"fmt"
	"regexp"
	"strconv"
)

func solve_3_1() {
	input, err := utils.ReadInput("03/03.0.input.txt")
	if err != nil {
		panic(err)
	}

	mulPattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(mulPattern)

	matches := re.FindAllStringSubmatch(input, -1)

	sum := 0
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += num1 * num2
	}

	fmt.Println(sum)
}
