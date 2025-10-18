package tree

import "testing"

func TestEquals(t *testing.T) {
	var nilNumeric *NumericNode
	var nilString *StringNode
	var randomTree *NumericNode = RandomValueTree(1000, 20, true)
	var stringTree1, stringTree2, stringTree3, stringTree4, stringTree5 *StringNode
	var stringTree6, stringTree7, stringTree8a, stringTree8b *StringNode
	var stringTree9a, stringTree9b, stringTree10a, stringTree10b *StringNode
	var stringTree11a, stringTree11b, stringTree12, stringTree13, stringTree14 *StringNode
	stringTree1, _ = CreateFromString("(a(b(bl)(br))(c(cl)(cr)))")
	stringTree2, _ = CreateFromString("(a(b(bl)(br))(c()(cr)))")
	stringTree3, _ = CreateFromString("(a(b(bl)(br))(c(cl)()))")
	stringTree4, _ = CreateFromString("(a(b(bl)())(c(cl)(cr)))")
	stringTree5, _ = CreateFromString("(a(b()(br))(c(cl)(cr)))")
	stringTree6, _ = CreateFromString("(a(b(bl)(br))(c(cx)(cr)))")
	stringTree7, _ = CreateFromString("(a(b(bl)(br))(c(cl)(cx)))")
	stringTree8a, _ = CreateFromString("(a()())")
	stringTree8b, _ = CreateFromString("(a()())")
	stringTree9a, _ = CreateFromString("(a(b)())")
	stringTree9b, _ = CreateFromString("(a(b)())")
	stringTree10a, _ = CreateFromString("(a()(b))")
	stringTree10b, _ = CreateFromString("(a()(b))")
	stringTree11a, _ = CreateFromString("(a()(b))")
	stringTree11b, _ = CreateFromString("(a(b)())")
	stringTree12, _ = CreateFromString("(1()())")
	stringTree13, _ = CreateFromString("(1()())")
	stringTree14, _ = CreateFromString("(a()())")
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
			args: args{t1: stringTree8a, t2: stringTree8b},
			want: true,
		},
		{
			name: "two element identical string trees",
			args: args{t1: stringTree9a, t2: stringTree9b},
			want: true,
		},
		{
			name: "two element identical string trees",
			args: args{t1: stringTree10a, t2: stringTree10b},
			want: true,
		},
		{
			name: "two element different string trees",
			args: args{t1: stringTree11a, t2: stringTree11b},
			want: false,
		},
		{
			name: "one element different type trees A",
			args: args{t1: stringTree12, t2: CreateNumeric([]string{"1"})},
			want: false,
		},
		{
			name: "one element different type trees B",
			args: args{t1: CreateNumeric([]string{"1"}), t2: stringTree13},
			want: false,
		},
		{
			name: "one string tree, one nil tree",
			args: args{t1: stringTree14, t2: nilString},
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
