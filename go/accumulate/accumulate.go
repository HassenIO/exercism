package accumulate

// Accumulate transforms am array of string into another array of string, given the `transformer` function
func Accumulate(input []string, transformer func(string) string) []string {
	output := make([]string, len(input))
	for i := range input {
		output[i] = transformer(input[i])
	}
	return output
}
