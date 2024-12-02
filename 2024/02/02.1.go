package main

import (
	"2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func is_valid_line_2_1(line string) bool {
	fields_str := strings.Fields(line)
	fields := make([]float64, len(fields_str))
	for i, f := range fields_str {
		fields[i], _ = strconv.ParseFloat(f, 64)
	}

	prev_sign := 0
	for idx, field := range fields {
		if idx < len(fields)-1 {
			signed_diff := field - fields[idx+1]
			diff := int(math.Abs(signed_diff))

			var sign int
			if signed_diff < 0 {
				sign = -1
			} else {
				sign = 1
			}

			if prev_sign != 0 && prev_sign != sign {
				return false
			} else {
				prev_sign = sign
			}

			if diff < 1 || diff > 3 {
				return false
			}
		}
	}

	return true
}

func solve_2_1() {
	input, err := utils.ReadInput("02/02.0.input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		if is_valid_line_2_1(line) {
			sum++
		}
	}

	fmt.Println(sum)
}
