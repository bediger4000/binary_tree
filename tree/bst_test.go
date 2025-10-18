package tree

import "testing"

func TestBstProperty(t *testing.T) {
	type args struct {
		root *NumericNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "balanced search tree",
			args: args{root: CreateNumeric([]string{"4", "2", "6", "1", "3", "5", "7"})},
			want: true,
		},
		{
			name: "not balanced search tree",
			args: args{root: &NumericNode{Data: 4,
				Left: &NumericNode{Data: 2,
					Left:  &NumericNode{Data: 1},
					Right: &NumericNode{Data: 3},
				},
				Right: &NumericNode{Data: 6,
					Left:  &NumericNode{Data: 7},
					Right: &NumericNode{Data: 5},
				},
			},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BstProperty(tt.args.root); got != tt.want {
				t.Errorf("BstProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}
