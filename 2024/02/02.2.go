package main

import (
	"2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func is_valid_line_2_2(line []string) bool {
	fields := make([]float64, len(line))
	for i, f := range line {
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

func solve_2_2() {
	input, err := utils.ReadInput("02/02.0.input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		has_valid := false
		fields_str := strings.Fields(line)
		// fmt.Println("---", fields_str)
		fields_count := len(fields_str)
		for i := 0; i < fields_count; i++ {
			fields := make([]string, len(fields_str))
			copy(fields, fields_str)

			// fmt.Println("idx", i, fields[0:i], fields[i+1:])
			line_with_missing := append(fields[0:i], fields[i+1:]...)

			if is_valid_line_2_2(line_with_missing) {
				// fmt.Println(line_with_missing)
				has_valid = true
			}
		}

		if has_valid {
			sum++
		}
	}

	fmt.Println(sum)
}
