package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("/Users/alex/go/projects/advent_of_code/day_1/input.txt")

	if err != nil {
		fmt.Println("Error opening input file.")
		return
	}

	scanner := bufio.NewScanner(file)

	var sum float64

	for scanner.Scan() {
		mass, parseErr := strconv.ParseFloat(scanner.Text(), 64)

		if parseErr != nil {
			fmt.Println("Error parsing mass.")
			break
		}

		sum += math.Floor(mass / 3) - 2
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Error reading from input file.")
	}

	fmt.Println(int(sum))
}
