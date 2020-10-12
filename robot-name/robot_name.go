package robotname

import (
	"errors"
	"fmt"
)

var a rune = 'A'
var b rune = 'A'
var n int = -1

type Robot struct {
	name string
}

var RobotNamesExhausted = errors.New("robot names exhausted")

func next() (string, error) {
	if a == 'Z' && b == 'Z' && n == 999 {
		return "", RobotNamesExhausted
	}
	if b == 'Z' && n == 999 {
		a++
		b = 'A'
		n = -1
	}
	if n == 999 {
		b++
		n = -1
	}
	n++
	return fmt.Sprintf("%c%c%03d", a, b, n), nil
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	name, error := next()
	r.name = name
	return r.name, error

}

func (r *Robot) Reset() {
	r.name = ""
}
