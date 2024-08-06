package utils

import "strconv"

func StringToInt(input string) (int, error) {
	intVar, err := strconv.Atoi(input)

	return intVar, err
}
