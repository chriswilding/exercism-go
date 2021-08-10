package stringset

import (
	"bytes"
	"fmt"
	"reflect"
)

type Set map[string]struct{}

func New() Set {
	return make(Set)
}

func NewFromSlice(items []string) Set {
	set := make(Set, len(items))

	for _, item := range items {
		set.Add(item)
	}

	return set
}

func (s Set) String() string {
	var i int
	var b bytes.Buffer

	b.WriteRune('{')

	for item := range s {
		if i == len(s)-2 {
			fmt.Fprintf(&b, "\"%s\", ", item)
		} else {
			fmt.Fprintf(&b, "\"%s\"", item)
		}
		i++
	}

	b.WriteRune('}')

	return b.String()
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(item string) bool {
	_, ok := s[item]
	return ok
}

func Subset(s1, s2 Set) bool {
	if len(s1) > len(s2) {
		return false
	}

	for item := range s1 {
		if s2.Has(item) {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	for item := range s1 {
		if s2.Has(item) {
			return false
		}
	}

	return true
}

func Equal(s1, s2 Set) bool {
	return reflect.DeepEqual(s1, s2)
}

func (s Set) Add(item string) {
	s[item] = struct{}{}
}

func Intersection(s1, s2 Set) Set {
	newSet := make(Set)

	for item := range s1 {
		if s2.Has(item) {
			newSet.Add(item)
		}
	}

	return newSet
}

func Difference(s1, s2 Set) Set {
	newSet := make(Set)

	for item := range s1 {
		if !s2.Has(item) {
			newSet.Add(item)
		}
	}

	return newSet
}

func Union(s1, s2 Set) Set {
	newSet := make(Set)

	for item := range s1 {
		newSet.Add(item)
	}

	for item := range s2 {
		newSet.Add(item)
	}

	return newSet
}
