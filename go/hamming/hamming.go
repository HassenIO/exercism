package hamming

import "errors"

// Distance calculates Hamming Distance between two strings.
func Distance(a, b string) (int, error) {
	if a == b {
		return 0, nil
	}
	if len(a) != len(b) {
		return -1, errors.New("strings must have the same length")
	}
	var hammingDist int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingDist++
		}
	}
	return hammingDist, nil
}
