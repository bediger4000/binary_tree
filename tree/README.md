# package tree notes

`package tree` is for the most part,
support functions that are common to at a few of the binary tree puzzles
or interview questions.

## Data Structures

In this package a binary tree consists of instances of
either type `Node`, or type `StringNode`.
The root of the tree is merely a pointer to an instance of one of those types,
not a instance of a special "tree" struct.

Connection to left and right sub-trees is through pointers.
Making both root and connections to sub-trees pointers means
that recursive functions can act on the tree,
pruning it or rearranging nodes.

### struct Node

```go
type Node struct {
    Data  int
    Left  *Node
    Right *Node
}
```

Numeric data, allowing the tree to be ordered.

Fields are exported because a lot of the problem require
either getters/setters (ugh) or direct access to instance internals.

### struct StringNode

```go
type StringNode struct {
    Data  string
    Left  *StringNode 
    Right *StringNode
}
```

### interface drawable

```go
type drawable interface {
    leftChild()  drawable
    rightChild() drawable
    isNil()      bool
}
```

This is really the only technically interesting data type.
With methods that satisfy `interface drawable` (which exist in [types.go](types.go)),
the same function (`func DrawPrefixed(out io.Writer, node drawable, prefix string)`)
can generate GraphViz output for trees of both `*Node` and `*StringNode`
types.

The traversal functions also use interface drawable to be somewhat more generic.

Dismayingly, `leftChild()` and `rightChild()` are both getters.

## Input/Output

### Creating a binary search tree from a list of numeric string representations

```go
func CreateNumeric(numberRepr []string) (root *Node)
```

Creates a binary search tree of type `*Node` from a slice of strings.
The strings have to be representations of integers.

### Creating a tree from a list of strings

This one's a little ugly.

```go
func CreateFromString(stringrep string) (root *StringNode)
```

This function will create a tree of type `*StringNode`.
The format of `stringrep` is something like a Lisp expression:
An input string `(root(left(gl)(gr))(right(gL)(gR))`
results in a tree like this:

          root
         /    \
      left     right
      /  \     /  \
    gl   gr   gL   gR

No error recovery, very brittle parsing, can't tolerate whitespace,
you have to indicate nil left children with `()` if you have non-nil right children.
`()` is option for nil right children, due to mediocre parsing.

### GraphViz drawing of a tree

```go
func Draw(root drawable)
func DrawPrefixed(out io.Writer, root drawable, prefix string)
```

Both of these accept the root of a binary tree,
comprised of either `*Node` or `*StringNode` elements,
and output [GraphViz](https://graphviz.org) directives
that can be used as input to `dot` to get a visualization
of the tree.

Programs `drawtree`, `invert`, `minimal_ht_tree`, `prune_tree`,
`reconstruct` all create GraphViz output,
always on stdout because I find redirection far more convenient
than an output-to-file command line flag.

`DrawPrefixed` gets used to create visualizations of two binary
trees side-by-side for comparison.
Use is a little bit more involved than `func Draw`.

```go
fmt.Printf("digraph g1 {\n")
fmt.Printf("subgraph cluster_0 {\n\tlabel=\"before\"\n")
tree.DrawPrefixed(os.Stdout, originalTree, "orig")
fmt.Printf("\n}\n") // close cluster_0
fmt.Printf("subgraph cluster_1 {\n\tlabel=\"after\"\n")
tree.DrawPrefixed(os.Stdout, changedTree, "mangled")
fmt.Printf("\n}\n") // close cluster_1
fmt.Printf("\n}\n") // close g1
```

Your code has to "open" and "close" the individual cluster's
GraphViz subgraphs.
You can do more than 2 subgraphs.
`func Draw` is implemented by callling `func DrawPrefixed`
with hard-coded arguments.

## Tree Properties

### Binary Search Tree

```go
func BstProperty(root *Node) bool
```

returns true if `root` (representing a numeric-valued binary tree)
has the official "binary search tree property",
where all of the nodes in the left subtree have a value less than root's value,
and all of the nodes in the right subtree have a value greater than root's value.
This allows for no duplicate value nodes.

A lot of puzzles specify creating a tree with this property,
or input trees with this property.
I'm sure that interviewers have asked merely how to check
if an input tree has this property.
It's worth having this to demonstrate that puzzles/problems
got solved correctly.

### Depth (or height) of tree

```go
func FindDepth2(root \*Node, depth int) (depth int, deepnode *Node)
```

Find the depth (or height, depending on how you want to look at it)
of a given binary tree.
This is the answer to one puzzle that asked merely for
an algorithm for the depth of a tree,
and support for the "compose minimum height tree from sorted array"
puzzle.

