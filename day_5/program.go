package day_5

import "fmt"

const (
	Address         = 0
	AddProgram      = 1
	MultiplyProgram = 2
	Input           = 3
	Output          = 4
	JumpIfTrue      = 5
	JumpIfFalse     = 6
	LessThan        = 7
	Equals          = 8
	EndOfProgram    = 99
)

func parseInstruction(instruction int) []int {
	isInstruction := true
	cutter := 100
	var instructions []int
	for instruction > 0 {
		instructions = append(instructions, instruction%cutter)
		instruction /= cutter
		if isInstruction {
			isInstruction = false
			cutter = 10
		}
	}

	if len(instructions) < 4 {
		for i := len(instructions); i < 4; i++ {
			instructions = append(instructions, 0)
		}
	}

	return instructions
}

func RunPrograms(programInputs []int, userInput int) {
	i := 0
	for {
		opcode := programInputs[i]
		var instructions []int
		if opcode > 99 {
			instructions = parseInstruction(opcode)
			opcode = instructions[0]
		}
		switch opcode {
		case EndOfProgram:
			return
		case Input:
			position := programInputs[i+1]
			programInputs[position] = userInput
			i += 2
			break
		case Output:
			position := programInputs[i+1]
			fmt.Println(programInputs[position])
			i += 2
			break
		case AddProgram:
			input1 := programInputs[i+1]
			input2 := programInputs[i+2]
			output := programInputs[i+3]

			if len(instructions) > 0 {
				if instructions[1] == Address {
					input1 = programInputs[input1]
				}
				if instructions[2] == Address {
					input2 = programInputs[input2]
				}
			} else {
				input1 = programInputs[input1]
				input2 = programInputs[input2]
			}

			programInputs[output] = input1 + input2

			i += 4
			break
		case MultiplyProgram:
			input1 := programInputs[i+1]
			input2 := programInputs[i+2]
			output := programInputs[i+3]

			if len(instructions) > 0 {
				if instructions[1] == Address {
					input1 = programInputs[input1]
				}
				if instructions[2] == Address {
					input2 = programInputs[input2]
				}
			} else {
				input1 = programInputs[input1]
				input2 = programInputs[input2]
			}

			programInputs[output] = input1 * input2

			i += 4
			break
		case JumpIfTrue:
			input1 := programInputs[i+1]
			input2 := programInputs[i+2]
			if len(instructions) > 0 {
				if instructions[1] == Address {
					input1 = programInputs[input1]
				}
				if instructions[2] == Address {
					input2 = programInputs[input2]
				}
			} else {
				input1 = programInputs[input1]
				input2 = programInputs[input2]
			}
			if input1 > 0 {
				i = input2
			} else {
				i = i + 3
			}
			break
		case JumpIfFalse:
			input1 := programInputs[i+1]
			input2 := programInputs[i+2]
			if len(instructions) > 0 {
				if instructions[1] == Address {
					input1 = programInputs[input1]
				}
				if instructions[2] == Address {
					input2 = programInputs[input2]
				}
			} else {
				input1 = programInputs[input1]
				input2 = programInputs[input2]
			}
			if input1 == 0 {
				i = input2
			} else {
				i = i + 3
			}
			break
		case LessThan:
			input1 := programInputs[i+1]
			input2 := programInputs[i+2]
			output := programInputs[i+3]

			if len(instructions) > 0 {
				if instructions[1] == Address {
					input1 = programInputs[input1]
				}
				if instructions[2] == Address {
					input2 = programInputs[input2]
				}
			} else {
				input1 = programInputs[input1]
				input2 = programInputs[input2]
			}

			if input1 < input2 {
				programInputs[output] = 1
			} else {
				programInputs[output] = 0
			}

			i += 4
			break
		case Equals:
			input1 := programInputs[i+1]
			input2 := programInputs[i+2]
			output := programInputs[i+3]

			if len(instructions) > 0 {
				if instructions[1] == Address {
					input1 = programInputs[input1]
				}
				if instructions[2] == Address {
					input2 = programInputs[input2]
				}
			} else {
				input1 = programInputs[input1]
				input2 = programInputs[input2]
			}

			if input1 == input2 {
				programInputs[output] = 1
			} else {
				programInputs[output] = 0
			}

			i += 4
			break
		}

	}
}
