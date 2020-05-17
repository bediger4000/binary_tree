# Binary Tree Intervew Questions

So very many of the "programming puzzle of the day",
or "dev job interview questions"
relate to binary trees.
This repo consolidates a number of scattered tree implementations into a Go package,
and includes code that solves typical questions.

## Building

The interview questions' code is almost entirely in package `tree`,
which lives in the `tree/` directory.

The demonstrations of the questions usually live in the top-level directory.

Build and use goes something like this:

    $ go build invert.go
    $ ./invert > invert.dot
    $ dot -Tpng -o invert.png invert.dot

A lot of the programs output [GraphViz](https://graphviz.org/) [dot-format](https://graphviz.gitlab.io/_pages/doc/info/lang.html) descriptions
of the answers.

## Questions and programs

* [invert](invert.go) a binary tree. I had it create a binary search tree so that the inversion is obvious.
* Create a [randomly valued](random.go) tree
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

The in-order traversal gives you an ordering of the elements.
You can reconstruct the original binary tree by adding elements
to a binary search tree in the pre-order traversal order, with "<=>"
determined by the in-order traversal.
