package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/Users/alex/go/projects/advent_of_code_2019/day_2/input.txt")

	if err != nil {
		fmt.Println("Error opening input file.")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	stringNums := strings.Split(scanner.Text(), ",")

	var opcodes []int

	for _, v := range stringNums {
		opcode, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		opcodes = append(opcodes, opcode)
	}

	opcodes[1] = 12
	opcodes[2] = 2

	var count int

	for i, op := range opcodes {
		count++
		if count == 4 {
			operation := opcodes[i - 3]
			input1 := opcodes[i - 2]
			input2 := opcodes[i - 1]

			var result int

			if operation == 1 {
				result = opcodes[input1] + opcodes[input2]
			} else if operation == 2 {
				result = opcodes[input1] * opcodes[input2]
			}

			opcodes[op] = result

			if i < len(opcodes) - 1 && opcodes[i + 1] == 99 {
				break
			}

			count = 0
		}
	}

	fmt.Println(opcodes)
	fmt.Println("opcodes[0]: ", opcodes[0])
}
