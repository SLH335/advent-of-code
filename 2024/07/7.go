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
	var result int64 = 0
	for _, line := range input {
		numbers := strings.Split(line, " ")

		test, _ := strconv.ParseInt(numbers[0][:len(numbers[0])-1], 10, 64)
		values := atoi(numbers[1:])

		fmt.Printf("test: %d, values: %v\n", test, values)

		sum := values[0]

		possible := calculate(test, sum, values[1:], addition)
		if possible {
			//fmt.Println(result)
			result += test
			continue
		}
		possible = calculate(test, sum, values[1:], multiplication)
		if possible {
			//fmt.Println(line, "possible")
			//fmt.Println(result)
			result += test
			continue
		}
		possible = calculate(test, sum, values[1:], concatenation)
		if possible {
			//fmt.Println(line, "possible")
			//fmt.Println(result)
			result += test
		} else {
			//fmt.Println(line, "not possible")
		}
	}
	fmt.Println("result:", result)
}

func b(input []string) {

}

type operation int

const (
	addition operation = iota
	multiplication
	concatenation
)

func calculate(test int64, sum int64, vals []int64, op operation) (possible bool) {
	if len(vals) == 0 {
		return false
	}
	switch op {
	case addition:
		sum += vals[0]
	case multiplication:
		sum *= vals[0]
	case concatenation:
		//fmt.Print("trying concatenation ", sum, " || ", vals[0])
		sum *= int64(math.Pow(10, float64(len(strconv.FormatInt(vals[0], 10)))))
		sum += vals[0]
		//fmt.Printf(" = %d\n", sum)
	}
	if test == sum {
		fmt.Println("found solution:", test)
		return true
	}
	possible = calculate(test, sum, vals[1:], addition)
	if possible {
		return true
	}
	possible = calculate(test, sum, vals[1:], multiplication)
	if possible {
		return true
	}
	possible = calculate(test, sum, vals[1:], concatenation)
	return possible
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

func atoi(strings []string) (ints []int64) {
	for _, str := range strings {
		integer, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			panic(err)
		}
		ints = append(ints, integer)
	}
	return ints
}
