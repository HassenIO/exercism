// Package strand is an exercise on exercism.io
package strand

var conversion = map[rune]byte{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA converts a DNA string to its RNA complement, given a set of rules.
func ToRNA(dna string) string {
	rna := []byte(dna)
	for pos, char := range dna {
		rna[pos] = conversion[char]
	}
	return string(rna)
}
