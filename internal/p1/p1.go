package p1

import (
	"fmt"
	"strconv"

	"github.com/simonasr/advent21/internal/input"
)

var (
	report []int
)

func Run() {
	fmt.Printf("Part1: %v\n", P1Part1())
	fmt.Printf("Part2: %v", P1Part2())
}

func P1Part1() int {
	report = readReport("p1-1-report.txt")
	var timesIncreased int = 0

	for index, element := range report {
		if index == 0 {
			continue
		}

		if didDepthIncreased(report[index-1], element) {
			timesIncreased++
		}
	}
	return timesIncreased
}

func P1Part2() int {
	var timesIncreased int = 0
	report = readReport("p1-1-report.txt")
	reportLenght := len(report)
	for index := 0; index < reportLenght-3; index++ {
		sum1 := report[index] + report[index+1] + report[index+2]
		sum2 := report[index+1] + report[index+2] + report[index+3]

		if didDepthIncreased(sum1, sum2) {
			timesIncreased++
		}
	}

	return timesIncreased
}

func readReport(report string) []int {
	var result []int
	in := input.ReadReport("p1-1-report.txt")
	for _, line := range in {
		line, _ := strconv.Atoi(line)
		result = append(result, line)
	}
	return result
}

func didDepthIncreased(first, second int) bool {
	return first < second
}
