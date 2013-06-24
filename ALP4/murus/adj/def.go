package adj

// (c) Christian Maurer   v. 121101 - license see murus.go

import
  . "murus/obj"
type
  AdjacencyMatrix interface {
// Sqare matrices with entries of some arbitrary type.
// The element of an object x row i and column k is called x(i,k).
// These matrices can be interpreted as graphs without loops:
// a(i,k) != e means, that there is an edge in the corresponding
// graph, outgoing from its i-th and incoming at its k-th node.
// The nodes of the graph and e are the objects given by New.

// x has the nodes n[i] and len(n) of rows/columns
// and e as empty Edge.
////  New (n []Any, e Any) *Imp

  Object

// Returns true, if ! x(i,i) for all i < x.Num().
// i.e. the corresponding graph does not contain loops.
  Ok () bool

// Returns the smallest i s.t. x(i,i) != e, if such exists;
// returns x.Num otherwise.
  Loop () uint

// Returns the number of rows/columns of x, defined by New.
  Num () uint

// Returns for i < x.Num() the i-th object of x, given by New;
// returns otherwise nil.
  Node (i uint) Any

// Returns for i, k < x.Num() the edge of x outgoing from
// its i-th incoming at its k-th node; returns e, if
// i >= x.Num() or k >= x.Num() or if there is no such edge.
  Edge (i, k uint) Any

// Returns true, iff x(i,i) == e for all i < x.Num()
// and x(i,k) == x(k,i) for all i, k < x.Num(), i.e.
// iff the corresponding graph does not contain loops
// and is undirected.
  Symmetric () bool

// Returns true, iff x(i,i) == e for all i < x.Num()
// and there is no pair, s.th. x(i,k) == x(k,i), i.e.
// iff the corresponding graph does not contain loops
// and is directed.
  Directed () bool

// TODO Spec
  Add (y AdjacencyMatrix)

// TODO Spec
  Write (l, c uint)

// x is mirrored at its diagonal, i.e. x(i,k) equals f(k,i) for the
// former entries f of x, i.e. if the corresponding graph is undirected
// nothing has happened, otherwise, all its edges are inverted.
  Invert ()

// If i or k >= x.Num(), nothing has happened. Otherwise:
//// If i != k, then x(i,k) == a (x(i,i) == e).
// x(i,k) == a.
  Put (i, k uint, a Any)

// x(i,k) == e.
  Del (i, k uint)

// Returns true, if all rows of x contain at least an entry "true",
// i.e. iff in the corresponding graph every node
// has at least one outgoing edge.
  Full () bool
}
