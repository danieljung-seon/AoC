package main

import (
	"2024/utils"
	"fmt"
	"regexp"
	"strconv"
)

func solve_3_2() {
	input, err := utils.ReadInput("03/03.0.input.txt")
	if err != nil {
		panic(err)
	}

	mulPattern := `mul\((\d+),(\d+)\)|don't\(\)|do\(\)`
	re := regexp.MustCompile(mulPattern)

	matches := re.FindAllStringSubmatch(input, -1)
	do := true

	sum := 0
	for _, match := range matches {
		instruction := match[0]
		if instruction == "don't()" {
			do = false
			continue
		}

		if instruction == "do()" {
			do = true
			continue
		}

		if do {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			sum += num1 * num2
		}
	}

	fmt.Println(sum)
}
