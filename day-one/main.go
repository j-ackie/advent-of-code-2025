package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var debugMode = flag.Bool("debug", false, "enable debug mode")
var includeAllClicks = flag.Bool("all-clicks", false, "include all clicks for password calculation")

const INITIAL_POSITION = 50
const NUM_POSITIONS = 100

func FindPassword(filename string, includeAllClicks bool, debug bool) int {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	fscanner := bufio.NewScanner(file)

	currentPosition := INITIAL_POSITION
	numZeros := 0

	for fscanner.Scan() {
		rotation := fscanner.Text()
		direction := rotation[0]
		steps, err := strconv.Atoi(rotation[1:])
		if err != nil {
			panic(err)
		}
		var newPosition int
		switch direction {
		case 'L':
			newPosition = currentPosition - steps
		case 'R':
			newPosition = currentPosition + steps
		}

		if includeAllClicks {
			var stepsToClick int
			if currentPosition == 0 {
				stepsToClick = NUM_POSITIONS
			} else if newPosition < currentPosition {
				stepsToClick = currentPosition
			} else {
				stepsToClick = NUM_POSITIONS - currentPosition
			}

			if steps >= stepsToClick {
				numZeros += ((steps - stepsToClick) / NUM_POSITIONS) + 1
			}

		} else if (newPosition%NUM_POSITIONS+NUM_POSITIONS)%NUM_POSITIONS == 0 {
			numZeros++
		}

		currentPosition = (newPosition%NUM_POSITIONS + NUM_POSITIONS) % NUM_POSITIONS

		if debug {
			fmt.Println("After", rotation, "position is", currentPosition)
			fmt.Println("\tNumber of zeros:", numZeros)
		}
	}

	return numZeros
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Usage: main [flags] <filename>")
		os.Exit(1)
	}
	filename := args[0]
	password := FindPassword(filename, *includeAllClicks, *debugMode)
	println("The password is:", password)
}
