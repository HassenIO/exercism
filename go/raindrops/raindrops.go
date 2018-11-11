// Package raindrops is an exercise on exercism.io
package raindrops

import "strconv"

// Convert a number to string following exercism.io's Raindrops exercise
func Convert(number int) string {
	text := ""
	if number%3 == 0 {
		text += "Pling"
	}
	if number%5 == 0 {
		text += "Plang"
	}
	if number%7 == 0 {
		text += "Plong"
	}
	if text == "" {
		text = strconv.Itoa(number)
	}
	return text
}
