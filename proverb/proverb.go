package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	output := make([]string, 0, len(rhyme))
	if len(rhyme) == 0 {
		return output
	}
	for i := 1; i < len(rhyme); i++ {
		output = append(output, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
	}
	return append(output, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
}
