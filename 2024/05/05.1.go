package main

import (
	"2024/utils"
	"fmt"
	"slices"
	"strings"
)

// Returns true if a number array satisfies all rules
func isValid(rules [][]int, nums []int) bool {
	for _, rule := range rules {
		left_idx := slices.Index(nums, rule[0])
		right_idx := slices.Index(nums, rule[1])
		if left_idx != -1 && right_idx != -1 && left_idx >= right_idx {
			return false
		}
	}
	return true
}

func solve_05_1() {
	input, err := utils.ReadInput("05/05.0.input.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(input, "\n\n")
	rules := parse_rules_05(strings.Split(parts[0], "\n"))
	numbers := parse_numbers_05(strings.Split(parts[1], "\n"))

	sum := 0
	for _, nums := range numbers {
		if isValid(rules, nums) {
			sum += nums[len(nums)/2]
		}
	}

	fmt.Println(sum)
}
