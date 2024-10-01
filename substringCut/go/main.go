package sub

// Написать функцию, которая на входе принимает строку и число K.
// Функция возвращает длину максимальной подстроки, которая содержит не более K различных символов из строки.

func getSubString(s string, k int) int {
	var tail, subMaxLen int
	var cnt = make(map[rune]int)
	for head, r := range s {
		cnt[r]++
		for len(cnt) > k {
			tailRune := rune(s[tail])
			cnt[tailRune]--
			if cnt[tailRune] == 0 {
				delete(cnt, tailRune)
			}
			tail++
		}
		if subMaxLen < head-tail+1 {
			subMaxLen = head - tail + 1
		}
	}
	return subMaxLen
}
