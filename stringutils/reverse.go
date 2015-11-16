package stringutil

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(inputString string) string {
	r := []rune(inputString)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}