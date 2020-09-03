package strain

type Ints []int

func (is Ints) Keep(predicate func(int) bool) Ints {
	var output Ints
	for _, i := range is {
		if ok := predicate(i); ok {
			output = append(output, i)
		}

	}
	return output
}

func (is Ints) Discard(predicate func(int) bool) Ints {
	return is.Keep(func(i int) bool {
		return !predicate(i)
	})
}

type Lists [][]int

func (ls Lists) Keep(predicate func([]int) bool) Lists {
	var output Lists
	for _, i := range ls {
		if ok := predicate(i); ok {
			output = append(output, i)
		}

	}
	return output
}

type Strings []string

func (ss Strings) Keep(predicate func(string) bool) Strings {
	var output Strings
	for _, i := range ss {
		if ok := predicate(i); ok {
			output = append(output, i)
		}

	}
	return output
}
