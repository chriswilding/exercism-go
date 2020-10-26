package house

import "strings"

var things = []string{
	"the house",
	"the malt",
	"the rat",
	"the cat",
	"the dog",
	"the cow with the crumpled horn",
	"the maiden all forlorn",
	"the man all tattered and torn",
	"the priest all shaven and shorn",
	"the rooster that crowed in the morn",
	"the farmer sowing his corn",
	"the horse and the hound and the horn",
}

var verbs = []string{
	"that Jack built.",
	"that lay in",
	"that ate",
	"that killed",
	"that worried",
	"that tossed",
	"that milked",
	"that kissed",
	"that married",
	"that woke",
	"that kept",
	"that belonged to",
}

func Verse(n int) string {
	var b strings.Builder
	b.WriteString("This is ")
	for i := n - 1; i > -1; i-- {
		b.WriteString(things[i])
		if i == 0 {
			b.WriteRune(' ')
		} else {
			b.WriteRune('\n')
		}
		b.WriteString(verbs[i])
		if i != 0 {
			b.WriteRune(' ')
		}
	}
	return b.String()
}

func Song() string {
	var b strings.Builder
	for i := 1; i < 13; i++ {
		b.WriteString(Verse(i))
		b.WriteString("\n\n")
	}
	return strings.TrimSpace(b.String())
}
