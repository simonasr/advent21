package p2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonasr/advent21/internal/input"
)

func Run() {
	fmt.Printf("2 Part1: %v\n", part1())
}

func part1() int {
	var horizontal, vertical int = 0, 0
	input := input.ReadReport("p2-1-input.txt")
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
