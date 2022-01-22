# Binary Tree Puzzles and Interview Questions

Many of the "programming puzzle of the day",
or "dev job interview questions"
relate to binary trees.
This repo contains a binary tree implementation in a Go package,
as well as code that solves puzzle-or-problem-of-the-day questions.

I do have other binary tree repos that illustrate problems too big
to fit in this repo:

* [Reconstruct a binary tree from a postorder traversal](https://github.com/bediger4000/postorder-tree-traversal)
* [AVL tree construction](https://github.com/bediger4000/avl_tree)
* [Multi-child tree symmetry](https://github.com/bediger4000/tree_symmetry)

Some of the problems and puzzles below should be their own repos because of their size,
but the convenience of a single binary tree package is too great to break them out.

## Building

Support code resides entirely in [package tree](./tree),
which lives in the `tree/` directory.
Answering the questions or solving the puzzles almost always means
doing something different than the generic support code does,
so the code or algorithms to solve problems lives in the top level directory.

Build and test goes something like this:

    $ make all
    $ ./runtests
    $ cd tree
    $ go test -v .

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

---

### Daily Coding Problem: Problem #842 [Medium]  

This problem was asked by Google.

Invert a binary tree.

For example, given the following tree:

```
    a
   / \
  b   c
 / \  /
d   e f
```

should become:

```
  a
 / \
 c  b
 \  / \
  f e  d
```

[Code #1](invert.go)

[Code #2](invert2.go)

[Code #3](invertnode.go)

For attempt #1,
I had it create a binary search tree of random-valued
nodes so that the inversion is obvious.
I wrote this one before generalizing `tree.NumericNode` and `tree.StringNode`
into interface `tree.Node`.
At that time, the data and left, right child pointers weren't exported,
nor did access functions exist.
I made `func (p *NumericNode) Invert()` a method of tree node struct and put it in the support code.

Attempt #2 I wrote well after generalizing numeric and string node types,
but I chose to use `tree.StringNode` so that I could get exactly
the output tree that the problem statement specifies.
I also used `func invert(node *tree.StringNode)` instead of
a method on the node's type.
I'm torn about object oriented programming,
and "inverting a tree" seems less like a operation performed on
a node, which would seem to need a method,
and more like operating on the entire data structure.

Attempt #3 uses `tree.Node`  Go interface types to see if that's possible,
and it is.

In all attempts, I found that checking for a nil node pointer
is best done in function or method `invert`,
to avoid complicating the inversion function with 2 tests for nil
children.
This has become a recurring motif.

This has been at least 3 different "Daily Coding Problems",
two rated "[Medium]"

---

### Daily Coding Problem: Problem #752 [Easy]

Also, Problem #107.

This problem was asked by Microsoft.

Print the nodes in a binary tree level-wise.
For example, the following should print 1, 2, 3, 4, 5.

```
  1
 / \
2   3
   / \
  4   5
```

#### Analysis

[Breadth-first traverse](breadthfirst.go) iterative traverse of tree.

This is an old one: instead of using a stack
(implicit function call stack, or an explicit data structure),
the algorithm uses a FIFO queue to keep track of
its place in the traverse of the tree.

```sh
$ go build breadthfirst.go
$ ./breadthirst '(1(2(4)(5))(3(6)(7)))'
1 2 3 4 5 6 7
$
```

Due to the simplicity of this problem,
maybe interviewers should use it only on entry-level candidates.

---

### Daily Coding Problem: Problem #622 [Easy]

This problem was asked by Google.

Given the root of a binary tree, return a deepest node.
For example, in the following tree, return d.

        a
       / \
      b   c
     /
    d

#### Analysis

* First cut at [finding depth](tree_depth.go) of tree, carries a struct around.
* Second cut at [finding depth](tree_depth2.go) of tree, completely recursive, returns deepest node.
* Third cut at [finding depth](tree_depth_node.go) of tree. Uses `tree.Node` interface,
`func tree.AllorderTraverseVisit`. No globals involved.
* Fourth cut at [finding depth](tree_depth_chan.go) of tree, using Go channels.
This is an example of Go-peculiar program structure,
leveraging channels and goroutines.

The problem statement confuses data with data structure.
The "deepest node" of the example isn't 'd', but 'd' is the value of the deepest node.
Interviewers should be clear, in my opinion.

This is a question for an entry-level position interview.
It involves a traverse of a binary tree where the value of the
nodes is only for identification,
so pre-, post- or in-order doesn't matter.
The only things to do are recognize a leaf node,
and keep track of depth in the tree.
The interviewer could look for orderly design process,
neat and tidy coding,
entry-level things like that.

If the candidate were to suggest test cases,
a single-node tree, left-heavy like the example,
a complementary right-deep test,
and some case in the middle might be in order.

I did the third version as an exercize: see if visitor functions and
interface type `tree.Node` could do the task.

I did the fourth version to show a Go-peculiar design pattern.
Have a goroutine generate some plausible values,
writing them to a channel.
The main goroutine reads from the channel and filters the plausible values
into desired value(s).
Recursive functions are simplified,
since the code for filtering the plausible values lives in the main goroutine,
and nothing has to be returned: possible answers get written to a channel.

---

### Reconstruct a tree from traversals

This is Daily Coding Problem: Problem #435 [Medium]

Given pre-order and in-order traversals of a binary tree,
write a function to reconstruct the tree.

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

A [clever solution](https://www.geeksforgeeks.org/construct-tree-from-given-inorder-and-preorder-traversal/) exists.
Isn't this O(n<sup>2</sup>), though?
It's also a [leetcode](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/submissions/) problem.

The in-order traversal gives you an ordering of the elements.
You can reconstruct the original binary tree by adding elements
to a binary search tree in the pre-order traversal order,
with "<=>" determined by the in-order traversal,
instead of using <, >, == built-in operators on data values to make comparisons.
[My code](reconstruct.go) constructs a `map[string]int` where the keys are strings
from the in-order traverse
and the values are the indices of those strings when they're
in-order in an array.
`func insert` in this code can decide which child to recurse down
by getting numeric values from the map and comparing those.
`func insert` looks a lot like an ordinary recursive function that
inserts values to create a binary search tree.

At the time I encountered this problem,
it suggested that my tree package `func Insert` could be generalized.
Originally, type `tree.Node` carried an int value.
I changed the int-valued struct to `tree.NumericNode`,
and made `tree.Node` into a Golang interface.
Later I added `tree.StringNode`,
for problems that require a binary tree of strings.

I created a generalized binary tree node,
type `tree.Node`, a Go interface type.
Types `tree.NumericNode` and `tree.StringNode`,
both structs,
have accompanying methods that make them satisfy the `tree.Node` interface.
Function `tree.GeneralCreateFromString` can and is used to
create binary trees with integer value nodes (`func CreateNumericFromString`)
and string value nodes (`func CreateFromString`).

#### Interview Analysis

An interviewer asking this question would have to decide what they wanted from the candidate.
If a candidate had that flash of insight
that let them create the clever algorithm, is that candidate suitable
for an "enterprise" programming role where boring, grind-it-out, lots of
boilerplate and standard following is necessary?

---

### Return all paths from the root to leaves

Another daily coding puzzle, "Daily Coding Problem: Problem #587 [Medium]".

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

I'd say this is actually about a medium difficulty interview problem,
suitable for whiteboarding with an above-junior-level programmer.
The candidate would have to understand recursive functions,
but use those recursive functions as a scaffold for the work to find the answer.

---

### Daily Coding Problem: Problem #540 [Easy]

Also: Daily Coding Problem: Problem #810 [Easy]

This problem was asked by Morgan Stanley.

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

---

### Cousin Nodes #1

#### Daily Coding Problem: Problem #487 [Medium]

This problem was asked by Yext. Whoever Yext is.

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

This solution doesn't meet the problem statement.
See Cousin Nodes #2, below.
I was misled by my intuition of what "cousins" means
geneologically in the USA and by the example given.

[My solution](cousins.go).

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

---

### Cousin Nodes #2

This problem was asked by Yext. Whoever Yext is.

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

[My solution](cousins2.go).

I see that I did Cousin Nodes #1 incorrectly.
I misread the problem statement: cousin nodes aren't
nodes that share a grandparent node, they're not biological cousins.
Cousin nodes are just in the same depth in the tree.

The solution has 3 parts:

1. Construct the tree
2. Find depth of the "particular node".
3. Find the parent of the "particular node".
4. Traverse tree, printing out or otherwise collecting
all nodes at the same depth as the "particular node"

This is actually easier than the problem that I made up
when I misread the problem statement.
Finding the depth of the "particular" node,
and finding all nodes at the same depth as that node
are completely recursive, and can be written compactly.
With a little care about recursion termination,
you can be pretty confident you've got the right answers.

I'm less certain this is a good interview question.
The heart of the difficulty is calling the nodes you want
to find "cousin" nodes.
"Cousin" has a well-defined familial or geneological meaning,
and it's not what this problem wants.
The explanatory diagram is misleading,
in that it implies that a biologically-inspired reading is correct.
It seems to me that this problem statement is the difficulty,
and that the interviewer might find out programming skill,
but will mainly be sarcastically amused when most of the candidates
don't get the correct answer.

---

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

---

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

I also wrote the [straightforward version](minimal_ht_tree2.go) of this.
It's a lot clearer,
but it always creates the same tree every time from any given output.

#### Interview Analysis

This seems like a harder interview question.
The interviewer should prepare to prompt the candidate.
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
Find the depth of the tree,
see if it's less than or equal to log<sub>2</sub>(N+1)+1.
[This program](testmin.go) does that.

![binary tree ht vs node count](minht.png?raw=true)

The above image shows that the actual, discrete,
tree heights follow a step function.
Height of a tree is constant while filling in the bottom row of leaves.
Complete binary trees have a height matching log<sub>2</sub>(N+1).
Binary trees with one leaf in the bottom row of leaveshave a height matching log<sub>2</sub>(N+1)+1.
All other trees have a height between the two values.

---

### Daily Coding Problem: Problem #545 [Hard]

Given a binary tree, find the lowest common ancestor (LCA) of two
given nodes in the tree. Assume that each node in the tree also has a
pointer to its parent.

According to the definition of LCA on Wikipedia: The lowest
common ancestor is defined between two nodes v and w as the lowest
node in T that has both v and w as descendants (where we allow a node
to be a descendant of itself).

#### Analysis

[My code](lca.go).

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

---

### Count nodes in less than linear time

This also appears as Daily Coding Problem: Problem #736 [Easy]

Given a complete binary tree,
count the number of nodes in faster than O(n) time.
Recall that a complete binary tree has every level filled except the last,
and the nodes in the last level are filled starting from the left.
"Complete" means: every level, except possibly the last,
is completely filled,
and all nodes in the last level are as far left as possible.
It can have between 1 and 2h nodes at the last level h.

#### Analysis

I know "complete binary tree" as "tree having the heap property".

Solutions easily found on the web don't do the actual work,
the authors merely describe how to do it.
It's harder to do it than to wave your hands about it.

Suppose you number the child pointers, 0 for left child pointers,
1 for right child pointers.
Give the leaf nodes labels based on which pointers the path
to the leaf node takes.
The leftmost leaf node is 000 in this diagram,
the rightmost is 111.
If you interpret the labels as binary numbers,
the leaf nodes are labeled sequentially.
Taking the bits of the binary numbers as instructions on which
child pointer to choose,
a node with a given label can be found in O(D) time.

![numbered binary tree](numberedtree.png?raw=true)

If the leftmost depth is 1 greater than the rightmost depth,
the tree isn't filled-in.
You can do a [binary search](https://en.wikipedia.org/wiki/Binary_search_algorithm)
on the labels of the leaf nodes,
because the digits of the labels are also the direction of
child node pointer on which to recurse to get to the leaf node.
You know that the leftmost node is labeled "00...00" and the
rightmost node is labeled "11...1"
(in this case it is 1 digit in length less than the leftmost node's label).
The next tree depth to probe is the node labeled (11..1 + 00..00)/2,
when interpreting node labels as numbers.
The division-by-2 can be accomplished by shifting the digits to the right,
minding the carry bit left from adding.
If you don't correctly handle the carry bit,
shifting the result of adding labels one bit to the right will give you gibberish when
viewed as the left/right pointer choices.

Another subtlety exists when choosing when you've found the
target of the search: I believe you want to compare the added-and-shifted
mid-point label to the left side of the binary search bracketing labels.
When the added-and-shifted mid-point equals the left-bracketing-label,
you've found the deepest rightmost node.

That gets you a O(log<sub>2</sub>N) search for the deepest
rightmost node.

[My code](tree_count.go) does just such a binary search.
It also counts the number of node accesses it does to find
the rightmost node in the final level of the tree.

![time complexity of node counting](countnodes.png?raw=true)

A log<sub>2</sub>(N) curve fits the experimental count of node accesses very well.
The constant factor here is about 9 - I just eyeballed that, I didn't curve fit.

#### Interview Analysis

This isn't a bad interview question,
if the interviewer is after a candidate's understanding of computer science.
If the interviewer is satisfied with a more-or-less handwaving explanation,
or giving this as a take-home problem, it's fine.
But it's not "easy" after the hand-waving
It contains a lot of subtleties that would cause wasted time in a whiteboard coding experience.
The candidate wouldn't demonstrate anything worthwhile,
just that they can puzzle over corner cases,
and understand how integer division works in their favorite language.
Letting a candidate analyze the problem out loud,
skipping over some of the details might be the best way to learn
if your candidate has problem solving abilities.
The node-labels-as-pointer-following-directions trick is also fairly subtle.
I discovered it by accident.
It's one of those "use a value as a number and also as something else" tricks
that can make a speedy algorithm, and give clarity to an analysis,
but are usually hard to see without lengthy puzzling over the problem.

---

### Daily Coding Problem: Problem #502 [Easy]

Given a binary tree, determine whether or not it is height-balanced. A
height-balanced binary tree can be defined as one in which the heights
of the two subtrees of any node never differ by more than one.

#### Analysis

[My code](ht_balanced.go).

This is esssentially the tree depth (or tree height) problem
framed differently.
As such, it's prey to all of the tree depth problem's difficulties.
The interview candidate might fall into the trap of trying to
write a `Balanced()` function that's recursive on its own,
rather than finding max depth of each subtree then ensuring that
any depth difference is not too great.
The interviewer might not get a feel for the candidate's
coding ability at all.

---

### Daily Coding Problem: Problem #133

Also, Daily Coding Problem: Problem #609 [Medium].

Given a node in a binary tree, return the next bigger element, also
known as the inorder successor.

For example, the inorder successor of 22 is 30.

      10
     /  \
    5    30
        /  \
      22    35

The Problem #609 version also says:

    You can assume each node has a parent pointer.

The example tree does have the Binary Search Tree property,
but the written statement doesn't say that input trees have that property.
My solution does not assume the input tree has that property.
Perhaps it should have.

#### Analysis

[My code](inorder_successor.go).

I thought this would be easier than it was.
My initial idea was to just do an in-order traverse of the tree,
and if either child node has the given value,
return the value of the current node.

This is wrong, as it misses the case where any node with 2 children
(10 or 30 in tree above)
is the given value.
The successor value is the value of the right child node.
It also misses the case where the given value is 35 in the tree above,
the maximum value in the tree.
It has no inorder successor.

For the tree above,
each value in the tree has these inorder successor values:

|Given Value|Inorder Successor|
|-----------|-----------------|
|  5 | 10 |
| 10 | 22 |
| 30 | 35 |
| 22 | 30 |
| 35 | - |

There's a missing case in the example tree the candidate would want to check.

![example binary tree](inorder_successor.png?raw=true)

If the given value is 6,
the inorder successor node has a value of 10.

I think there's another corner case,
where the given value doesn't exist in the tree.
I suppose you could parse the problem's language
to declare that case doesn't exist.

I declare this problem statement to be cunningly misleading.

This isn't one of those "flash of insight necessary" problems.
In that respect, it's a decent interview problem.
The interviewer should probably tailor their expectations
for solutions based on what the candidate claims their experience is.
Less senior developers would be lucky to write code that partially works.
More senior developers might get a good analysis of the problem,
but still have trouble writing code that works on every case.

The candidate would do well to analyze test cases first for this problem.
That would give the candidate enough information to have a fighting
chance to get a correct algorithm.

[I ended up with code](inorder_successor.go)
that had a complicated recursive in-order traverse.
I interpreted "binary tree" in the problem to mean "unordered",
not a Binary Search Tree.
If "binary tree" means "binary search tree",
then my code isn't the greatest: it will take forever to search
for a given value in a large tree.

The search is one function that gets called with a `*tree.NumericNode`,
the given value,
and an instance of a 3-valued type.
The 3 values of this type represent:

1. Looking for the given value
2. Found the given value
3. Found the inorder successor

Note that the search function receives one of these 3 values
as an argument,
and it needs to consider what the values mean when returned
from a recursive call.

If the search function receives "found the given value" as an
argument, the current node is the inorder successor.
It returns its own value and "found the inorder successor".

If the current node is not the inorder successor,
the left sub-tree can contain the given value.
If the left child has the given value,
the current node is the inorder successor.
The search function can return the current node's value,
and "found the inorder successor".

If the current node has the given value,
either the right child is the inorder successor,
or the inorder successor is up the tree.

There's a subtlety if the programmer calls the search function
with a nil (or NULL) current node.
This is common practice because it avoids having two checks
(left and right child) for nil/NULL pointers,
and it declutters recursive traverse code.
A nil/NULL current node pointer has to return the value of the 3-valued type argument
it's called with, rather than "looking for the given value".

Of course, the quick-and-dirty method of creating an array from
an in-order traverse of the tree, finding the given value in the array,
then returning the value at the next index would also work.
It might be worth the candidate's time to mention this.
The drawbacks?
O(n) extra space for the array,
O(log N) extra time for a binary search of the array.
The recursive method only uses O(D) (D depth of tree) extra space.

I'm not at all sure what difference the parent pointers of the
Problem #609 version would make.
The "missing case" of the inorder successor of 6 in the tree above
basically makes the parent pointer useless.
I guess you could find the first number, then follow parent pointers
until you find the next node with a bigger value than the first number,
but off hand, I can't say that would always work.

---

### Daily Coding Problem: Problem #490 [Medium]

This problem was asked by Yelp.

The horizontal distance of a binary tree node describes how far left or right
the node will be when the tree is printed out.

More rigorously, we can define it as follows:

* The horizontal distance of the root is 0.
* The horizontal distance of a left child is hd(parent) - 1.
* The horizontal distance of a right child is hd(parent) + 1.

For example, for the following tree, hd(1) = -2, and hd(6) = 0.

               5
             /   \
           /       \
          3         7
        /  \      /   \
      1     4    6     9
     /                /
    0                8

The bottom view of a tree, then,
consists of the lowest node at each horizontal distance.
If there are two nodes at the same depth and horizontal distance,
either is acceptable.

For this tree, for example, the bottom view could be [0, 1, 3, 6, 8, 9].

Given the root to a binary tree, return its bottom view.

#### Analysis

This is a largish problem for a whiteboard,
the candidate who gets this question should talk it out,
put some very high level pseudocode on the whiteboard
to show you understand it before doing any "coding".

I doubt there's an actual "best solution".
[My solution](bottomview.go):

1. Build a tree with nodes that have depth and horizontal distance fields.
2. Traverse tree, filling in depth and distance fields.
3. Traverse tree again, filling in a Go map with the deepest
node for any given horizontal distance.
4. Construct an array of the deepest nodes that appear in the map.
5. Sort the array based on horizontal distance.
6. Print the sorted array.

These steps could be condensed into a single pass:

1. While building the tree, fill in depth and distance fields.
Get this working, it's harder than either building the tree
or filling in depth and distance.
2. While building the tree and calculating depth and distance,
insert nodes into the map.

I guess this would be an O(n) algorithm,
because it visits every single node in the tree.

This could be a good interview question,
if the interviewer adjusts expectations for the level of the job.
Entry level programmers could be coaxed to discuss binary trees,
and how they might calculate depth and distance recursively.
Mid-level programmers might sketch an overall solution,
and implement parts on the whiteboard.
Senior programmers might end up implementing the whole thing,
if allowed enough time.

---

### Daily Coding Problem: Problem #490 [Medium]

This question was asked by Apple.

Given a binary tree, find a minimum path sum from root to a leaf.

For example, the minimum path in this tree is [10, 5, 1, -1], which
has sum 15.

      10
     /  \
    5    5
     \     \
       2    1
           /
         -1

#### Analysis

[My solution](minpathsum.go).

First off, the request is phrased poorly:
do they want the sum, the path or both?
The candidate would be wise to ask.
Returning the path is more work than just the sum.

Since the problem seems to allow negative numbers as node values,
there's no way to short-circuit a complete traverse of the tree.
The recursion needs to carry around a minimum sum found so far,
and the associated path because of this.

Depending on what language the candidate does this problem,
solving it may entail some array management code
to keep from re-appending pieces of the paths through the tree.
That's really the only tricky piece to the recursion.

The candidate could distinguish themself by suggesting
test case input trees - a single element tree would be 
a decent test case,
as would extreme left- and right-hand-side minimum-sum-paths.
Maybe a tree with a large negative value in the leaf node would be
a good test case.

The interviewer should only expect a simple recursive solution.
There's no opportunity for cleverness or short-circuiting a
traverse of the entire tree.

This might actually count as a medium-difficulty problem.
It requires no great insight to solve, but the candidate would
have to incorporate a few extras in an otherwise simple recursive solution.

---

### Daily Coding Problem: Problem #482 [Medium]

This problem was asked by Google.

Given a binary search tree and a range [a, b] (inclusive), return the sum
of the elements of the binary search tree within the range.

For example, given the following tree:

        5
       / \
      3   8
     / \ / \
    2  4 6  10

and the range [4, 9], return 23 (5 + 4 + 6 + 8).

#### Analysis

The obvious solution is to do a full traverse of the input tree, 
carrying around a pointer to an int, the sum so far.
At each node, check if the value is in the range [a, b],
and add it to the pointed-to-int.
I did this in `func visit1` of my [solution code](bst_value_sum.go).

The problem says "Given a binary search tree".
That means that a program can use that property to avoid visiting
nodes that have values less than a and greater than b.
I did this in `func visit2` of my [solution](bst_value_sum.go).

My thought is that any time a binary tree problem or puzzle
says that the tree has the binary search property,
it's worth considering how to make that problem faster
using the BST property.
Seeing a candidate recognize and use that property is probably
what this question is all about, for an interviewer.

The candidate should consider test cases,
like giving a range that doesn't match any node's value,
or excludes all nodes, or includes only one node.
The shape of the input tree could also matter,
so consider a tree that's effectively a linked list:
all of the right-hand-children are non-null,
but none of the left-hand-children are.

I don't believe this qualifies as a "medium" hard question.
Consider Daily Coding Problem #475, which follows.
That's a "medium" too, and much more difficult.

---

### Daily Coding Problem: Problem #475 [Medium]

This problem was asked by Google.

Implement locking in a binary tree.
A binary tree node can be locked or unlocked only if all of its descendants
or ancestors are not locked.

Design a binary tree node class with the following methods:

* is_locked, which returns whether the node is locked
* lock, which attempts to lock the node.
If it cannot be locked,
then it should return false.
Otherwise, it should lock it and return true.
* unlock, which unlocks the node.
If it cannot be unlocked, then it should return false.
Otherwise, it should unlock it and return true.

You may augment the node to add parent pointers or any other property you
would like. You may assume the class is used in a single-threaded program,
so there is no need for actual locks or mutexes. Each method should run in
O(h), where h is the height of the tree.

#### Analysis

It's weird that "locking" takes places in a single-threaded program
with no need for actual locks or mutexes.
I suspect this is to avoid all the ugliness of what needs to get locked, and when.

The problem statment says:

    A binary tree node can be locked or unlocked only if all of its descendants
    or ancestors are not locked.

This translates to "locking a node locks the sub-tree below it",
and "you can't lock a node inside a locked sub-tree".

The problem has a giveaway hint:

    You may augment the node to add parent pointers or any other property you would like.

If you add parent pointers,
finding out if any ancestors are locked is O(log<sub>2</sub>(N+1)),
where N is the number of nodes in the tree, i.e. the depth of the tree.
The nodes can keep a count of locked descendants,
eliminating the need for traversing sub-trees to find any locked nodes..
Any node, even unlocked nodes,
that have more than zero locked descendants aren't eligible to lock.
Unlocking a node involves chasing parent pointers to the root,
decrementing the count of locked descendants.

[My solution](locktree.go).

I created a locking binary tree type,
and a program that lets you interactively create trees, lock and unlock nodes, and inspect trees and nodes.

    ./locktree
    Locked binary tree explorer
    > ?
    locking node binary tree explorer
    Operations:
    print - print lisp-like string rep of tree
    checkall - show lock status of all nodes
    check N - show lock status of node with value N
    lock N - lock node with value N
    unlock N - unlock node with value N
    find N - print info about node with value N
    create (...) - parse lisp-like tree rep, use it thereafter
    > create (0(1()(2))(3()()))
    > lock 2
    locked node with value 2 at 0xc0000983f0
    > print
    (0/U(1/U()(2/L))(3/U))
    > lock 1
    did not lock node with value 1 at 0xc0000983c0

The program didn't lock node with value 1 because node with value 2,
in it's right sub-tree, was already locked.

The conditions on locking and unlocking are to keep a sub-tree of a given
node locked, and not re-lock sub-trees of nodes in the locked sub-tree.
It's good that the problem says not to use mutexes,
and that the code is to be used in single-threaded programs,
because actually locking a node would seem to involve
locking the entire tree to chase parent pointers.
I'm at a loss to explain what's going on with this problem,
except that maybe it's nonsensical nature is to get candidates
to solve a problem they've never seen before.

From that standpoint,
the interviewer has to watch for 2 things:

1. That the candidate has the insight that a parent point
can allow O(h) lock and unlocks by chasing those parent pointers.
2. Programming mechanics.
Chasing pointers might be unfamiliar to people who don't do C or Go
any more. The C++ subculture is such that raw pointers seem to be
considered taboo.
Maybe this problem exists to seperate the pointer-familiar sheep,
from the pointer-less goats.

Perhaps the interviewer is supposed to be satisifed with a design,
or a design and an implementation with some flaws that would get ironed out
by a little testing.

---

### Daily Coding Problem: Problem #644 [Easy]

A unival tree (which stands for "universal value") is a tree where all
nodes under it have the same value.

Given the root to a binary tree, count the number of unival subtrees.

For example, the following tree has 5 unival subtrees:

```
   0
  / \
 1   0
    / \
   1   0
  / \
 1   1
```

<!--
In package tree lisp-like tree notation, that's:
(0 (1) (0 (1 (1) (1)) (0)))
-->

#### Build and run Unival tree program

```sh
$ go build unival.go
$ .unival '(0 (1) (0 (1 (1) (1)) (0)))'

5 unival subtrees
```

#### Unival tree analysis

I think I see the 5 unival subtress: 4 leaf nodes, which are vacuously unival trees, since all of their subtrees,
of which there are zero, have the same value.
The final unival tree is the all-1-value subtree.

The solution actually took me 1 decent false start,
and some thinking.
The solution is entirely recursive,
although the problem statement hides it.

Leaf nodes have no children, therefore they have the same data
value as their children.
A leaf node is a unival tree.
If the current node has the same value as its left and right children,
and the child nodes are roots of unival subtress,
the current node and its subtress form a unival subtree.
You can do this in a post-order traverse,
but there's a complicated decision about whether the
current node is the root of a unival subtree.

I'm going to go with this problem is under-rated.
It's at least a medium,
given the complicated tests to decide if the current node
heads a unival subtree.
You've got to handle either or both child nodes don't exist,
and you've got to decide if they exist,
do the make the current node the root of a unival subtree.
One child might be the root of a unival subtree,
but have a different value than the current node.

The candidate could score points by noting that one child
could be the root of a unival subtree,
while the other isn't,
so more testing is required to verify correctness.

---

### Daily Coding Problem: Problem #651 [Medium]

This problem was asked by LinkedIn.

Determine whether a tree is a valid binary search tree.

A binary search tree is a tree with two children, left and right, and
satisfies the constraint that the key in the left child must be less
than or equal to the root and the key in the right child must be greater
than or equal to the root.

#### Analysis

This particular problem statement seems sloppily worded.
It reads as if they only want the candidate to consider
3-node trees.
It also conflates "value of node's data" with "node itself".
Sloppiness like that will lead to amazing bugs.

The other oddity in the problem statement is that it says:

```
the key in the left child must be less than or equal to the root and the
key in the right child must be greater than or equal to the root.
```

When building such a tree, you get to make a decision:
does a value equal to that of the current node become the left
child node's value, or the right child node's value?

I'm not even going to implement this one: it's garbage.
This could be implemented by 2 if-then-else tests.
There's no art or craft to this one.

The only reason to ask this question, if my reading is correct,
is to see if a candidate can withstand ambiguity and do a really
dumb task without complaining.
Candidates who are asked to do this problem should reject that
corporation's job offer. There's better environments to work in,
don't take the job.

---

### Daily Coding Problem: Problem #442 [Hard]  

This problem was asked by Netflix.

A Cartesian tree with sequence S is a binary tree defined by the following two properties:

* It is heap-ordered,
so that each parent value is strictly less than that of its children.
* An in-order traversal of the tree produces nodes with values that correspond
exactly to S.

For example, given the sequence [3, 2, 6, 1, 9], the resulting Cartesian tree would be:

```
      1
    /   \   
  2       9
 / \
3   6
```

Given a sequence S, construct the corresponding Cartesian tree.

#### Analysis

There are [several](https://en.wikipedia.org/wiki/Cartesian_tree#Efficient_construction)
methods for constructing a Cartesian tree,
none of them obvious.

I haven't done this yet.

---

### Daily Coding Problem: Problem #422 [Easy] 

This problem was asked by Salesforce.

Write a program to merge two binary trees.
Each node in the new tree should hold a value equal to the sum of the values of
the corresponding nodes of the input trees.

If only one input tree has a node in a given position,
the corresponding node in the new tree should match that input node.

#### Analysis

[My version of it](merge.go) requires 2 trees on the command line,
each in a lisp-like syntax.
The "-g" flag gives you [GraphViz](https://graphviz.org/) dot-language output,
which can be converted into an image for viewing.

```sh
$ go build merge.o
$ ./merge -g '(1()(3(3)(3)))' '(1(2(2)(2))())'  > m.dot
$ dot -Tpng -o m.png m.dot
$ feh m.png # Or whatever your favorite viewer is
```

![2 trees and merged tree](merge.png?raw=true)

This makes a good interview problem for a candidate for a junior-level job.
It's just nonsensical enough that nobody has done it on the job,
but it requires some Computer Science and some experience.

The problem can be handled with standard binary-tree-recursive thinking.
Realize that you need a recursive function,
find the base case where recursion stops,
handle the recursive call(s).
There may be more than one base case,
where a nil argument pointer gets passed in to avoid duplicate tests for nil
where it handles recursive call(s).

The "merge" function takes two nodes (one from first tree, one from second tree)
and returns a merged node.

* If the merge function gets two nil nodes, it returns nil.
This is the case where recursion stops.
* If the merge functions get a non-nil and a nil node,
it creates a new node with the value of the non-nil node,
and calls the merge function on the non-nil node's children
to fill in left and right children of the new node.
Return the new node.
This is 2 cases in the code, left tree nil, right tree non-nil,
and vice versa.
* If the merge function gets 2 non-nil nodes,
ereate a new node with the sum of the input nodes' values,
call the merge function on the left children of each node
to create the left child of the new node.
Call the merge function on the right children of each node
to create the right child of the new node,
Return the new node.

This does require a little insight to realize that the
requirement for
```
If only one input tree has a node in a given position,
the corresponding node in the new tree should match that input node.
```
means chasing only the non-nil child.
You don't have to write a `CopyTree` function,
the merge function can take care of it.

The interviewer can watch for experience indicators,
like having the merge function handle both nil arguments
rather than having duplicated tests for non-nil in the code.
The interviewer could elicit more from the candidate by
asking for test cases,
or if the candidate has missed a trick,
suggesting a test case that triggers undesired behavior.

---

### Daily Coding Problem: Problem #405 [Hard]

This problem was asked by Apple.

Given a tree, find the largest tree/subtree that is a BST.

Given a tree, return the size of the largest tree/subtree that is a BST.

#### Analysis

[My version](largest_bst.go) does this from the bottom up:

* A nil pointer (it's in Go) is a zero-sized BST.
* A leaf node is a 1-node-sized BST.
* If a node's value is greater than it's left-child's value,
and less than it's right-child's value,
and both right and left sub-trees are BSTs,
then the size of the BST is 1 + size of left-subtree + size of right-subtree.
* It's possible for a node to not be the root of a Binary Search Tree,
or one or the other children to not have the BST property.
In that case, the code decides to use the largest of the 2 sub-tress
that have the BST property.

I used Go's multiple returns from a function to indicate the size of the largest
BST in the sub-tree, and whether or not the node is the root of a BST.

My faith in checking for nil pointers on entrance to the recursive function only grows.
This lets you keep your code cleaner,
without redundant left- and righ-child checks for nil that visually clutter the code.

I think this algorithm runs in O(n) where n is the number of nodes in the whole tree.
I don't think there's a better run time,
given that the algorithm has to check every interior node's value against the values of the left and right child nodes.

As far as an interview problem goes,
it's pretty good.
A candidate can solve it using the usual recursive algorithm reasoning,
finding a base case (nil pointer or leaf node) as the formal argument node,
then working through what the recursive function has to do to provide
the desired answer.

The recursive function has to watch for a few cases,
like one child pointer nil, the other non-nil,
and the cases of current node as root of BST vs not root of BST.

My solution assumes that the interviewer wants to find binary search trees
that go all the way to the leaf nodes.
It's possible that the interviewer wanted a BST that encompasses only part of the entire tree.
That would be a harder problem,
but the candidate could probably do some elaborate checking on left and right
child values even if they don't comprise roots of BSTs themselves.
This seems kind of nonsensical, though.

It looks to me like the interviewer could watch a candidate reason out the
recursive function,
especially the "one nil, one non-nil child node" cases to see if the candidate
has a good grasp of algorithms.
There's enough code to write to ensure that the candidate can actually write
programs in the language in question.
The problem of finding sub-trees that have the BST property
is probably not one
that many candidates will have run into at work,
so it'probably a good problem in the sense that nobody
has a memorized solution.

Compared to other "hard" problems, this one really isn't.
It's probably just a "medium".

---

### Daily Coding Problem: Problem #664 [Easy]

This problem was asked by Google.

Given a binary tree of integers,
find the maximum path sum between two nodes.
The path must go through at least one node,
and does not need to go through the root.

#### Analysis

I haven't done this one yet.

---

### Daily Coding Problem: Problem #702 [Medium]

This problem was asked by Google.

Given the root to a binary tree,
implement serialize(root),
which serializes the tree into a string,
and deserialize(s),
which deserializes the string back into the tree.

For example, given the following Node class

```
class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
```

The following test should pass:

```
node = Node('root', Node('left', Node('left.left')), Node('right'))
assert deserialize(serialize(node)).left.left.val == 'left.left'
```

#### Analysis

"Medium"!?! The serializing isn't particularly difficult,
but deserializing any Lisp S-expression type representation will be tough.
If you allow whitespace, it becomes a lot more difficult.

My `tree` package implements this in [create.go](tree/create.go).
Specifically, functions `tree.CreateFromString`, `tree.CreateNumericFromString`
do the deserialization,
and function `tree.Printf` does the serialization.
These functions get used in many of the other problem solutions.

[My code](serialize.go) is in Go, not Python as above,
but it meets the problem statement.
The [utility program](readtree.go) to create GraphViz output of a tree
almost meets the problem statement.

As an interview question this might not be bad.
The interviewer gets to see a lot of coding for the de-serialization part.
The candidate can note where corner cases like "Missing ')'" might cause problems.
A node with a right child, but no left child is also a test case.
This would be a harsh problem for whiteboarding - it's harder than many other
"[Hard]" coding problems.

---

### Daily Coding Problem: Problem #748 [Easy]

This problem was asked by Apple.

Given the root of a binary tree,
find the most frequent subtree sum.
The subtree sum of a node is the sum of all values under a node,
including the node itself.

For example, given the following tree:

```
  5
 / \
2  -5
```

Return 2 as it occurs twice: once as the left leaf,
and once as the sum of 2 + 5 - 5.

#### Analysis

---

### Daily Coding Problem: Problem #750 [Medium]

This problem was asked by Jane Street.

Generate a finite,
but an arbitrarily large binary tree quickly in O(1).

That is, generate() should return a tree whose size is
unbounded but finite.

#### Analysis

[Code](rando1.go)

I was unsatisfied with this program.
After writing it,
re-reading the problem statement makes me think
my solution isn't what they want at all. 
Setting the number of nodes in the "random" tree
might not be what the interviewer desires.

Beyond that, what is the "O(1)" associated with?
For sorting, the number of swaps is important.
For making a bunch of "random" things
you have to generate a bunch of at least pseudo-random numbers,
maybe to have randomly-valued nodes in the binary tree,
or maybe to decide which branch of a node to add any
further nodes.

If we use "choosing a random number" as the operation we track O() for,
my program misses the goal.
My program creates O(n) random numbers for node values.
It also creates O(log<sub>2</sub>n) random numbers to
decide how to partition an array of node values into
sub-trees.
I failed this interview.

#### Around the web

* [Stackoverflow](https://stackoverflow.com/questions/49502112/construct-binary-tree-in-o1) solutions
* [Daily Coding Problem](https://www.dailycodingproblem.com/blog/big-tree/https://www.dailycodingproblem.com/blog/big-tree/) blog solutions

The consensus appears to be te solution is "generate the tree lazily".
That is,
node-generation only happens if and when a node gets accessed.
This strikes me as cheating.

### Interview Analysis

Are you willing to cheat to pass an interview?
Maybe that's a good thing for something like a pen testing position,
but outside of that, I don't know.
I would be unlikely to come up with lazy tree generation as a solution
during an interview.
Maybe I could find it if this was a take-home problem.

In any case, lazy generation seems like it's very vaguely gestured to
by the problem statement.
If the candidate should decide on lazy generation,
they should probably ask questions of the interviewer to see if
that's what's desired.

It's entirely possible that the question is deliberately loosely phrased to
encourage candidates to ask questions, allowing the interviewer to see how a
candidate thinks.

---

### Daily Coding Problem: Problem #751 [Hard]

Also, Daily Coding Problem: Problem #808 [Hard]

This problem was asked by Palantir.

Typically, an implementation of in-order traversal of a binary tree has
O(h) space complexity,
where h is the height of the tree.
Write a program to compute the in-order traversal of a
binary tree using O(1) space.

#### Analysis

This question wants a "Morris traversal".

#### Interview Analysis
---

### Daily Coding Problem: Problem #793 [Medium]

This problem was asked by Yahoo.

Recall that a full binary tree is one in which each node is either a
leaf node,
or has two children.
Given a binary tree,
convert it to a full one by removing nodes with only one child.

For example, given the following tree:

```
         0
      /     \
    1         2
  /            \
3                 4
  \             /   \
    5          6     7
```

You should convert it to:

```
     0
  /     \
5         4
        /   \
       6     7
```

#### Analysis

I haven't done this one yet.

This can be solved recursively,
by removing single-children from the bottom up.
Recurse to leaf nodes,
then on the way back up the tree,
delete nodes with only 1 child.

#### Interview Analysis
---
### Daily Coding Problem: Problem #307

This problem was asked by Oracle.

Given a binary search tree,
find the floor and ceiling of a given integer.
The floor is the highest element in the tree
less than or equal to an integer,
while the ceiling is the lowest element in the tree
greater than or equal to an integer.

If either value does not exist, return None.

#### Analysis

I haven't done this one yet.

#### Interview Analysis

---

### Daily Coding Problem: Problem #36

Also: Daily Coding Problem: Problem #992 [Medium] 

This problem was asked by Dropbox.

Given the root to a binary search tree,
find the second largest node in the tree.

#### Analysis

The largest valued node in a binary search tree is the right-most node.
The second-largest-valued node depends on the shape of the tree.

```
    1        2        1
   / \      /          \
  0   2    1            2
```

The largest valued node in the above 3 trees has the value 2,
but the second-largest-value can be its parent or its left child.

[My code](secondlargest.go) prints the largest and second-largest values.

#### Interview Analysis

This is relatively easy for a Daily Coding Problem "[Medium]" problem.

The problem statement is problematic. 
"Second largest node in the tree"
confuses "node" with "value of node's data".
So what does "find the second largest node" mean?
I'm sure most people will find the 2nd-largest-value of nodes
in the tree, but it's certainly possible to decide that the
"second largest node" is the parent of the 2nd-largest population subtree.
Does "find the node" mean return a pointer to that node,
or returning the second-largest-value of nodes in the tree,
or just printing out the second-largest-value as I did?

This problem statement should be re-phrased
to ask for something less ambiguous.
An interviewer asking this question will get a lot of clarifying
questions from some candidates,
and some shocking code from a few candidates that both
think unconventionally, and are either afraid to ask questions,
or are so sure of themselves that they just do some bizarro thing
like figure out which immediate child of the root node has the biggest
tree.

Of the more conventionally-minded candidates,
or those candidates that understand the distinction between "data"
and "data structure",
this problem could potentially weed out those that don't know about
binary trees.

One lesson in this one for candidates is: ask clarifying questions.
Do you understand what "largest node" means?
What does "find the node" mean - return the node, return the value,
or just output something when the code receives that node as argument?

Another lesson is that these coding problems aren't always well thought out,
or even able to help an interviewer ascertain your skills. 

---

### Daily Coding Problem: Problem #936 [Medium]

This problem was asked by Google.

Given a binary search tree and a range [a, b] (inclusive),
return the sum of the elements of the binary search tree within the range.

For example, given the following tree:

```
    5
   / \
  3   8
 / \ / \
2  4 6  10
```

and the range [4, 9],
return 23 (5 + 4 + 6 + 8).

#### Analysis

I haven't done this one yet.
#### Interview Analysis

---
### Daily Coding Problem: Problem #1020 [Easy] 

This problem was asked by Google.

Given the root of a binary search tree,
and a target K,
return two nodes in the tree whose sum equals K.

For example, given the following tree and K of 20

```
    10
   /   \
 5      15
       /  \
     11    15
```

Return the nodes 5 and 15.

#### Analysis

I haven't done this one yet.

The statement says "binary search tree",
so instead of traversing the whole tree to find two nodes,
it might pay off to use the BST property with some node's value
to find K - value, if it exists.
If it doesn't exist, move to another node.
Use the BST property to decide which subtree to move to.

#### Interview Analysis
---
### Daily Coding Problem: Problem #1057 [Easy]

This problem was asked by Amazon.

Given an integer N, construct all possible binary search trees with N nodes.

#### Analysis

I haven't done this in a while, and not in this context.

---
### Daily Coding Problem: Problem #1072 [Easy]

This problem was asked by Facebook.

Given a binary tree, return the level of the tree with minimum sum.

#### Analysis

I haven't done this yet.

The problem seems poorly phrased:
it asks the candidate to assume numerically-valued nodes,
but are they integer values or floating point?
Summing floating point numbers of radically varying magnitude
can have unintuitive results.

---
### Daily Coding Problem: Problem #1098 [Easy] 

This problem was asked by Oracle.

Given a binary search tree, find the floor and ceiling of a given integer.
The floor is the highest element in the tree less than or equal to an integer,
while the ceiling is the lowest element in the tree greater than or equal to an integer.

If either value does not exist, return None.

#### Analysis

I haven't done this yet.

The phrase "binary search tree" is undoubtedly important.
