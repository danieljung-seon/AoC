package main

import (
	"2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func execute_07_1(op_left int, op_right int, operator string) int {
	if operator == "+" {
		return op_left + op_right
	}

	if operator == "*" {
		return op_left * op_right
	}

	panic("Invalid operator")
}

func solve_07_1_helper(nums []int, operator string, expected int) bool {
	if len(nums) == 2 {
		return execute_07_1(nums[0], nums[1], operator) == expected
	}

	op_left := nums[0]
	op_right := nums[1]
	partial_result := execute_07_1(op_left, op_right, operator)

	nums = append([]int{partial_result}, nums[2:]...)

	return solve_07_1_helper(nums, "+", expected) || solve_07_1_helper(nums, "*", expected)
}

func test_07_1(nums []int, expected int) bool {
	return solve_07_1_helper(nums, "+", expected) || solve_07_1_helper(nums, "*", expected)
}

func solve_07_1() {
	input, err := utils.ReadInput("07/07.0.input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		expected, _ := strconv.Atoi(parts[0])
		inputs := strings.Split(parts[1], " ")

		nums := []int{}
		for _, input := range inputs {
			num, _ := strconv.Atoi(input)
			nums = append(nums, num)
		}

		res := test_07_1(nums, expected)
		if res {
			sum += expected
		}
	}

	fmt.Println(sum)
}
