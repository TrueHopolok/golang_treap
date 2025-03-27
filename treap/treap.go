/*
[Treap] is a binray tree with heap property that realized via random priority generation.
This package's data structure uses [Implicit keys] variation.
This allows to work as dynamic array with ability to split, merge, insert, delete, find in a logarithmic time.

# Package is unsafe to be used in parallel goroutines.

[Treap]: https://en.wikipedia.org/wiki/Treap
[Implicit keys]: https://en.wikipedia.org/wiki/Treap#Implicit_treap
*/
package treap

import (
	rand "math/rand/v2"
)

/*
Internal struct that is the treap itself.
Not used as a main type since cause problems with initialization.
*/
type node struct {
	value    int
	size     int
	priority int
	lson     *node
	rson     *node
}

/*
Recalculate node's size by checking all children's sizes.

# Time complexity:
  - Constant - requires constant amount of operations;
*/
func sync(n *node) {
	if n == nil {
		return
	}
	n.size = 1
	if n.lson != nil {
		n.size += n.lson.size
	}
	if n.rson != nil {
		n.size += n.rson.size
	}
}

/*
Merges 2 nodes into 1 node with its root being node with the highest priority.

# Time complexity:
  - Logarithmic - time complexity is equal to height of the treap;
*/
func merge(n1 *node, n2 *node) *node {
	if n1 == nil {
		return n2
	} else if n2 == nil {
		return n1
	}

	if n1.priority > n2.priority {
		n1.rson = merge(n1.rson, n2)
		sync(n1)
		return n1
	} else {
		n2.lson = merge(n1, n2.lson)
		sync(n2)
		return n2
	}
}

/*
Splits node into 2 by provided index.

	if index out of range: do nothing

# Time complexity:
  - Logarithmic - time complexity is equal to height of the treap;
*/
func split(n *node, index int) (l *node, r *node) {
	if n == nil {
		return nil, nil
	}

	if index < 0 {
		return nil, n
	} else if index >= n.size {
		return n, nil
	}

	position := index
	if n.lson != nil {
		position -= n.lson.size
	}

	if position < 0 {
		// split left son
		l, r = split(n.lson, index)
		n.lson = r
		sync(n)
		return l, n
	} else if position > 0 {
		l, r = split(n.rson, index)
		n.rson = l
		sync(n)
		return n, r
	} else {
		r = n.rson
		n.rson = nil
		sync(n)
		return n, r
	}
}

/*
Saves nodes values into provided slice.

Requirements:
  - Slice with the size of the treap;
  - Position set to 0.
  - Provided node being the root of the treap

This requirements are necessary on the 1st function call.

	if requirements not satisfied: may throw a panic

# Time complexity:
  - Linear - time complexity is equal to size of the treap;
*/
func export(values []int, position int, n *node) {
	if n == nil {
		return
	}
	if n.lson != nil {
		export(values, position, n.lson)
		position += n.lson.size
	}
	values[position] = n.value
	export(values, position+1, n.rson)
}

/*
Main type of a data structure that stores a single pointer.
That pointer is pointing to the root node.

This type shouldn't be used to initialize a varible.
Use `Create()` or `*Treap` type instead.
*/
type Treap struct {
	root *node
}

/*
Correctly initialize a Treap data structure.
Insert all given values to the back by calling `PushBack()` method.

# Time complexity:
  - Loglinear - time complexity is equal to height of the treap multiplied by amount of provided values;
*/
func Create(values ...int) *Treap {
	t := &Treap{nil}
	t.PushBack(values...)
	return t
}

/*
Merges 2 treaps. Returns resulted treap.
Old treaps must not be used afterwards.

# Time complexity:
  - Logarithmic - time complexity is equal to height of the highest treap;
*/
func Merge(t1 *Treap, t2 *Treap) *Treap {
	if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	}
	return &Treap{merge(t1.root, t2.root)}
}

/*
Split treap by provided index.
Returns 2 resulted treaps:

	1st: treap index <= given index
	2nd: treap index >  given index

Old treap must not be used afterwards.

# Time complexity:
  - Logarithmic - time complexity is equal to height of the highest treap;
*/
func Split(t *Treap, index int) (tl *Treap, tr *Treap) {
	if t == nil {
		return nil, nil
	}
	tl, tr = &Treap{nil}, &Treap{nil}
	tl.root, tr.root = split(t.root, index)
	return
}

