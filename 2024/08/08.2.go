package main

import (
	"2024/utils"
	"fmt"
	"slices"
	"strings"
)

func get_anti_nodes_for_pair_2(position_a Position, position_b Position, max_x int, max_y int) []Position {
	// Calculate the direction vector
	diff_x := position_b.x - position_a.x
	diff_y := position_b.y - position_a.y

	// Get the unit step
	step_x := diff_x
	step_y := diff_y

	result := []Position{position_a}

	// Move forward from position_a
	curr_x := position_a.x + step_x
	curr_y := position_a.y + step_y
	for curr_x >= 0 && curr_x < max_x && curr_y >= 0 && curr_y < max_y {
		if curr_x != position_b.x || curr_y != position_b.y {
			result = append(result, Position{x: curr_x, y: curr_y})
		}
		curr_x += step_x
		curr_y += step_y
	}

	// Move backward from position_a
	curr_x = position_a.x - step_x
	curr_y = position_a.y - step_y
	for curr_x >= 0 && curr_x < max_x && curr_y >= 0 && curr_y < max_y {
		result = append(result, Position{x: curr_x, y: curr_y})
		curr_x -= step_x
		curr_y -= step_y
	}

	return result
}

func get_anti_nodes_2(nodes map[string][]Position, max_x int, max_y int) map[string][]Position {
	anti_nodes := make(map[string][]Position)

	for key, node := range nodes {
		for _, position_a := range node {
			for _, position_b := range node {
				// Skip the same position
				if position_a.x == position_b.x && position_a.y == position_b.y {
					continue
				}

				anti_nodes_for_pair := get_anti_nodes_for_pair_2(position_a, position_b, max_x, max_y)
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

func parse_08_2(input string) map[string][]Position {
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

func render_08_2(positions []Position, max_x int, max_y int) {
	for y := 0; y < max_y; y++ {
		for x := 0; x < max_x; x++ {
			if slices.ContainsFunc(positions, func(p Position) bool {
				return p.x == x && p.y == y
			}) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func solve_08_2() {
	input, err := utils.ReadInput("08/08.0.input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(input, "\n")
	max_x := len(lines[0])
	max_y := len(lines)

	nodes := parse_08_2(input)
	anti_nodes := get_anti_nodes_2(nodes, max_x, max_y)

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

	// render_08_2(unique_positions, max_x, max_y)

	fmt.Println(len(unique_positions))
}
