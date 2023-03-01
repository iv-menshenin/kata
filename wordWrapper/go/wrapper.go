package _go

import "strings"

func Wrap(s string, colSz int) []string {
	var wrapped []string
	for s != "" {
		line := s
		if len(line) > colSz {
			if l := strings.LastIndex(line[:colSz+1], " "); l > -1 {
				line = line[:l]
			} else {
				line = line[:colSz]
			}
		}
		wrapped = append(wrapped, line)
		s = strings.TrimSpace(s[len(line):])
	}
	return wrapped
}
