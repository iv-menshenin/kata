package anagram

import (
	"strings"
	"testing"
)

type testedFunc func(string, string) bool

func testSomeAnagramChecker(t *testing.T, fn testedFunc) {
	var tests = []struct {
		name string
		arg1 string
		arg2 string
		want bool
	}{
		{
			name: "empty",
			want: false,
		},
		{
			name: "one_char",
			arg1: "a",
			arg2: "a",
			want: true,
		},
		{
			name: "anagram_1",
			arg1: "foo",
			arg2: "ofo",
			want: true,
		},
		{
			name: "anagram_2",
			arg1: "foobaranagram",
			arg2: "anarobramfoag",
			want: true,
		},
		{
			name: "anagram_cyr",
			arg1: "инфографика_ABC",
			arg2: "ииоAгBнраCфф_ка",
			want: true,
		},
		{
			name: "not_anagram_wrong_len",
			arg1: "bar",
			arg2: "baar",
			want: false,
		},
		{
			name: "not_anagram_correct_len",
			arg1: "bar",
			arg2: "baa",
			want: false,
		},
	}
	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			got := fn(test.arg1, test.arg2)
			if test.want != got {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}

func benchSomeAnagramChecker(b *testing.B, fn testedFunc) {
	var tests = []struct {
		arg1 string
		arg2 string
	}{
		{
			arg1: strings.Repeat("foo", 100),
			arg2: strings.Repeat("ofo", 100),
		},
		{
			arg1: strings.Repeat("bar-anagram2", 100),
			arg2: strings.Repeat("baarna-amgr1", 100),
		},
		{
			arg1: strings.Repeat("инфографика_ABCинфографика_ABCинфографика_ABC", 100),
			arg2: strings.Repeat("ииоAгBнраCфф_каииоAгBнраCфф_каииоAгBнраCфф_ка", 100),
		},
	}
	b.ResetTimer()
	for nn := 0; nn < b.N; nn += len(tests) {
		for _, test := range tests {
			_ = fn(test.arg1, test.arg2)
		}
	}
}
