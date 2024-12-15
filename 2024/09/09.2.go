package main

import (
	"2024/utils"
	"fmt"
	"strconv"
	"strings"
)

type Config struct {
	start int
	end   int
	len   int
	val   int
}

func solve_09_2() {
	input, err := utils.ReadInput("09/09.0.input.txt")
	if err != nil {
		panic(err)
	}

	raws := strings.Split(input, "")
	nums := make([]int, 0)
	values := make([]Config, 0)
	spaces := make([]Config, 0)

	for idx, raw := range raws {
		num, err := strconv.Atoi(raw)
		if err != nil {
			panic(err)
		}

		if idx%2 != 0 {
			spaces = append(spaces, Config{
				start: len(nums),
				end:   len(nums) + num - 1,
				len:   num,
				val:   idx / 2,
			})
		} else {
			values = append(values, Config{
				start: len(nums),
				end:   len(nums) + num - 1,
				len:   num,
				val:   idx / 2,
			})
		}

		for i := 0; i < num; i++ {
			if idx%2 == 0 {
				nums = append(nums, idx/2)
			} else {
				nums = append(nums, -1)
			}
		}
	}

	b := len(values)
	for b > 0 {
		find_len := values[b-1].len

		for i := 0; i < len(spaces); i++ {
			if spaces[i].len >= find_len {
				if spaces[i].start >= values[b-1].start {
					break
				}

				for v := 0; v < find_len; v++ {
					nums[spaces[i].start+v] = values[b-1].val
					nums[values[b-1].start+v] = -1
				}

				spaces[i].len -= find_len
				spaces[i].start += find_len
				spaces[i].end += find_len
				break
			}
		}

		b--
	}

	sum := 0
	str := ""
	for idx, num := range nums {
		if num != -1 {
			sum += num * idx
			str += strconv.Itoa(num)
		} else {
			str += "."
		}
	}

	fmt.Println(sum)
}

// 8587288893605 -> too high
