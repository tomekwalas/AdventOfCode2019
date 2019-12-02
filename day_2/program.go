package day_2

const (
	AddProgram      = 1
	MultiplyProgram = 2
	EndOfProgram    = 99
)

func Parse(programInputs []int) []int {
	i := 0
	for {
		typeOfProgram := programInputs[i]

		if typeOfProgram == EndOfProgram {
			return programInputs
		}
		num1 := programInputs[programInputs[i+1]]
		num2 := programInputs[programInputs[i+2]]
		resultPosition := programInputs[i+3]

		if typeOfProgram == AddProgram {
			programInputs[resultPosition] = num1 + num2
		} else if typeOfProgram == MultiplyProgram {
			programInputs[resultPosition] = num1 * num2
		}
		i += 4
	}
}
