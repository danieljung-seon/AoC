package main

import (
	"strconv"
	"strings"
)

func parse_rules_05(rules_raw []string) [][]int {
	rules := make([][]int, len(rules_raw))
	for i, rule := range rules_raw {
		parts := strings.Split(rule, "|")
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])
		rules[i] = []int{left, right}
	}
	return rules
}

func parse_numbers_05(numbers_raw []string) [][]int {
	numbers := make([][]int, len(numbers_raw))

	for i, number := range numbers_raw {
		parts := strings.Split(number, ",")
		for _, part := range parts {
			number, _ := strconv.Atoi(part)
			numbers[i] = append(numbers[i], number)
		}
	}

	return numbers
}
