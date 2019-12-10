package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y float64
}

func main() {
	file, err := os.Open("/Users/alex/go/projects/advent_of_code_2019/day_3/input.txt")

	if err != nil {
		fmt.Println("Error opening input file.")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	wire1Input := strings.Split(scanner.Text(), ",")

	scanner.Scan()

	wire2Input := strings.Split(scanner.Text(), ",")

	wire1Points := getPoints(wire1Input)
	wire2Points := getPoints(wire2Input)

	shortestDistance := -1

	for point := range wire1Points {
		if wire2Points[point] {
			distance := manhattanDistanceFromOrigin(point)

			if distance < shortestDistance || shortestDistance < 0  {
				shortestDistance = distance
			}
		}
	}

	fmt.Println("Shortest distance: ", shortestDistance)
}

func getPoints(wireInput []string) map[Point]bool {
	var x, y float64
	points := make(map[Point]bool)

	for _, input := range wireInput {
		direction := input[0:1]
		step, err := strconv.Atoi(input[1:])

		if err != nil {
			panic(err)
		}

		for i:= 0; i < step; i++ {
			if direction == "R" {
				x = x + 1
			} else if direction == "L" {
				x = x - 1
			} else if direction == "D" {
				y = y - 1
			} else if direction == "U" {
				y = y + 1
			}

			points[Point{x, y}] = true
		}
	}

	return points
}

func manhattanDistanceFromOrigin(point Point) int {
	return int(math.Abs(point.x) + math.Abs(point.y))
}
