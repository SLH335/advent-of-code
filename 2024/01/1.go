package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func a(input []string) {
	var diff int

	numsLeft := []int{}
	numsRight := []int{}

	for _, line := range input {
		n, err := strconv.Atoi(strings.Split(line, "   ")[0])
		if err != nil {
			log.Printf("Invalid number: '%s'", line)
		}
		m, err := strconv.Atoi(strings.Split(line, "   ")[1])
		if err != nil {
			log.Printf("Invalid number: '%s'", line)
		}
		numsLeft = append(numsLeft, n)
		numsRight = append(numsRight, m)
	}

	slices.Sort(numsLeft)
	slices.Sort(numsRight)

	for i := 0; i < len(numsLeft); i++ {
		diff += int(math.Abs(float64(numsLeft[i] - numsRight[i])))
	}

	fmt.Println(diff)
}

func b(input []string) {
	numsLeft := []int{}
	numsRight := []int{}

	for _, line := range input {
		n, err := strconv.Atoi(strings.Split(line, "   ")[0])
		if err != nil {
			log.Printf("Invalid number: '%s'", line)
		}
		m, err := strconv.Atoi(strings.Split(line, "   ")[1])
		if err != nil {
			log.Printf("Invalid number: '%s'", line)
		}
		numsLeft = append(numsLeft, n)
		numsRight = append(numsRight, m)
	}

	sim := 0
	for i := 0; i < len(numsLeft); i++ {
		n := numsLeft[i]
		for _, m := range numsRight {
			if n == m {
				sim += n
			}
		}
	}

	fmt.Println(sim)
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
