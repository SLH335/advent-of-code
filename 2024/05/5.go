package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func a(input []string) {
	rules, updates := parse(input)

	sumCorrectMiddles := 0
	for _, update := range updates {
		correctOrder := true
		for j := 0; j < len(update)-1; j++ {
			for _, rule := range rules {
				if rule[0] == update[j+1] && rule[1] == update[j] {
					correctOrder = false
					break
				}
			}
		}
		if correctOrder {
			sumCorrectMiddles += update[len(update)/2]
		}
	}
	fmt.Println(sumCorrectMiddles)
}

func b(input []string) {
	rules, updates := parse(input)

	sumCorrectedMiddles := 0
	for i, update := range updates {
		initiallyCorrectOrder := true
		for {
			foundMistake := false
			for j := 0; j < len(update)-1; j++ {
				for _, rule := range rules {
					if rule[0] == update[j+1] && rule[1] == update[j] {
						tmp := updates[i][j]
						updates[i][j] = updates[i][j+1]
						updates[i][j+1] = tmp
						foundMistake = true
						initiallyCorrectOrder = false
					}
				}
			}
			if !foundMistake {
				break
			}
		}
		if !initiallyCorrectOrder {
			sumCorrectedMiddles += updates[i][len(update)/2]
		}
	}
	fmt.Println(sumCorrectedMiddles)
}

func parse(input []string) (rules [][]int, updates [][]int) {
	parsingRules := true
	for _, line := range input {
		if line == "" {
			parsingRules = false
			continue
		}

		if parsingRules {
			rule := atoi(strings.Split(line, "|"))
			rules = append(rules, rule)
		} else {
			update := atoi(strings.Split(line, ","))
			updates = append(updates, update)
		}
	}
	return rules, updates
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

func atoi(strings []string) (ints []int) {
	for _, str := range strings {
		integer, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		ints = append(ints, integer)
	}
	return ints
}
