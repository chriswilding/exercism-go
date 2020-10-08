package allergies

const (
	eggs uint = 1 << iota
	peanuts
	shellfish
	strawberries
	tomatoes
	chocolate
	pollen
	cats
)

var substanceToEnum = map[string]uint{
	"eggs":         eggs,
	"peanuts":      peanuts,
	"shellfish":    shellfish,
	"strawberries": strawberries,
	"tomatoes":     tomatoes,
	"chocolate":    chocolate,
	"pollen":       pollen,
	"cats":         cats,
}

var substances = [...]string{"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies(score uint) []string {
	var allergies []string
	for _, substance := range substances {
		if AllergicTo(score, substance) {
			allergies = append(allergies, substance)
		}
	}
	return allergies
}

func AllergicTo(score uint, substance string) bool {
	enum, _ := substanceToEnum[substance]
	return score&enum > 0
}
