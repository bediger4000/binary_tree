package tree

import (
	"reflect"
	"testing"
)

func TestCreateNumeric(t *testing.T) {
	type args struct {
		numberRepr []string
	}
	tests := []struct {
		name     string
		args     args
		wantRoot *NumericNode
	}{
		{
			name:     "one element",
			args:     args{[]string{"42"}},
			wantRoot: &NumericNode{Data: 42},
		},
		{
			name:     "three element",
			args:     args{[]string{"1", "0", "2"}},
			wantRoot: &NumericNode{Data: 1, Left: &NumericNode{Data: 0}, Right: &NumericNode{Data: 2}},
		},
		{
			name:     "three element left",
			args:     args{[]string{"2", "1", "0"}},
			wantRoot: &NumericNode{Data: 2, Left: &NumericNode{Data: 1, Left: &NumericNode{Data: 0}}},
		},
		{
			name:     "three element right",
			args:     args{[]string{"0", "10", "100"}},
			wantRoot: &NumericNode{Data: 0, Right: &NumericNode{Data: 10, Right: &NumericNode{Data: 100}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRoot := CreateNumeric(tt.args.numberRepr); !reflect.DeepEqual(gotRoot, tt.wantRoot) {
				t.Errorf("CreateNumeric() = %v, want %v", gotRoot, tt.wantRoot)
			}
		})
	}
}

func TestCreateFromString(t *testing.T) {
	type args struct {
		stringrep string
	}
	tests := []struct {
		name     string
		args     args
		wantRoot *StringNode
	}{
		{
			name:     "single string tree 1",
			args:     args{stringrep: "(a)"},
			wantRoot: &StringNode{Data: "a"},
		},
		{
			name:     "single string tree 2",
			args:     args{stringrep: "(a()())"},
			wantRoot: &StringNode{Data: "a"},
		},
		{
			name:     "left child string tree",
			args:     args{stringrep: "(a(b)())"},
			wantRoot: &StringNode{Data: "a", Left: &StringNode{Data: "b"}},
		},
		{
			name:     "right child string tree",
			args:     args{stringrep: "(a()(c))"},
			wantRoot: &StringNode{Data: "a", Right: &StringNode{Data: "c"}},
		},
		{
			name: "big string tree with whitespace",
			args: args{stringrep: " (a ( b ( d ) (e ( f) (g ))) ( c ) ) "},
			wantRoot: &StringNode{Data: "a",
				Left: &StringNode{Data: "b",
					Left: &StringNode{Data: "d"},
					Right: &StringNode{Data: "e",
						Left:  &StringNode{Data: "f"},
						Right: &StringNode{Data: "g"},
					},
				},
				Right: &StringNode{Data: "c"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRoot := CreateFromString(tt.args.stringrep); !reflect.DeepEqual(gotRoot, tt.wantRoot) {
				t.Errorf("CreateFromString() = %v, want %v", gotRoot, tt.wantRoot)
			}
		})
	}
}

func TestCreateNumericFromString(t *testing.T) {
	type args struct {
		stringrep string
	}
	tests := []struct {
		name     string
		args     args
		wantRoot *NumericNode
	}{
		{
			name:     "single item numeric tree 1",
			args:     args{stringrep: "(42)"},
			wantRoot: &NumericNode{Data: 42},
		},
		{
			name:     "single item numeric tree 2",
			args:     args{stringrep: "(42()())"},
			wantRoot: &NumericNode{Data: 42},
		},
		{
			name:     "left child numeric tree",
			args:     args{stringrep: "(42(90)())"},
			wantRoot: &NumericNode{Data: 42, Left: &NumericNode{Data: 90}},
		},
		{
			name:     "right child numeric tree",
			args:     args{stringrep: "(90()(9))"},
			wantRoot: &NumericNode{Data: 90, Right: &NumericNode{Data: 9}},
		},
		{
			name:     "arbitrary numeric tree",
			args:     args{stringrep: "(90(12()(3))(6(4)(90()(9))))"},
			wantRoot: &NumericNode{Data: 90, Left: &NumericNode{Data: 12, Right: &NumericNode{Data: 3}}, Right: &NumericNode{Data: 6, Left: &NumericNode{Data: 4}, Right: &NumericNode{Data: 90, Right: &NumericNode{Data: 9}}}},
		},
		{
			name:     "arbitrary numeric tree with whitespace",
			args:     args{stringrep: " ( 90 ( 12 () (3 ) )(6 (4) (90 ( )( 9 ) ) ) )"},
			wantRoot: &NumericNode{Data: 90, Left: &NumericNode{Data: 12, Right: &NumericNode{Data: 3}}, Right: &NumericNode{Data: 6, Left: &NumericNode{Data: 4}, Right: &NumericNode{Data: 90, Right: &NumericNode{Data: 9}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRoot := CreateNumericFromString(tt.args.stringrep); !reflect.DeepEqual(gotRoot, tt.wantRoot) {
				t.Errorf("CreateNumericFromString() = %v, want %v", gotRoot, tt.wantRoot)
			}
		})
	}
}
