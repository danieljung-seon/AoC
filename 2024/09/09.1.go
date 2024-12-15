package main

import (
	"2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func solve_09_1() {
	input, err := utils.ReadInput("09/09.0.input.txt")
	if err != nil {
		panic(err)
	}

	raws := strings.Split(input, "")
	nums := make([]int, 0)

	for idx, raw := range raws {
		num, err := strconv.Atoi(raw)
		if err != nil {
			panic(err)
		}

		for i := 0; i < num; i++ {
			if idx%2 == 0 {
				nums = append(nums, idx/2)
			} else {
				nums = append(nums, -1)
			}
		}
	}

	sum := 0
	idx := 0
	until := len(nums)

	for idx < until {
		num := nums[idx]

		if num == -1 {
			last := nums[until-1]
			if last == -1 {
				until--
				continue
			}

			sum += idx * last
			until--
		} else {
			sum += idx * num
		}

		idx++
	}

	fmt.Println(sum)
}
