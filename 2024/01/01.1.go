package main

import (
	"2024/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func solve_1() {
	input, err := utils.ReadInput("01/01.0.input.txt")
	if err != nil {
		panic(err)
	}

	var list1 []int
	var list2 []int
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		num1, _ := strconv.Atoi(fields[0])
		num2, _ := strconv.Atoi(fields[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	sum := 0
	for idx, num := range list1 {
		diff := num - list2[idx]
		diff = int(math.Abs(float64(diff)))
		sum += diff
	}

	fmt.Println(sum)
}
