package tree

import "testing"

func TestEquals(t *testing.T) {
	var nilNumeric *NumericNode
	var nilString *StringNode
	var randomTree *NumericNode = RandomValueTree(1000, 20, true)
	var stringTree1 *StringNode = CreateFromString("(a(b(bl)(br))(c(cl)(cr)))")
	var stringTree2 *StringNode = CreateFromString("(a(b(bl)(br))(c()(cr)))")
	var stringTree3 *StringNode = CreateFromString("(a(b(bl)(br))(c(cl)()))")
	var stringTree4 *StringNode = CreateFromString("(a(b(bl)())(c(cl)(cr)))")
	var stringTree5 *StringNode = CreateFromString("(a(b()(br))(c(cl)(cr)))")
	var stringTree6 *StringNode = CreateFromString("(a(b(bl)(br))(c(cx)(cr)))")
	var stringTree7 *StringNode = CreateFromString("(a(b(bl)(br))(c(cl)(cx)))")
	type args struct {
		t1 Node
		t2 Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "both numeric trees nil",
			args: args{t1: nilNumeric, t2: nilNumeric},
			want: true,
		},
		{
			name: "both string trees nil",
			args: args{t1: nilString, t2: nilString},
			want: true,
		},
		{
			name: "one string tree nil, one numeric tree nil",
			args: args{t1: nilString, t2: nilNumeric},
			want: true,
		},
		{
			name: "zero-count numeric trees",
			args: args{t1: CreateNumeric([]string{}), t2: CreateNumeric([]string{})},
			want: true,
		},
		{
			name: "one element numeric trees, identical",
			args: args{t1: CreateNumeric([]string{"1"}), t2: CreateNumeric([]string{"1"})},
			want: true,
		},
		{
			name: "two element numeric trees, different",
			args: args{t1: CreateNumeric([]string{"1", "0"}), t2: CreateNumeric([]string{"0", "1"})},
			want: false,
		},
		{
			name: "three element numeric trees, identical",
			args: args{t1: CreateNumeric([]string{"1", "0", "2"}), t2: CreateNumeric([]string{"1", "0", "2"})},
			want: true,
		},
		{
			name: "one element identical string trees",
			args: args{t1: CreateFromString("(a()())"), t2: CreateFromString("(a()())")},
			want: true,
		},
		{
			name: "two element identical string trees",
			args: args{t1: CreateFromString("(a(b)())"), t2: CreateFromString("(a(b)())")},
			want: true,
		},
		{
			name: "two element identical string trees",
			args: args{t1: CreateFromString("(a()(b))"), t2: CreateFromString("(a()(b))")},
			want: true,
		},
		{
			name: "two element different string trees",
			args: args{t1: CreateFromString("(a()(b))"), t2: CreateFromString("(a(b)())")},
			want: false,
		},
		{
			name: "one element different type trees A",
			args: args{t1: CreateFromString("(1()())"), t2: CreateNumeric([]string{"1"})},
			want: false,
		},
		{
			name: "one element different type trees B",
			args: args{t1: CreateNumeric([]string{"1"}), t2: CreateFromString("(1()())")},
			want: false,
		},
		{
			name: "one string tree, one nil tree",
			args: args{t1: CreateFromString("(a()())"), t2: nilString},
			want: false,
		},
		{
			name: "random numeric tree, self compare",
			args: args{t1: randomTree, t2: randomTree},
			want: true,
		},
		{
			name: "string tree, self compare",
			args: args{t1: stringTree1, t2: stringTree1},
			want: true,
		},
		{
			name: "string trees A",
			args: args{t1: stringTree1, t2: stringTree2},
			want: false,
		},
		{
			name: "string trees B",
			args: args{t1: stringTree1, t2: stringTree3},
			want: false,
		},
		{
			name: "string trees C",
			args: args{t1: stringTree1, t2: stringTree4},
			want: false,
		},
		{
			name: "string trees D",
			args: args{t1: stringTree1, t2: stringTree5},
			want: false,
		},
		{
			name: "string trees E",
			args: args{t1: stringTree1, t2: stringTree6},
			want: false,
		},
		{
			name: "string trees F",
			args: args{t1: stringTree1, t2: stringTree7},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equals(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
