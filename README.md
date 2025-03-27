# Treap with implicit keys on GO

This data structure is basicly dynamic array of integers. 

Use `go doc treap` or read comments in the package file to see full documentation.

### How to use?

```go
t := treap.Create() // to create new treap

t.PushBack(1, 2, 3, 4) // insert to the back of the treap
t.PushFront(1, 2, 3, 4) // insert to the front of the treap
t.Insert(4, 5) // insert 5 in the 4 position

t.Size() // return amount of the elements in the treap
t.Find(4) // return value of the element on the 4th position
t.Export() // return all elements' values in the new slice

t.Cut(0, 3) // Deletes all elements from 0th to 3rd indexes
t.Delete(0) // Delete 1 element from the 0th position
```
