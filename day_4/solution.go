package day_4

import "fmt"

func GetSolution() {
	correctPasswords := 0
	for i := 153517; i <= 630395; i++ {
		if ValidatePartTwo(i) {
			correctPasswords += 1
		}
	}

	fmt.Println(correctPasswords)
}
