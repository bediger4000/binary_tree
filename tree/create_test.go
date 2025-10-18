package tree

import (
	"bytes"
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
			args: args{stringrep: "(a ( b ( d ) (e ( f) (g ))) ( c ) )"},
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
			if gotRoot, err := CreateFromString(tt.args.stringrep); err != nil {
				t.Errorf("CreateFromString() failed parse: %v", err)
			} else {
				if !reflect.DeepEqual(gotRoot, tt.wantRoot) {
					t.Errorf("CreateFromString() = %v, want %v", gotRoot, tt.wantRoot)
				}
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
			args:     args{stringrep: "( 90 ( 12 () (3 ) )(6 (4) (90 ( )( 9 ) ) ) )"},
			wantRoot: &NumericNode{Data: 90, Left: &NumericNode{Data: 12, Right: &NumericNode{Data: 3}}, Right: &NumericNode{Data: 6, Left: &NumericNode{Data: 4}, Right: &NumericNode{Data: 90, Right: &NumericNode{Data: 9}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRoot, err := CreateNumericFromString(tt.args.stringrep); err != nil {
				t.Errorf("CreateNumericFromString: parsing string rep: %v", err)
			} else {
				if !reflect.DeepEqual(gotRoot, tt.wantRoot) {
					t.Errorf("CreateNumericFromString() = %v, want %v", gotRoot, tt.wantRoot)
				}
			}
		})
	}
}

func TestPrintf(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		wantOut string
	}{
		{
			name:    "root node only",
			arg:     "(a)",
			wantOut: "(a)",
		},
		{
			name:    "root node, explicit nil left child",
			arg:     "(a    ()  )",
			wantOut: "(a)",
		},
		{
			name:    "root node, explicit nil children",
			arg:     "(a    ()  ())",
			wantOut: "(a)",
		},
		{
			name:    "left child only",
			arg:     "(a (left))",
			wantOut: "(a (left))",
		},
		{
			name:    "right child only",
			arg:     "(a ()     (right))",
			wantOut: "(a () (right))",
		},
		{
			name:    "both children",
			arg:     "(a ( left )     (right))",
			wantOut: "(a (left) (right))",
		},
		{
			name:    "multiple levels",
			arg:     "(root(left(leftleft)(leftright))(right(rightleft)(rightright)))",
			wantOut: "(root (left (leftleft) (leftright)) (right (rightleft) (rightright)))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root, err := CreateFromString(tt.arg)
			if err != nil {
				t.Errorf("TestPrintf - parse problem: %v", err)
			}
			out := &bytes.Buffer{}
			Printf(out, root)
			var gotOut string
			if gotOut = out.String(); gotOut != tt.wantOut {
				t.Errorf("Printf() = %v, want %v", gotOut, tt.wantOut)
			}
			// recycle gotOut, ensure CreateFromString works on it
			root2, err := CreateFromString(gotOut)
			if err != nil {
				t.Errorf("TestPrintf - re-parse problem: %v", err)
			}
			out2 := &bytes.Buffer{}
			Printf(out2, root2)
			if gotOut2 := out2.String(); gotOut2 != gotOut {
				t.Errorf("Printf() = %v, want %v", gotOut2, gotOut)
			}
		})
	}
}
