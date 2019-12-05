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

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			memory := make([]int, len(opcodes))
			copy(memory, opcodes)
			memory[1] = noun
			memory[2] = verb

			runComputer(&memory)

			if memory[0] == 19690720 {
				fmt.Println("noun: ", noun)
				fmt.Println("verb: ", verb)
				fmt.Println("100 * noun + verb: ", 100 * noun + verb)
			}
		}
	}
}

func runComputer(instructions *[]int) {
	var count int
	instr := *instructions

	for i, op := range instr {
		count++
		if count == 4 {
			operation := instr[i - 3]
			input1 := instr[i - 2]
			input2 := instr[i - 1]

			var result int

			if operation == 1 {
				result = instr[input1] + instr[input2]
			} else if operation == 2 {
				result = instr[input1] * instr[input2]
			}

			instr[op] = result

			if i < len(instr) - 1 && instr[i + 1] == 99 {
				break
			}

			count = 0
		}
	}
}
