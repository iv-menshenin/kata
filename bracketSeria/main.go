package bracketSeria

func makeBracketRecursive(n int) []string {
	if n == 0 {
		return nil
	}
	if n == 1 {
		return []string{"()"}
	}
	var m = map[string]struct{}{}
	for i := 1; i < n; i++ {
		for _, l := range makeBracketRecursive(i) {
			for _, r := range makeBracketRecursive(n - i) {
				m[l+r] = struct{}{}
				m[r+l] = struct{}{}
			}
		}
	}
	for _, v := range makeBracketRecursive(n - 1) {
		m["("+v+")"] = struct{}{}
	}
	var result = make([]string, 0)
	for k := range m {
		result = append(result, k)
	}
	return result
}

func makeBracketCycle(n int) []string {
	type parSet struct {
		set    string
		opened int
		closed int
	}
	var generated []parSet
	for pos := 0; pos < n*2; pos++ {
		if len(generated) == 0 {
			generated = make([]parSet, 1)
		}
		for i := range generated {
			var next = make([]parSet, 0, 2)
			if item := generated[i]; item.closed < item.opened {
				item.set += ")"
				item.closed++
				next = append(next, item)
			}
			if item := generated[i]; item.opened < n {
				item.set += "("
				item.opened++
				next = append(next, item)
			}
			generated[i] = next[0]
			generated = append(generated, next[1:]...)
		}
	}
	var result = make([]string, len(generated))
	for i, r := range generated {
		result[i] = r.set
	}
	return result
}
