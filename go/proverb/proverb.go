// Package proverb is an exercise of exercism.io
package proverb

import "fmt"

// Proverb should return a proverb using a list of words.
func Proverb(rhyme []string) []string {
	proverb := []string{}
	for i := 0; i < len(rhyme)-1; i++ {
		proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}
	if len(rhyme) > 0 {
		proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	}
	return proverb
}
