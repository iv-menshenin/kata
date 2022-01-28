package pointsSymmetry

import "testing"

func Test_isSymmetric(t *testing.T) {
	type args struct {
		pointCloud []coords
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{pointCloud: []coords{}},
			want: true,
		},
		{
			name: "only_on_symmetry_axis",
			args: args{pointCloud: []coords{{x: 12, y: 5}, {x: 12, y: 8}}},
			want: true,
		},
		{
			name: "only_two_points_on_the_same_y_axis",
			args: args{pointCloud: []coords{{x: -13, y: 225}, {x: 312, y: 225}}},
			want: true,
		},
		{
			name: "only_two_points_on_the_same_y_axis_odd",
			args: args{pointCloud: []coords{{x: -312, y: 225}, {x: 13, y: 225}}},
			want: true,
		},
		{
			name: "symmetric_left_side",
			args: args{pointCloud: []coords{{x: -7, y: 1}, {x: -1, y: 2}, {x: 1, y: 1}, {x: -5, y: 2}}},
			want: true,
		},
		{
			name: "symmetric_right_side",
			args: args{pointCloud: []coords{{x: 7, y: 1}, {x: 1, y: 2}, {x: -1, y: 1}, {x: 5, y: 2}}},
			want: true,
		},
		{
			name: "symmetric_left_side_odd",
			args: args{pointCloud: []coords{{x: -2, y: 10}, {x: -1, y: -2}, {x: 0, y: -2}, {x: 1, y: 10}}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.pointCloud); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
