package anagram

// Identify if the two words obtained are anagrams

func isAnagramMapped(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return false
	}
	var m = make(map[int32]int)
	for _, r := range a {
		m[r]++
	}
	for _, r := range b {
		if m[r]--; m[r] == 0 {
			delete(m, r)
		}
	}
	return len(m) == 0
}
