package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func a(input []string) {
	sum := 0
	check := "XMAS"
	for i, line := range input {
		for j, char := range line {
			if char != rune(check[0]) {
				continue
			}
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					for m := 1; m < len(check); m++ {
						x := i + k*m
						y := j + l*m
						if x < 0 || x >= len(input) {
							break
						}
						if y < 0 || y >= len(line) {
							break
						}
						if input[x][y] != byte(check[m]) {
							break
						}
						if m == 3 {
							sum++
						}
					}
				}
			}
		}
	}
	fmt.Println(sum)
}

func b(input []string) {
	sum := 0
	for i, line := range input {
		for j, char := range line {
			if char != 'A' {
				continue
			}
			if i-1 < 0 || j-1 < 0 || i+1 >= len(input) || j+1 >= len(line) {
				continue
			}
			tl, tr, bl, br := input[i+1][j+1], input[i+1][j-1], input[i-1][j+1], input[i-1][j-1]

			if (tl == 'M' && br == 'S' || tl == 'S' && br == 'M') &&
				(tr == 'M' && bl == 'S' || tr == 'S' && bl == 'M') {
				sum++
			}
		}
	}
	fmt.Println(sum)
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
