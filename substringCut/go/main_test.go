package sub

import "testing"

func Test_getSubString(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			args: args{
				s: "",
				k: 2,
			},
			want: 0,
		},
		{
			name: "all",
			args: args{
				s: "aaa",
				k: 1,
			},
			want: 3,
		},
		{
			name: "jjhhdsh",
			args: args{
				s: "askdjjhhdshtc",
				k: 4,
			},
			want: 8,
		},
		{
			name: "asa",
			args: args{
				s: "asafghjkl",
				k: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSubString(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("getSubString() = %v, want %v", got, tt.want)
			}
		})
	}
}
