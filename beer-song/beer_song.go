package beer

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func action(verse int) string {
	switch verse {
	case 1:
		return "Take it down and pass it around"
	case 0:
		return "Go to the store and buy some more"
	default:
		return "Take one down and pass it around"
	}
}

func container(verse int) string {
	if verse == 1 {
		return "bottle"
	}
	return "bottles"
}

func quantity(verse int) string {
	if verse == 0 {
		return "no more"
	}
	return strconv.Itoa(verse)
}

func next(verse int) int {
	if verse == 0 {
		return 99
	}
	return verse - 1
}

func toString(verse int) string {
	return fmt.Sprintf("%s %s", quantity(verse), container(verse))
}

func capitalise(s string) string {
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func Verse(verse int) (string, error) {
	if verse < 0 || verse > 100 {
		return "", errors.New("invalid verse")
	}
	bns := toString(verse)
	var b strings.Builder
	b.WriteString(capitalise(bns))
	b.WriteString(" of beer on the wall, ")
	b.WriteString(bns)
	b.WriteString(" of beer.\n")
	b.WriteString(action(verse))
	b.WriteString(", ")
	nbns := toString(next(verse))
	b.WriteString(nbns)
	b.WriteString(" of beer on the wall.\n")
	return b.String(), nil
}

func Verses(from, to int) (string, error) {
	if to > from {
		return "", errors.New("to must be equal to of greater than from")
	}
	var b strings.Builder
	for i := from; i >= to; i-- {
		v, e := Verse(i)
		if e != nil {
			return "", e
		}
		b.WriteString(v)
		b.WriteRune('\n')
	}
	return b.String(), nil
}

func Song() string {
	s, _ := Verses(99, 0)
	return s
}
