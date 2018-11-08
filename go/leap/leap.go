// Package leap is related to "Leap exercise" on exercism.io.
package leap

// IsLeapYear checks if a given year is a leap year or not.
func IsLeapYear(year int) bool {
	return year%400 == 0 || year%100 != 0 && year%4 == 0
}
