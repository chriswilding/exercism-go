package diffsquares

func SquareOfSum(input int) int {
	sum := input * (input + 1) / 2
	return sum * sum
}

func SumOfSquares(input int) int {
	return (input * (input + 1) * (2*input + 1)) / 6
}

func Difference(input int) int {
	return SquareOfSum(input) - SumOfSquares(input)
}
