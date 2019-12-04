package day_4

func reverse(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func toDigits(code int) []int {
	var digits []int
	for code > 0 {
		digits = append(digits, code%10)
		code = code / 10
	}

	return reverse(digits)
}

func Validate(code int) bool {
	digits := toDigits(code)
	hasDoubleDigit := false
	for i, digit := range digits {
		for j, innerDigit := range digits {
			if j < i {
				continue
			}
			if digit == innerDigit && i != j {

				hasDoubleDigit = true
			}
			if digit > innerDigit {
				return false
			}
		}
	}

	return hasDoubleDigit
}

func ValidatePartTwo(code int) bool {
	digits := toDigits(code)
	multipleDigits := make(map[int]int)

	for i, digit := range digits {
		for j, innerDigit := range digits {
			if j < i {
				continue
			}
			if digit == innerDigit && i != j {
				multipleDigits[digit] += 2
			}
			if digit > innerDigit {
				return false
			}
		}
	}

	for _, v := range multipleDigits {
		if v == 2 {
			return true
		}
	}
	return false
}
