# Binary Tree Intervew Questions

So very many of the "programming puzzle of the day",
or "dev job interview questions"
relate to binary trees.
This repo consolidates a number of scattered tree implementations into a Go package,
and includes code that solves typical questions.

## Building

Support code resides entirely in package `tree`,
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

This was another Daily Coding Problem. Can't remember how easy it was said to be.

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
