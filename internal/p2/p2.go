package p2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonasr/advent21/internal/input"
)

func Run() {
	input := input.ReadReport("p2-1-input.txt")

	fmt.Printf("2 Part1: %v\n", Part1(input))
	fmt.Printf("2 Part2: %v\n", Part2(input))
}

func Part1(input []string) int {
	var horizontal, vertical int = 0, 0
	for _, line := range input {
		x, y := goingTo(line)
		horizontal += x
		vertical += y
	}
	return horizontal * vertical
}

func goingTo(input string) (int, int) {
	var x, y int = 0, 0
	movement := strings.Split(input, " ")
	direction := movement[0]
	distance, _ := strconv.Atoi(movement[1])
	switch direction {
	case "forward":
		x += distance
	case "up":
		y -= distance
	case "down":
		y += distance
	}
	return x, y
}

func Part2(input []string) int {
	var horizontal, depth, aim int = 0, 0, 0
	for _, line := range input {
		x, y, z := goingToV2(line, aim)
		horizontal += x
		depth += y
		aim += z
	}
	return horizontal * depth
}

func goingToV2(input string, aim int) (int, int, int) {
	var x, y, z int = 0, 0, 0
	movement := strings.Split(input, " ")
	direction := movement[0]
	distance, _ := strconv.Atoi(movement[1])
	switch direction {
	case "forward":
		x += distance
		y += aim * distance
	case "up":
		z -= distance
	case "down":
		z += distance
	}
	return x, y, z
}
