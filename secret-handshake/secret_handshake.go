package secret

const (
	wink uint = 1 << iota
	doubleBlink
	closeYourEyes
	jump
	reverse
)

func Handshake(code uint) []string {
	output := make([]string, 0, 5)

	if code&wink > 0 {
		output = append(output, "wink")
	}
	if code&doubleBlink > 0 {
		output = append(output, "double blink")
	}
	if code&closeYourEyes > 0 {
		output = append(output, "close your eyes")
	}
	if code&jump > 0 {
		output = append(output, "jump")
	}
	if code&reverse > 0 {
		for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
			output[i], output[j] = output[j], output[i]
		}
	}

	return output
}
