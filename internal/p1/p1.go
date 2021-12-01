package p1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	Report         []int
	TimesIncreased int = 0
)

func Run() {
	Report = readReport()

	for index, element := range Report {
		if index == 0 {
			continue
		}

		if didDepthIncreased(Report[index-1], element) {
			TimesIncreased++
		}
	}
	fmt.Print(TimesIncreased)
}

func readReport() []int {
	var lines []int
	file, err := os.Open("p1-1-report.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		depth := scanner.Text()
		i, _ := strconv.Atoi(depth)
		lines = append(lines, i)
	}
	return lines
}

func didDepthIncreased(first, second int) bool {
	return first < second
}
