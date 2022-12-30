package factor

func Generate(i int) []int {
	var factors []int
	var candidate int = 2
	for i > 1 {
		for ; i%candidate == 0; i /= candidate {
			factors = append(factors, candidate)
		}
		candidate++
	}
	return factors
}
