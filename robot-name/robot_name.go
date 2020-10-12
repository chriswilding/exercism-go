package robotname

import "math/rand"

var names map[string]struct{}

type Robot struct {
	name string
}

func generate() string {
	var runes [5]rune
	runes[0] = rune(65 + rand.Intn(26))
	runes[1] = rune(65 + rand.Intn(26))
	runes[2] = rune(48 + rand.Intn(9))
	runes[3] = rune(48 + rand.Intn(9))
	runes[4] = rune(48 + rand.Intn(9))
	return string(runes[:])
}

func newName() string {
	if names == nil {
		names = make(map[string]struct{})
	}
	for {
		name := generate()
		_, seen := names[name]
		if !seen {
			names[name] = struct{}{}
			return name
		}
	}
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = newName()
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
