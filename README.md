# Binary Tree Intervew Questions

So very many of the "programming puzzle of the day",
or "dev job interview questions"
relate to binary trees.
This repo consolidates a number of scattered tree implementations into a Go package,
and includes code that solves typical questions.

I do have other binary tree repos that illustrate problems too big
to fit in this repo:

* [Reconstruct a binary tree from a postorder traversal](https://github.com/bediger4000/postorder-tree-traversal)
* [AVL tree construction](https://github.com/bediger4000/avl_tree)

## Building

Support code resides entirely in [package tree](./tree),
which lives in the `tree/` directory.
Answering the questions or solving the puzzles almost always means
doing something different than the generic support code does,
so the code or algorithms to solve problems lives in the top level directory.

Build and use goes something like this:

    $ go build invert.go
    $ ./invert > invert.dot
    $ dot -Tpng -o invert.png invert.dot

A lot of the programs output [GraphViz](https://graphviz.org/) [dot-format](https://graphviz.gitlab.io/_pages/doc/info/lang.html) descriptions
of the answers.
That way you can visually check that the code does what it's supposed to do.

## Questions and programs

I have a bigger binary tree question,
recreate binary tree from its post-order traversal,
in its [own repo](https://github.com/bediger4000/postorder-tree-traversal)
Still can't believe they marked that one "medium".

* Create a [randomly valued](random.go) tree.
* Create a GarphViz [drawing](drawtree.go) of a tree.
This code creates a binary search tree (BST) by inserting values as they appear on the command line.
I believe you can create a BST of any shape by inserting the values of
nodes of a BST with the desired shape in breadth-first order.
* [invert](invert.go) a binary tree. I had it create a binary search tree so that the inversion is obvious.
For some reason, I made `func (p *Node) Invert()` a method of tree node struct and put it in the support code.
* First cut at [finding depth](tree_depth.go) of tree, carries a struct around.
* Second cut at [finding depth](tree_depth2.go) of tree, completely recursive, returns deepest node.

### Reconstruct a tree from traversals

This was one of the "Daily Coding Puzzle" problems.
A [clever solution](https://www.geeksforgeeks.org/construct-tree-from-given-inorder-and-preorder-traversal/) exists.
Isn't this O(n^2), though?
It's also a [leetcode](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/submissions/) problem.

Given pre-order and in-order traversals of a binary tree, write a function to reconstruct the tree.

For example, given the following preorder traversal:

    [a, b, d, e, c, f, g]

And the following inorder traversal:

    [d, b, e, a, f, c, g]

You should return the following tree:

        a
       / \
      b   c
     / \ / \
    d  e f  g

#### Analysis

The in-order traversal gives you an ordering of the elements.
You can reconstruct the original binary tree by adding elements
to a binary search tree in the pre-order traversal order,
with "<=>" determined by the in-order traversal,
instead of using <, >, == built-in operators to make comparisons.
[My code](reconstruct.go) constructs a `map[string]int` where the keys are strings
from the in-order traverse
and the values are the indices of those strings when they're
in-order in an array.
`func insert` in this code can decide which child to recurse down
by getting numeric values from the map and comparing those.
`func insert` looks a lot like an ordinary recursive function that
inserts values to create a binary search tree.

This suggests that my current tree package `func Insert` could be
generalized to accept a `Node`, a value of type `interface{}`,
a node-creation function probably of type `func(interface{}) Node`,
and a comparison function of type `func(interface{},interface{}) int`.
String-valued nodes, integer-valued nodes and floating-point valued
nodes could all be created and inserted based on what the comparison
function returned.

An interviewer asking this question would have to decide what they wanted from the candidate. 
If a candidate had that flash of insight that let them create the clever algorithm,
is that candidate suitable for an "enterprise" programming role where boring, grind-it-out,
lots of boilerplate and standard following is necessary?

### Return all paths from the root to leaves

Another "Daily Coding Puzzle".

Given a binary tree, return all paths from the root to leaves.

For example, given the tree

      1
     / \
    2   3
       / \
      4   5

it should return [[1, 2], [1, 3, 4], [1, 3, 5]].

#### Analysis

The phrasing of the answer seems to assume the use of Python.
[My program](tree_paths.go) creates a binary search tree from number representations
on the command line,
then traverses the newly-created tree.
It adds each node's value to the current path when that node
gets visited by the traverse.
At leaf nodes, it copies the path, and keeps that copy.
As the traverse leaves a node,
the code trims the node's value from the current path.
I had to write a recursive visitor function that includes
pre-order and post-order function calls,
and write a type that could be used to accumulate paths at leaf nodes,
and also kept the current path updated.

### Daily Coding Problem: Problem #540 [Easy]

In Ancient Greece,
it was common to write text with the first line going left to right,
the second line going right to left,
and continuing to go back and forth.
This style was called "boustrophedon".

Given a binary tree, write an algorithm to print the nodes in
boustrophedon order.

For example, given the following tree:

           1
         /   \
        /     \
      2         3
     / \       / \
    4   5     6   7

You should return [1, 3, 2, 4, 5, 6, 7].

#### Analysis

I [did this](bous.go) with two stacks, one to hold nodes for a rightward
breadth-first pass, the other for a leftward breadth-first pass.
I switched stacks from which to pop parent nodes when the current
parent-node-stack turns up empty.
You could probably perform the task with a single double-ended queue and
a special marker node that indicates when to change direction.

There is no way this is an "easy" whiteboarding problem.
The child nodes in a rightward pass are the parent nodes in the
succeeding leftward pass: you have to think an entire "layer" of the
tree ahead to decide what operation (pop or dequeue) to do in 
the next pass, and you have to think an entire layer behind to decide
in what order to push or enqueue the child nodes.
It's also very easy to make mistakes in the code,
because pushing child nodes on the stacks happens in different orders
depending on the direction the code is traversing the parent nodes' layer.

You'll want to try 1, 2, 3, and 4 deep input trees, both complete,
and with a partial deepest layer.

    $ ./bous 7 3 11 1 5 9 13 0 2 4 6 8 10 12 14 
    7 11 3 1 5 9 13 14 12 10 8 6 4 2 0 

should show you a boustrophedon traverse of a complete binary search tree
of depth 4.

### Cousin Nodes

This was another Daily Coding Problem.
Can't remember how easy it was said to be.

Two nodes in a binary tree can be called cousins if they are on the same level
of the tree but have different parents.
For example,
in the following diagram 4 and 6 are cousins.

        1
       / \
      2   3
     / \   \
    4   5   6

Given a binary tree and a particular node,
find all cousins of that node.

#### Analysis

At first, I thought this was a fairly bogus question,
but this might actually be a good interview problem.
It does contain the opportunity to do a recursive or iterative search
of a binary tree, allowing the candidate to demonstrate algorithm knowledge.
It has a task that requires synthesizing a solution from several parts
(finding a node, finding a parent node, finding uncle, finding cousin(s)),
allowing the candidate to demonstrate thinking through a nonsensical problem,
and let's face it, business logic is often an ambiguity wrapped up in a special
case, inside some regulatory capture. 
It has the opportunity to point out test cases (can't find node, can't find
grandparent, can't find uncle, 0 through 2 cousins),
and it allows the candidate to demonstrate some insight (only 1 node can be
parent of cousins, it's a binary tree).

### Prune a tree so that subtrees containing all 0s are removed

Given a binary tree where all nodes are either 0 or 1, prune the tree so
that subtrees containing all 0s are removed.

For example, given the following tree:

       0
      / \
     1   0
        / \
       1   0
      / \
     0   0

should be pruned to:

       0
      / \
     1   0
        /
       1

We do not remove the tree at the root or its left child because it still
has a 1 as a descendant.

#### Analysis

My [code](prune_tree.go) does this with a post-order traverse
of the entire original tree, pruning the tree as it goes.

By pruning left and right children,
then checking to see if both child nodes are null and
the current node's data is 0,
the code can return the node or nil.

A leaf node with 0 value gets pruned by setting its parent's pointer
to it to nil.
If both child pointers of the parent have a nil value,
or get set to a nil value in the traverse,
the parent can be pruned by returning nil.

Perceiving a post-order traverse as a way to prune requires
the ability to not shy away from doing something via brute force,
and also a willingness to put off the work of pruning until
both of a node's subtrees have been pruned.

I suspect that a lot of candidates would try to traverse the
tree and write a function `prunable(node *Node) bool`,
calling pruneable on each child node.
This would get you wrapped around the axle of recursion,
because `prunable()` would have to recurse each node's two
child sub-trees, and then you'd want to write code to recurse into
each of the sub-trees.

This is probably a pretty good interview question,
if you want to see if the candidate has that flash of intuition,
which may not arrive during the pressure of a whiteboard question.

### Minimum height binary search tree

I seem to have mislaid all other information about this problem
other than this simple statement:

Given a sorted array increasing order of unique integers,
create a binary search tree with minimal height.

#### Analysis

Ordinarily, this would take quite a bit of work to get correct,
but the "sorted array" clause makes it possible.

[My code](minimal_ht_tree.go) take advantage of the sort by
using the middle element of the array as the root of the root.
The sub-arrays on either side of the middle element will be
binary search trees of the left and right children of the root
of the tree,
so the code can recurse.
The base case actually comprises 3 special cases:
1, 2 and 3 element sorted arrays each get treated on their own.
it was too hard to deal with a length 2 array with the general case
code, and 3 element was too easy to not treat specially.
Also, there's a choice of two arrangements for the length 2 case:

    1
     \
      2

or

      2
     /
    1

I have my code choosing one of them pseudo-randomly,
just for equity.
That is,
it can create different minimal height trees from different runs with the same input.

I also write the [straightforward version](minimal_ht_tree2.go) of this.
It's a lot clearer,
but it always creates the same tree from any given output.

This seems like a harder interview question,
the interviewer should prepare to prompt the candidate.
It requires some insight to notice that "sorted array" and
"binary search tree" have a relationship
via the middle-node-becomes-root,
and that the relationship is recursive.
Also, the choice of middle node for an even-length sorted array
makes the actual coding a little harder.
The candidate probably writes code to always chose one or the other.
Calculating the middle index of a 0-indexed array causes humans troubles.
It's easy to be off-by-one, and it's not the more familiar fence post problem.
The interviewer shouldn't expect anything close to a good solution.

The candidate could amaze the interviewer by providing insightful test cases.
Not only is a 1-node tree a good test case,
but so are 2-node trees.
Test cases that cause the algorithm to decide between two
"middle indexes" of the array would be good.
Something like an input of [0, 1, 2, 3] has 4 different trees of height 3.
That's worth a test.
"Slightly incomplete" trees, missing only 1, 2, 3 or so nodes
from having a complete set of leaf nodes would be good test cases.

The candidate could also amaze the interviewer by proposing a test
for whether or not the tree is of minimal height.
A complete binary tree of depth D has N = 2<sup>D</sup>-1 nodes.
We know how many nodes are in the tree, we got an array as input.
Find the depth of the tree, see if it's less than or equal to log2(N+1)+1.
[This program](testmin.go) does that.

### Daily Coding Problem: Problem #545 [Hard]

Given a binary tree, find the lowest common ancestor (LCA) of two
given nodes in the tree. Assume that each node in the tree also has a
pointer to its parent.

According to the definition of LCA on Wikipedia: The lowest
common ancestor is defined between two nodes v and w as the lowest
node in T that has both v and w as descendants (where we allow a node
to be a descendant of itself).

#### Analysis

It says "binary tree", not "binary search tree",
so you can't assume an ordering.

The "assume that each node in the tree also has a pointer to its parent"
is almost certainly a clue that they don't want the obvious solution,
which is to find paths to V- and W-valued nodes,
then compare paths to find the last common ancestor.
The node that appears in both paths last is deepest in the tree.

Apparently [this 18 page paper](https://www.cs.bgu.ac.il/~segal/PAPERS2/tarj.pdf)
describes a sub-linear algorithm for finding the LCA,
using back-links to parents.
It's entirely unobvious, and I refuse to bother with it.

This is a weird problem to ask in an interview.
Unless the interviewer wants a candidate who's read, understood,
and memorized all of Robert Tarjan's many, many algorithms,
nobody will pass this. Everyone will do the O(n) time algorithm,
or waste all their time trying to recreate something inobvious.

### Daily Coding Problem: Problem #502 [Easy]

Given a binary tree, determine whether or not it is height-balanced. A
height-balanced binary tree can be defined as one in which the heights
of the two subtrees of any node never differ by more than one.

#### Analysis

This is esssentially the tree depth (or tree height) problem
framed differently.
As such, it's prey to all of the tree depth problem's difficulties.
The interview candidate might fall into the trap of tryig to
write a `Balanced()` function that's recursive on its own,
rather than finding max depth of each subtree then ensuring that
any depth difference is not too great.
The interviewer might not get a feel for the candidate's
coding ability at all.
