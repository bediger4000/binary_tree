package tree

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	type args struct {
		node  *NumericNode
		value int
	}
	tests := []struct {
		name string
		args args
		want *NumericNode
	}{
		{
			name: "insert single value",
			args: args{node: nil, value: 10},
			want: &NumericNode{Data: 10},
		},
		{
			name: "insert less than",
			args: args{node: &NumericNode{Data: 10}, value: 9},
			want: &NumericNode{Data: 10, Left: &NumericNode{Data: 9}},
		},
		{
			name: "insert greater than",
			args: args{node: &NumericNode{Data: 10}, value: 100},
			want: &NumericNode{Data: 10, Right: &NumericNode{Data: 100}},
		},
		{
			name: "recursive insert right",
			args: args{node: &NumericNode{Data: 10, Right: &NumericNode{Data: 100}}, value: 101},
			want: &NumericNode{Data: 10, Right: &NumericNode{Data: 100, Right: &NumericNode{Data: 101}}},
		},
		{
			name: "don't insert dupe",
			args: args{node: &NumericNode{Data: 10}, value: 10},
			want: &NumericNode{Data: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Insert(tt.args.node, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
