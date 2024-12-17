package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func a(input []string) {
	antinodes := getPositions(input, 1, 1)
	fmt.Println(len(antinodes))
}

func b(input []string) {
	antinodes := getPositions(input, 0, max(len(input), len(input[0])))
	fmt.Println(len(antinodes))
}

func getPositions(input []string, lower, upper int) (antinodes []pos) {
	antennas := parseAntennas(input)

	for _, frequency := range antennas {
		if len(frequency) < 2 {
			continue
		}
		for i, pos1 := range frequency {
			for j, pos2 := range frequency {
				if i == j {
					continue
				}
				nodes := getAntinodes(pos1, pos2, len(input), len(input[0]), lower, upper)
				antinodes = append(antinodes, nodes...)
			}
		}
	}
	slices.SortFunc(antinodes, func(a, b pos) int {
		if a.x < b.x {
			return -1
		} else if a.x > b.x {
			return 1
		}
		if a.y < b.y {
			return -1
		} else if a.y > b.y {
			return 1
		}
		return 0
	})
	antinodes = slices.Compact(antinodes)

	antinodes = slices.DeleteFunc(antinodes, func(a pos) bool {
		if a.x < 0 || a.y < 0 || a.x >= len(input) || a.y >= len(input[0]) {
			return true
		}
		return false
	})

	return antinodes
}

type pos struct {
	x int
	y int
}

func parseAntennas(input []string) (antennas map[rune][]pos) {
	antennas = map[rune][]pos{}

	for x, line := range input {
		for y, char := range line {
			if char == '.' {
				continue
			}
			antennas[char] = append(antennas[char], pos{x, y})
		}
	}
	return antennas
}

func (p1 pos) add(p2 pos) (res pos) {
	return pos{p1.x + p2.x, p1.y + p2.y}
}

func (p1 pos) sub(p2 pos) (res pos) {
	return pos{p1.x - p2.x, p1.y - p2.y}
}

func (p pos) mul(n int) (res pos) {
	return pos{p.x * n, p.y * n}
}

func getAntinodes(p1, p2 pos, xSize, ySize, lower, upper int) (antinodes []pos) {
	offset := p1.sub(p2)
	nodes := []pos{}

	for i := lower; i <= upper; i++ {
		nodes = append(nodes, p1.add(offset.mul(i)), p2.sub(offset.mul(i)))
	}

	for _, node := range nodes {
		if !(node.x < 0 || node.y < 0 || node.x >= xSize || node.y >= ySize) {
			antinodes = append(antinodes, node)
		}
	}

	return antinodes
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
