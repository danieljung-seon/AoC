package main

import (
	"2024/utils"
	"fmt"
	"slices"
	"strings"
)

// Returns true if a number array violates any rules
func isInvalid(rules [][]int, nums []int) bool {
	for _, rule := range rules {
		left_idx := slices.Index(nums, rule[0])
		right_idx := slices.Index(nums, rule[1])
		if left_idx != -1 && right_idx != -1 && left_idx >= right_idx {
			return true
		}
	}
	return false
}

// Fix a single array by swapping numbers until it satisfies all rules
func fixArray(nums []int, rules [][]int) []int {
	result := make([]int, len(nums))
	copy(result, nums)

	for isInvalid(rules, result) {
		for _, rule := range rules {
			left_idx := slices.Index(result, rule[0])
			right_idx := slices.Index(result, rule[1])

			if left_idx != -1 && right_idx != -1 && left_idx >= right_idx {
				result[left_idx], result[right_idx] = result[right_idx], result[left_idx]
				break
			}
		}
	}
	return result
}

func solve_05_2() {
	input, err := utils.ReadInput("05/05.0.input.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(input, "\n\n")
	rules := parse_rules_05(strings.Split(parts[0], "\n"))
	numbers := parse_numbers_05(strings.Split(parts[1], "\n"))

	sum := 0
	for _, nums := range numbers {
		if isInvalid(rules, nums) {
			fixed := fixArray(nums, rules)
			sum += fixed[len(fixed)/2]
		}
	}

	fmt.Println(sum)
}
