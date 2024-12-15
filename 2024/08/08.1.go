package main

import (
	"2024/utils"
	"fmt"
	"slices"
	"strings"
)

type Position struct {
	x int
	y int
}

//
// a....
// .....
// ..b..
// x: 0 - 2 = -2
// y: 0 - 2 = -2

//
// b....
// .....
// ..a..
// x: 2 - 0 = 2
// y: 2 - 0 = 2

//
// ..a..
// .....
// b....
// x: 2 - 0 = 2
// y: 0 - 2 = -2

//
//
// ..b..
// .....
// a....
//
// x: 0 - 2 = -2
// y: 2 - 0 = 2

func get_anti_nodes_for_pair(position_a Position, position_b Position, max_x int, max_y int) []Position {
	diff_x := position_a.x - position_b.x
	diff_y := position_a.y - position_b.y

	anti_node_1 := Position{x: position_a.x + diff_x, y: position_a.y + diff_y}
	anti_node_2 := Position{x: position_b.x - diff_x, y: position_b.y - diff_y}

	var result []Position
	if anti_node_1.x >= 0 && anti_node_1.x < max_x && anti_node_1.y >= 0 && anti_node_1.y < max_y {
		result = append(result, anti_node_1)
	}

	if anti_node_2.x >= 0 && anti_node_2.x < max_x && anti_node_2.y >= 0 && anti_node_2.y < max_y {
		result = append(result, anti_node_2)
	}

	return result
}

func get_anti_nodes(nodes map[string][]Position, max_x int, max_y int) map[string][]Position {
	anti_nodes := make(map[string][]Position)

	for key, node := range nodes {
		for _, position_a := range node {
			for _, position_b := range node {
				// Skip the same position
				if position_a.x == position_b.x && position_a.y == position_b.y {
					continue
				}

				anti_nodes_for_pair := get_anti_nodes_for_pair(position_a, position_b, max_x, max_y)
				for _, anti_node := range anti_nodes_for_pair {
					// Check if this anti_node is already in the list
					exists := slices.ContainsFunc(anti_nodes[key], func(p Position) bool {
						return p.x == anti_node.x && p.y == anti_node.y
					})

					if !exists {
						anti_nodes[key] = append(anti_nodes[key], anti_node)
					}
				}
			}
		}
	}

	return anti_nodes
}

func parse_08_1(input string) map[string][]Position {
	antennas := make(map[string][]Position)

	lines := strings.Split(input, "\n")
	for y, line := range lines {
		parts := strings.Split(line, "")

		for x, part := range parts {
			if part == "." {
				continue
			}

			antennas[part] = append(antennas[part], Position{x: x, y: y})
		}
	}

	return antennas
}

func solve_08_1() {
	input, err := utils.ReadInput("08/08.0.input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(input, "\n")
	max_x := len(lines[0])
	max_y := len(lines)

	nodes := parse_08_1(input)
	anti_nodes := get_anti_nodes(nodes, max_x, max_y)

	// Flatten and deduplicate anti_nodes
	var all_positions []Position
	for _, positions := range anti_nodes {
		all_positions = append(all_positions, positions...)
	}

	// Remove duplicates using slices.ContainsFunc
	var unique_positions []Position
	for _, pos := range all_positions {
		exists := slices.ContainsFunc(unique_positions, func(p Position) bool {
			return p.x == pos.x && p.y == pos.y
		})
		if !exists {
			unique_positions = append(unique_positions, pos)
		}
	}

	fmt.Println(len(unique_positions))
}
