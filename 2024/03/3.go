package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func a(input string) {
	sum := 0
	for i := 0; i < len(input); i++ {
		searchLen := min(len(input)-i, 12)
		matched, _ := regexp.MatchString("^mul\\([0-9]{1,3},[0-9]{1,3}\\).*", input[i:i+searchLen])
		if matched {
			mul := strings.Replace(strings.Split(input[i:i+searchLen], ")")[0], "mul(", "", 1)
			nums := atoi(strings.Split(mul, ","))
			sum += nums[0] * nums[1]
		}
	}

	fmt.Println(sum)
}

func b(input string) {
	enabled := true
	sum := 0
	for i := 0; i < len(input); i++ {
		searchLen := min(len(input)-i, 12)
		if strings.HasPrefix(input[i:i+searchLen], "don't()") {
			enabled = false
			continue
		} else if strings.HasPrefix(input[i:i+searchLen], "do()") {
			enabled = true
			continue
		}
		if !enabled {
			continue
		}
		matched, _ := regexp.MatchString("^mul\\([0-9]{1,3},[0-9]{1,3}\\).*", input[i:i+searchLen])
		if matched {
			mul := strings.Replace(strings.Split(input[i:i+searchLen], ")")[0], "mul(", "", 1)
			nums := atoi(strings.Split(mul, ","))
			sum += nums[0] * nums[1]
		}
	}

	fmt.Println(sum)
}

func main() {
	args := os.Args[1:]

	inputFile := "input.txt"
	if len(args) > 1 && args[1] == "easy" {
		if args[0] == "a" {
			inputFile = "input_easy_a.txt"
		} else {
			inputFile = "input_easy_b.txt"
		}
	}

	input := strings.Join(readFile(inputFile), "\n")

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
