package utils

import "strconv"

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func StringToInt(str string) int {
	intVar, err := strconv.Atoi(str)
	HandleError(err)
	return intVar
}
