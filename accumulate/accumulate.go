package accumulate

func Accumulate(inputs []string, callback func(string) string) []string {
	output := make([]string, len(inputs))
	for i := 0; i < len(inputs); i++ {
		output[i] = callback(inputs[i])
	}
	return output
}
