package factor

func Generate(i int) (factors []int) {
	for candidate := 2; i > 1; candidate++ {
		for ; i%candidate == 0; i /= candidate {
			factors = append(factors, candidate)
		}
	}
	return
}
