package utils

// checkError checks
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
