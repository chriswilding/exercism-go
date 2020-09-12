package reverse

func Reverse(input string) string {
	runes := []rune(input)
	l := len(runes)
	for i := 0; i < l/2; i++ {
		runes[i], runes[l-1-i] = runes[l-1-i], runes[i]
	}
	return string(runes)
}
