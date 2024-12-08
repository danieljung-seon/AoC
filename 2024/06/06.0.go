package main

import (
	"fmt"
)

func isOutside(guard [2]int, maxY, maxX int) bool {
	x := guard[0]
	y := guard[1]

	return x < 0 || x > maxX-1 || y < 0 || y > maxY-1
}

func move(guard [2]int, direction [2]int) [2]int {
	y := guard[0]
	x := guard[1]
	dy := direction[0]
	dx := direction[1]

	return [2]int{y + dy, x + dx}
}

func render(visited map[string]bool, obstacles map[string]bool, maxY, maxX int, guard_start [2]int) {
	fmt.Println("<pre style=\"word-wrap: break-word; white-space: pre-wrap;\">")
	for y := 0; y < maxY+1; y++ {
		for x := 0; x < maxX; x++ {
			key := fmt.Sprintf("%d,%d", y, x)
			if visited[key] {
				fmt.Print("X")
			} else if y == guard_start[0] && x == guard_start[1] {
				fmt.Print("^")
			} else if obstacles[key] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("</pre>")
}