/*
Insert value into provided index.
Splits treap into 2 parts.
Merges with new node and return the resulted treap.

In case index out range method calls:

	if index <= 0: t.PushFront(value)
	if index >= size-1: t.PushBack(value)

# Time complexity:
  - Logarithmic - time complexity is equal to height of the highest treap;
*/
func (t *Treap) Insert(index int, value int) {
	if t == nil {
		return
	} else if t.root == nil {
		t.root = &node{value: value, size: 1, priority: rand.Int(), lson: nil, rson: nil}
		return
	}
	if index <= 0 {
		t.PushFront(value)
		return
	} else if index >= t.root.size-1 {
		t.PushBack(value)
		return
	}
	l, r := split(t.root, index-1)
	l = merge(l, &node{value: value, size: 1, priority: rand.Int(), lson: nil, rson: nil})
	t.root = merge(l, r)
}

/*
Insert all provided values to the front of the treap.

# Time complexity:
  - Logarithmic - time complexity is equal to height of the highest treap;
*/
func (t *Treap) PushFront(values ...int) {
	if t == nil {
		return
	}
	var vroot *node
	for _, value := range values {
		n := &node{value: value, size: 1, priority: rand.Int(), lson: nil, rson: nil}
		vroot = merge(n, vroot)
	}
	t.root = merge(vroot, t.root)
}

/*
Insert all provided values to the back of the treap.

# Time complexity:
  - Logarithmic - time complexity is equal to height of the highest treap;
*/
func (t *Treap) PushBack(values ...int) {
	if t == nil {
		return
	}
	var vroot *node
	for _, value := range values {
		n := &node{value: value, size: 1, priority: rand.Int(), lson: nil, rson: nil}
		vroot = merge(vroot, n)
	}
	t.root = merge(t.root, vroot)
}

/*
Delete all elements in the given range.
Method works by splitting treap into 3 parts,
and then merging 2 necessary parts togheter.

Some properties of the deletion range:

	if index_left > index_right: do nothing
	if index_left >= size: do nothing
	if index_right < 0: do nothing

# Time complexity:
  - Logarithmic - time complexity is equal to height of the highest treap;
*/
func (t *Treap) Cut(index_left int, index_right int) {
	if t == nil {
		return
	} else if t.root == nil {
		return
	} else if index_left > index_right {
		return
	} else if index_right < 0 || index_left >= t.root.size {
		return
	}
	l, k := split(t.root, index_left-1)
	_, r := split(k, index_right)
	t.root = merge(l, r)
}

/*
Delete 1 element from the treap by provided index.
Method works by splitting treap into 3 parts,
and then merging 2 necessary parts togheter.

	if index < 0 || index >= size: do nothing

Method is replacement of a `Cut()` method but for 1 position to delete instead of range.

# Time complexity:
  - Logarithmic - time complexity is equal to height of the highest treap;
*/
func (t *Treap) Delete(index int) {
	if t == nil {
		return
	} else if t.root == nil {
		return
	} else if index < 0 || index >= t.root.size {
		return
	}
	l, k := split(t.root, index-1)
	_, r := split(k, index)
	t.root = merge(l, r)
}

/*
Returns size of a treap.

# Time complexity:
  - Constant - requires constant amount of operations;
*/
func (t *Treap) Size() int {
	if t == nil {
		return 0
	} else if t.root == nil {
		return 0
	}
	return t.root.size
}

/*
Return the element on the given index.

	if index out of range: return 0

# Time complexity:
  - Logarithmic - time complexity is equal to height of the highest treap;
*/
func (t *Treap) Find(index int) int {
	if t == nil {
		return 0
	} else if t.root == nil {
		return 0
	} else if index < 0 || index >= t.root.size {
		return 0
	}
	for n := t.root; n != nil; {
		position := index
		lson := n.lson
		var lsize int
		if lson != nil {
			lsize = lson.size
			position -= lsize
			index -= lsize
		}
		if position < 0 {
			index += lsize
			n = lson
		} else if position > 0 {
			index--
			n = n.rson
		} else {
			return n.value
		}
	}
	return 0
}

/*
Returns all values of the treap as slice of the integers.
All indexes are the same as in the treap.
This method is recommended if a lot of look up operations will be coming.

# Time complexity:
  - Linear - time complexity is equal to size of the treap;
*/
func (t *Treap) Export() []int {
	if t == nil {
		return nil
	} else if t.root == nil {
		return nil
	}
	values := make([]int, t.root.size)
	export(values, 0, t.root)
	return values
}
