package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func a(input []string) {
	reports := [][]int{}
	for _, line := range input {
		if line == "" {
			continue
		}
		levels := []int{}
		for _, level := range strings.Split(line, " ") {
			n, err := strconv.Atoi(level)
			if err != nil {
				fmt.Println("Error reading level", level)
				continue
			}
			levels = append(levels, n)
		}
		reports = append(reports, levels)
	}

	var safeReports int
	for _, report := range reports {
		var lastLevel int
		ascending := true
		safe := true
		for i, level := range report {
			diff := int(math.Abs(float64(level - lastLevel)))
			if lastLevel != 0 && (diff < 1 || diff > 3) {
				safe = false
				break
			}
			if i == 1 {
				ascending = (level - lastLevel) > 0
			}
			if (level - lastLevel) > 0 != ascending {
				safe = false
				break
			}
			lastLevel = level
		}
		if safe {
			safeReports++
		}
	}
	fmt.Println(safeReports)
}

func b(input []string) {
	reports := getReports(input)

	safeReports := 0
	for _, report := range reports {
		safe := isSafe(report)
		if safe {
			safeReports++
		} else {
			for i := 0; i < len(report); i++ {
				toleranceReport := removeIndex(report, i)
				safe = isSafe(toleranceReport)
				if safe {
					safeReports++
					break
				}
			}
		}
	}
	fmt.Println(safeReports)
}

func getReports(input []string) [][]int {
	reports := [][]int{}
	for _, line := range input {
		if line == "" {
			continue
		}
		levels := []int{}
		for _, level := range strings.Split(line, " ") {
			n, err := strconv.Atoi(level)
			if err != nil {
				fmt.Println("Error reading level", level)
				continue
			}
			levels = append(levels, n)
		}
		reports = append(reports, levels)
	}
	return reports
}

func isSafe(report []int) bool {
	var lastLevel int
	ascending := true
	safe := true
	for j, level := range report {
		if j == 1 {
			ascending = (level - lastLevel) > 0
		}
		diff := int(math.Abs(float64(level - lastLevel)))
		if lastLevel != 0 && (diff < 1 || diff > 3) || (level-lastLevel) > 0 != ascending {
			safe = false
			break
		}
		lastLevel = level
	}
	return safe
}

func removeIndex(report []int, index int) []int {
	newReport := []int{}
	for i := 0; i < len(report); i++ {
		if i != index {
			newReport = append(newReport, report[i])
		}
	}
	return newReport
}

func main() {
	args := os.Args[1:]

	inputFile := "input.txt"
	if len(args) > 1 && args[1] == "easy" {
		inputFile = "input_easy.txt"
	}

	input := readFile(inputFile)

	if len(args) == 0 || args[0] == "a" {
		a(input)
	} else if args[0] == "b" {
		b(input)
	} else {
		fmt.Println("Invalid argument, speficy a or b")
		os.Exit(1)
	}
}

func readFile(name string) []string {
	// open input file
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	// add each line to string array
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
