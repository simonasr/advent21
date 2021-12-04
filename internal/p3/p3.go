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
	fmt.Printf("3 Part2: %v\n", Part2(input))
}

func Part1(input []string) int {
	linesTotal := len(input)
	most, least := mostCommon(linesTotal, countOnes(input))
	gammaRate := gammaRate(most)
	epsilonRate := epsilonRate(least)
	powerConsumption := gammaRate * epsilonRate

	return int(powerConsumption)
}

func Part2(input []string) int {
	oxygen := GetRating(input, "oxygen")
	co2 := GetRating(input, "co2")
	result := oxygen * co2

	return int(result)
}

func countOnes(input []string) []int {
	ones := []int{0}
	for _, line := range input {
		for index, b := range line {
			Bit, _ := strconv.Atoi(string(b))
			if index < len(ones) {
				ones[index] += Bit
			} else {
				ones = append(ones, Bit)
			}
		}
	}
	return ones
}

func gammaRate(numbers []int) int64 {
	gammaBin := numbersToBin(numbers)
	gammaDecimal, _ := strconv.ParseInt(gammaBin, 2, 64)
	return gammaDecimal
}

func epsilonRate(numbers []int) int64 {
	gammaBin := numbersToBin(numbers)
	gammaDecimal, _ := strconv.ParseInt(gammaBin, 2, 64)
	return gammaDecimal
}

func numbersToBin(numbers []int) string {
	var bin []string
	for _, i := range numbers {
		bin = append(bin, strconv.Itoa(i))
	}
	gammaString := strings.Join(bin, "")
	return gammaString
}

func mostCommon(total int, ones []int) ([]int, []int) {
	var most, least []int
	var res int
	for _, v := range ones {
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

func GetRating(list []string, ratingType string) int {
	var filteredInput, newFilteredList []string
	var listZeros, listOnes []string
	filteredInput = list
	lineLength := len(list[0])

	for offset := 0; offset < lineLength; offset++ {
		for _, line := range filteredInput {
			bitAtOffset, _ := strconv.Atoi(string(line[offset]))
			if bitAtOffset == 1 {
				listOnes = append(listOnes, line)
			} else {
				listZeros = append(listZeros, line)
			}
		}

		if len(listOnes) >= len(listZeros) {
			switch ratingType {
			case "oxygen":
				newFilteredList = listOnes
			case "co2":
				newFilteredList = listZeros
			}
		} else {
			switch ratingType {
			case "oxygen":
				newFilteredList = listZeros
			case "co2":
				newFilteredList = listOnes
			}
		}
		if len(newFilteredList) == 1 {
			break
		}
		filteredInput = newFilteredList
		newFilteredList = newFilteredList[:0]
		listOnes = listOnes[:0]
		listZeros = listZeros[:0]
	}
	result, _ := strconv.ParseInt(newFilteredList[0], 2, 64)
	return int(result)
}
