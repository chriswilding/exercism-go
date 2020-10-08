package allergies

var substances = [...]string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func substanceToValue(substance string) uint {
	switch substance {
	case "eggs":
		return 1
	case "peanuts":
		return 2
	case "shellfish":
		return 4
	case "strawberries":
		return 8
	case "tomatoes":
		return 16
	case "chocolate":
		return 32
	case "pollen":
		return 64
	case "cats":
		return 128
	}
	panic("unreachable")
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
	value := substanceToValue(substance)
	return score&value > 0
}
