package listops

type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

type IntList []int

func (l IntList) Foldl(callback binFunc, initial int) int {
	for i := 0; i < l.Length(); i++ {
		value := l[i]
		initial = callback(initial, value)
	}
	return initial
}

func (l IntList) Foldr(callback binFunc, initial int) int {
	for i := l.Length() - 1; i >= 0; i-- {
		value := l[i]
		initial = callback(value, initial)
	}
	return initial
}

func (l IntList) Filter(callback predFunc) IntList {
	if l.Length() == 0 {
		return l
	}
	var output IntList
	for i := 0; i < l.Length(); i++ {
		item := l[i]
		if result := callback(item); result {
			output = append(output, item)
		}
	}
	return output
}

func (l IntList) Length() int {
	return len(l)
}

func (l IntList) Map(callback unaryFunc) IntList {
	output := make(IntList, l.Length())
	for i := 0; i < l.Length(); i++ {
		output[i] = callback(l[i])
	}
	return output
}

func (l IntList) Reverse() IntList {
	output := make(IntList, l.Length())
	for i := 0; i < l.Length(); i++ {
		output[l.Length()-i-1] = l[i]
	}
	return output
}

func (l IntList) Append(o IntList) IntList {
	return append(l, o...)
}

func (l IntList) Concat(lists []IntList) IntList {
	for _, list := range lists {
		l = append(l, list...)
	}
	return l
}
