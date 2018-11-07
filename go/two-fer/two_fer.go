// Package twofer is an exercise on exercism
package twofer

import "fmt"

// ShareWith returns a string in the form of "One for X, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
