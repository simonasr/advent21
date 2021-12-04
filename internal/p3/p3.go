package p3

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/simonasr/advent21/internal/input"
)

func Run() {
	input := input.ReadReport("p3-1-input.txt")

	fmt.Printf("3 Part1: %v\n", Part1(input))
	// fmt.Printf("3 Part2: %v\n", Part2(input))
}

func Part1(input []string) int {
	Occurencies := []int{0}
	linesTotal := len(input)
	for _, line := range input {
		for index, b := range line {
			Bit, _ := strconv.Atoi(string(b))
			if index < len(Occurencies) {
				Occurencies[index] += Bit
			} else {
				Occurencies = append(Occurencies, Bit)
			}
		}
	}
	most, least := mostCommon(linesTotal, Occurencies)
	gammaRate := gammaRate(most)
	epsilonRate := epsilonRate(least)
	powerConsumption := gammaRate * epsilonRate
	return int(powerConsumption)
}

func gammaRate(numbers []int) int64 {
	gammaBin := gammaRateBin(numbers)
	gammaDecimal, _ := strconv.ParseInt(gammaBin, 2, 64)
	return gammaDecimal
}

func epsilonRate(numbers []int) int64 {
	gammaBin := gammaRateBin(numbers)
	gammaDecimal, _ := strconv.ParseInt(gammaBin, 2, 64)
	return gammaDecimal
}

func gammaRateBin(numbers []int) string {
	var bin []string
	for _, i := range numbers {
		bin = append(bin, strconv.Itoa(i))
	}
	gammaString := strings.Join(bin, "")
	return gammaString
}

func mostCommon(total int, countOnes []int) ([]int, []int) {
	var most, least []int
	var res int
	for _, v := range countOnes {
		// prevent division by zero
		if v == 0 {
			res = 0
		} else {
			res = total / v
		}
		if res > 1 {
			most = append(most, 0)
			least = append(least, 1)
		} else {
			most = append(most, 1)
			least = append(least, 0)
		}
	}
	return most, least
}
