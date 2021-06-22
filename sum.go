package calculator

var logMessage = "[LOG]"

var Version = "1.0"

func internalsum(number int) int {
	return number - 1
}

func Sum(N1, N2 int) (N3 int) {
	N3 = N1 + N2
	return N3
}
