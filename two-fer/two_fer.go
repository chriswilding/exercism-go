// Package twofer implements the twofer exercism from exercism.io
package twofer

import "fmt"

// ShareWith returns a string describing sharing something. If name is empty name is equivalent to "you"
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
