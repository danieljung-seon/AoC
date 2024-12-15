package main

type Point struct {
	x int
	y int
	v int
}

func get_value(parsed [][]int, point Point) int {
	if point.y < 0 || point.y >= len(parsed) || point.x < 0 || point.x >= len(parsed[0]) {
		return -1
	}

	return parsed[point.y][point.x]
}
