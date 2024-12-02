package main

import (
	"2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func solve_2() {
	input, err := utils.ReadInput("01/01.0.input.txt")
	if err != nil {
		panic(err)
	}

	var listLeft []int
	mapRight := make(map[int]int)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		num1, _ := strconv.Atoi(fields[0])
		listLeft = append(listLeft, num1)

		num2, _ := strconv.Atoi(fields[1])
		mapRight[num2]++
	}

	sum := 0
	for _, num := range listLeft {
		sum += num * mapRight[num]
	}

	fmt.Println(sum)
}
