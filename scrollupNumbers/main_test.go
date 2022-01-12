package scrollupNumbers

import "testing"

type scrollupNumbersFn func([]int) string

func testScrollupNumbers(t *testing.T, fn scrollupNumbersFn) {
	tests := []struct {
		name string
		arg  []int
		want string
	}{
		{
			name: "nil",
			arg:  nil,
			want: "",
		},
		{
			name: "once",
			arg:  []int{4},
			want: "4",
		},
		{
			name: "one_tuple",
			arg:  []int{2, 1, 0, 4, 3},
			want: "0-4",
		},
		{
			name: "mixed",
			arg:  []int{12, 0, 2, 11, 3, 4, 5, 6, 13},
			want: "0,2-6,11-13",
		},
		{
			name: "mixed_with_appendix",
			arg:  []int{12, 2, 11, 3, 44, 4, 5, 6, 13},
			want: "2-6,11-13,44",
		},
	}
	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			if got := fn(test.arg); got != test.want {
				t.Errorf("got: '%s', want: '%s'", got, test.want)
			}
		})
	}
}

func Test_scrollupNumbersSerial(t *testing.T) {
	testScrollupNumbers(t, scrollupNumbersSerial)
}

func Test_scrollNumbersBoilerPrint(t *testing.T) {
	testScrollupNumbers(t, scrollNumbersBoilerPrint)
}

func Test_scrollNumbersSimpled(t *testing.T) {
	testScrollupNumbers(t, scrollNumbersSimpled)
}
