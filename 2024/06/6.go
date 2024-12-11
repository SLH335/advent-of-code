package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func a(input []string) {
	currPos := pos{}
	nextPos := pos{}

	for x, line := range input {
		for y, char := range line {
			if !slices.Contains([]rune{'^', '>', 'v', '<'}, char) {
				continue
			}
			currPos = pos{x, y, char}
			nextPos = getNextPos(currPos)
		}
	}

	for nextPos.x >= 0 && nextPos.x < len(input) && nextPos.y >= 0 && nextPos.y < len(input[0]) {
		switch input[nextPos.x][nextPos.y] {
		case '.', 'X':
			replaceChar(input, currPos.x, currPos.y, 'X')
			replaceChar(input, nextPos.x, nextPos.y, nextPos.dir)
			currPos = nextPos
			nextPos = getNextPos(nextPos)
		case '#':
			replaceChar(input, currPos.x, currPos.y, nextDir(currPos.dir))
			nextPos = getNextPos(pos{currPos.x, currPos.y, nextDir(currPos.dir)})
		}
	}
	replaceChar(input, currPos.x, currPos.y, 'X')

	visitedPositions := 0
	for _, line := range input {
		for _, char := range line {
			if char == 'X' {
				visitedPositions++
			}
		}
	}
	fmt.Println(visitedPositions)
}

func b(input []string) {
	locationsRaw := findPossibleObstacleLocations(input)

	fmt.Println("found", len(locationsRaw), "locations including duplicates")

	locations := []pos{}
	for _, loc1 := range locationsRaw {
		alreadyPresent := false
		for _, loc2 := range locations {
			if loc1 == loc2 {
				alreadyPresent = true
			}
		}
		if !alreadyPresent {
			locations = append(locations, loc1)
		}
	}
	fmt.Println("found", len(locations), "distinct locations")

	loopLocations := 0
	for i, location := range locations {
		if i%50 == 0 {
			fmt.Println("checked", i, "locations")
		}
		loop := checkLoopWithObstacle(input, location)
		if loop {
			loopLocations++
		}
	}

	fmt.Println(loopLocations)
}

type pos struct {
	x   int
	y   int
	dir rune
}

func getNextPos(currPos pos) (nextPos pos) {
	switch currPos.dir {
	case '^':
		nextPos = pos{currPos.x - 1, currPos.y, '^'}
	case '>':
		nextPos = pos{currPos.x, currPos.y + 1, '>'}
	case 'v':
		nextPos = pos{currPos.x + 1, currPos.y, 'v'}
	case '<':
		nextPos = pos{currPos.x, currPos.y - 1, '<'}
	}
	return nextPos
}

func nextDir(currDir rune) (nextDir rune) {
	switch currDir {
	case '^':
		return '>'
	case '>':
		return 'v'
	case 'v':
		return '<'
	case '<':
		return '^'
	default:
		panic("invalid dir")
	}
}

func replaceChar(input []string, x int, y int, newChar rune) {
	newLine := []rune(input[x])
	newLine[y] = newChar
	input[x] = string(newLine)
}

func findPositions(input []string) (currPos pos, nextPos pos) {
	currPos = pos{-1, -1, -1}
	for x, line := range input {
		for y, char := range line {
			if !slices.Contains([]rune{'^', '>', 'v', '<'}, char) {
				continue
			}
			currPos = pos{x, y, char}
			nextPos = getNextPos(currPos)
		}
	}
	return currPos, nextPos
}

func findPossibleObstacleLocations(inputOriginal []string) (locations []pos) {
	input := make([]string, len(inputOriginal))
	copy(input, inputOriginal)

	currPos, nextPos := findPositions(input)

	for nextPos.x >= 0 && nextPos.x < len(input) && nextPos.y >= 0 && nextPos.y < len(input[0]) {
		switch input[nextPos.x][nextPos.y] {
		case '.', 'X':
			replaceChar(input, currPos.x, currPos.y, 'X')
			replaceChar(input, nextPos.x, nextPos.y, nextPos.dir)
			currPos = nextPos
			nextPos = getNextPos(nextPos)
			locations = append(locations, pos{nextPos.x, nextPos.y, 0})
		case '#':
			replaceChar(input, currPos.x, currPos.y, nextDir(currPos.dir))
			nextPos = getNextPos(pos{currPos.x, currPos.y, nextDir(currPos.dir)})
		}
	}
	replaceChar(input, currPos.x, currPos.y, 'X')

	return locations
}

func checkLoopWithObstacle(inputOriginal []string, location pos) bool {
	input := make([]string, len(inputOriginal))
	copy(input, inputOriginal)

	if location.x < 0 || location.x >= len(input) || location.y < 0 || location.y >= len(input[0]) {
		return false
	}
	replaceChar(input, location.x, location.y, 'O')
	//fmt.Printf("checking with obstacle:\n%s\n-----------------------\n", strings.Join(input, "\n"))

	turningPoints := []pos{}

	currPos, nextPos := findPositions(input)
	if currPos.dir == -1 {
		return false
	}

	for nextPos.x >= 0 && nextPos.x < len(input) && nextPos.y >= 0 && nextPos.y < len(input[0]) {
		switch input[nextPos.x][nextPos.y] {
		case '.':
			replaceChar(input, currPos.x, currPos.y, '.')
			replaceChar(input, nextPos.x, nextPos.y, nextPos.dir)
			currPos = nextPos
			nextPos = getNextPos(nextPos)
		case '#', 'O':
			nextDir := nextDir(currPos.dir)
			replaceChar(input, currPos.x, currPos.y, nextDir)
			currPos.dir = nextDir
			nextPos = getNextPos(currPos)

			for _, turningPoint := range turningPoints {
				if turningPoint.x == currPos.x && turningPoint.y == currPos.y && turningPoint.dir == currPos.dir {
					//fmt.Println("identical turning point:", turningPoint)
					return true
				}
			}
			turningPoints = append(turningPoints, currPos)
		}
	}

	return false
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
