package gra

// (c) Christian Maurer   v. 130118 - license see murus.go

import (
  . "murus/obj"
  "murus/adj"
)
type
  Demo byte; const (
  Depth = Demo(iota); Cycle; Euler; TopSort; ConnComp; Breadth; SpanTree; nDemos
)
type
  Demoset [nDemos]bool

// Sets of nodes with an irreflexiva relation:
// two nodes are related, iff they are connected by an edge, where there
// are no loops (i.e. no node is connected with itself by an edge).
// If the relation is symmetric, the graph is called "undirected",
// if it is strict, "directed" (i.e. all edges have a direction).
//
// The edges have a number of type uint as value; either all edges
// have the value 1 ore their value is given by their type Valuator.
// The outgoing edges of a node are enumerated (starting with 0);
// the node, with which a node is connected by its n-th outgoing edge,
// is denoted as its n-th neighbournode.
//
// In any nonempty graph exactly one node is distinguished as postactual
// and one as actual node (those two nodes must not be identical).
// Each graph has an actual subgraph, an actual way and a nodestack.
//
// A subgraph U of a graph G consists of all nodes of G
// and of those edges of G, that connect only nodes of U.
// A path in a graph is a sequence of nodes and from each of those
// - excluding from the last one - an outgoing edge to the next node.
// A simple path is a path of pairwise disjoint nodes.
// A circle is a simple path with an additional edge
// from the last node of the path to its first.
// Paths and circles are subgraphs.
type
  Graph interface {

// Pre: e == nil or e is of Type Valuator.
// x is Empty (with andefined actual and postactual node,
// empty actual subgraph, empty actual way and empty nodestack).
// x is directed, if g, otherwise undirected.
// x has the type of n as nodetype.
// If e == nil, then x has no edgetype and all edges of x
// have value 1; otherwise, x has the type of e as edgetype,
// which defines the values of the edges of x.
//  New (d bool, n, e Any) *Imp

  Object

// marks the postactual node, if x is not empty.
  Marker

  Persistor

// If x is persistent, it is saved.
  Terminate ()

// Returns true, iff x is directed.
  Directed () bool

// Returns the number of nodes of x.
  Num () uint

// Returns the number of nodes in the actual subgraph of x.
  NumAct () uint

// Pre: p is defined on nodes.
// Returns the number of nodes of x, for which p returns true.
  NumPred (p Pred) uint

// Returns the number of edges of x.
  Num1 () uint

// If n is not of the nodetype of x or
// if n is already contained as node in x,
// nothing has happend. Otherwise:
// n is inserted as node in x.
// If x was empty, then n is now the postactual and actual node of x,
// otherwise, n is now the actual node and the former actual node
// is now the postactual node of x.
  Ins (n Any)

// If x is empty or has an edgetype or
// if the postacual node of x coincides with the actual node of x,
// then nothing has happened. Otherwise:
// An edge is inserted from the postactual to the actual node of x
// (if these two nodes were already connected by an edge,
// then its direction might have been changed.)
  Edge ()

// If x is empty or has no edgetype or
// if the postactual node of x coincides with the actual node or
// if e is not of the edgetype of x,
// nothing has happened. Otherwise:
// e is inserted in x as edge from the postactual to the actual node of x
// (if these two nodes were already connected by an edge,
// that edge is replaced by e).
  Edge1 (e Any)

// If x is empty or has an edgetype or
// if n or n1 is not of the nodetype of x or
// if n or n1 is not contained in x or
// if n and n1 coincide or
// if there is already an edge from n to n1,
// nothing has happened. Otherwise:
// n is now the postactual and n1 the actual node of x
// and there is an edge from n to n1.
  Edge2 (n, n1 Any)

// If x is empty or has no edgetype or
// if n or n1 is not of the nodetype of x or
// if n or n1 is not contained in x or if n and n1 coincide or
// if there is an edge from n to n1 or
// if e is not of the edgetype of x,
// nothing has happened. Otherwise:
// n is now the postactual and n1 the actual node of x and
// e is inserted in x as edge from n to n1
// (if n and n1 were already connected by an edge,
// that edge is replaced by e).
  Edge3 (n, e, n1 Any)

// Pre: x has no edgetype.
// TODO Spec
  Matrix () *adj.Imp

// Pre: x has no edgetype.
// TODO Spec
  Def (a *adj.Imp)

// Returns true, iff
// the postactual node does not coincide with the actual node of x and
// there is an edge from the postactual to the actual node in x.
  Edged () bool

// Returns true, iff
// the postactual node does not coincide with the actual node of x
// and there is an edge from the actual to the postactual node in x.
  EdgedInv () bool

// Returns true, iff n is contained as node in x.
// In this case n is now the actual node of x.
// The postactual node of x is the same as before.
  Ex (n Any) bool

// Returns true, if n and n1 are contained as nodes in x and do not coincide.
// In this case n now is the postactual and n1 the actual node of x.
  Ex2 (n, n1 Any) bool

// Pre: p is defined on nodes.
// Returns true, if there is a node in x, for which p returns true.
// In this case some such node is now the actual node of x.
// The postactual node of x is the same as before.
  ExPred (p Pred) bool

// Pre: p is defined on edges.
// Returns true, iff there is an edge in x, for which p returns true.
// In this case the neighbour nodes of some such edge are now
// the postactual and the actual node of x (if x is directed,
// the node, from which the edge goes out, is the postactual node.
  ExPred1 (p Pred) bool

// Pre: p and p1 are defined on nodes.
// Returns true,
// iff there are two different nodes n and n1 with p(n) and p(n1).
// In this case now some node n with p(n) is the postactual node
// and some node n1 with p1(n1) is the actual node in x.
  ExPred2 (p, p1 Pred) bool

// Returns nil, if x is empty.
// Returns otherwise a copy of the actual node of x.
  Get () Any

// Returns nil, if x is empty or has no edgetype or
// if there is no edge from the postactual to the actual node of x or
// if the postactual node of x coincides with the actual node.
// Returns otherwise a copy of the edge
// from the postactual to the actual node of x.
  Get1 () Any

// Returns (nil, nil), if x is empty. Returns otherwise
// copies of the postactual and of the actual node of x.
  Get2 () (Any, Any)

// Returns (nil, nil, nil), if x is empty or has no edgetype or
// if the postactual node of x coincides with the actual node of x or
// if there is no edge from the postactual node to the actual node of x.
// Returns otherwise a copies of the postactual node, the actual node
// and the edge from the postactual to the actual node of x.
  Get3 () (Any, Any, Any)

// If x is empty or
// if n is not of the nodetype of x, nothing has happened. Otherwise:
// The actual node of x is replaced by n.
  Put (n Any)

// If x is empty or if e has no edgetype or
// if e is not of the edgetype of x or
// if there is no edge from the postactual to the actual node of x,
// nothing has happened. Otherwise:
// The edge from the postactual to the actual node of x is replaced by e.
  Put1 (e Any)

// If x is empty or
// if n or n1 is not of the nodetype of x or
// if the postactual node of x coincides with the actual node,
// nothing had happened. Otherwise:
// The postactual node of x is replaced by n
// and the actual node of x is replaced by n1.
  Put2 (n, n1 Any)

// If x is empty or if e has no edgetype or
// if n or n1 is not of the nodetype of x or
// if e is not of the edgetype of x or
// if the postactual node of x coincides with the actual node or
// if there is no edge from the postactual node to the actual node of x,
// nothing had happened. Otherwise:
// The postactual node of x is replaced by n,
// the actual node of x is repaced by n1 and the edge
// from the postactual node to the actual node is replaced by e.
  Put3 (n, n1, e Any)

// If x is empty, nothing has happened. Otherwise:
// The former actual node of x and
// all its outgoing and incoming edges are deleted.
// If x is now not empty, some other node is now the actual node
// and coincides with the postactual node of x.
// The actual way and the actual subgraph of x are empty. 
  Del ()

// If there was an edge between the postactual and the actual node of x,
// it is now deleted from x.
  Del1 ()

// Returns true, iff x is empty or
// if the postactual node coincides with the actual node of x or
// if there is a way from the postactual to the actual node in x.
  Conn () bool

// Pre: p is defined on nodes.
// Returns true, iff x is empty or
// the postactual node coincides with the actual node of x or
// if p returns true for the actual node and there is a way
// from the postactual node of x to the actual node, that contains
// - apart from the postactual node - only nodes, for which p returns true.
  ConnCond (p Pred) bool

// If x is empty, nothing had happened. Otherwise:
// If there is a way from the postactual to the actual node of x,
// the actual way of x is a shortest such way
// (shortest w.r.t. the sum of the values of its edges,
// hence, if x has no edgetype, w.r.t. their number).
// If there is no way from the postactual to the actual node of x,
// the actual way consists only of the postactual node.
// The actual subgraph of x is the actual way of x.
  Actualize () // TODO name

// Pre: p is defined on nodes.
// If x is empty, nothing had happened. Otherwise:
// If p returns true for the actual node and there is a way
// from the postactual to the actual node of x, that contains
// - apart from the postactual node - only nodes, for which p returns true,
// the actual way of x is a shortest such way
// w.r.t. the sum of the values of its edges
// (hence, if x has no edgetype, w.r.t. their number).
// Otherwise the actual way consists only of the postactual node.
// The actual subgraph of x is the actual way of x.
  ActualizePred (p Pred) // TODO name

// Returns the sum of the values of all edges of x
// (hence, if x has no edgetype, the number of the edges of x).
  Len () uint

// Returns the sum of the values of all edges in the actual subgraph of x
// (hence, if x has no edgetype, the number of the edges in the subgraph).
  LenAct () uint

// Returns 0, if x is empty.
// Returns otherwise the number of the outgoing edges of actual node of x.
  NumLocal () uint // TODO name

// If x is not directed, nothing had happened. Otherwise:
// The directions of all edges of x are reversed.
  Inv ()

// If x is not directed, nothing had happened. Otherwise:
// The directions of all outgoing and incoming edges
// of the actual node of x are reversed.
  InvAct ()

// If x is empty, nothing had happened. Otherwise:
// The actual and the postactual node of x are exchanged.
// The actual way of x consists only of the postactual node of x
// and the actual subgraph of x is the actual way of x.
  Reposition () // TODO name

// If x is empty, nothing had happened. Otherwise:
// The postactual node coincides with the actual node of x,
// where for v true that is the node, that was the former actual node of x,
// and for !v that one, that was the former postactual node of x.
// The actual way of x consists only of this node
// and the actual subgraph of x is the actual way.
  Position (f bool) // TODO name

// Returns true, iff x is empty or the actual node of x
// coincides with the postactual node of x.
  Positioned () bool // TODO name

// If x is empty or if i >= number of nodes outgoing from the actual node
// nothing had happened. Otherwise:
// For v:  The i-th neighbour node of the last node of the actual way
//         of x is appended to it as new last node.
// For !v: The last node of the actual way of x is deleted from it,
//         if it had not only one node (i does not play any role in this case).
// The last node of the actual way of x is the actual node of x and
// the actual subgraph of x is its actual way.
  Step (i uint, f bool)

// The actual node of x is pushed on the nodestack of x.
  Save ()

// If the nodestack of x is empty, nothing had happened. Otherwise:
// The actual node is now the top of the nodestack
// and this node is pulled from the nodestack of x.
  Restore ()

////////////////////////////////////////////////////////////////////////////////
// experimental stuff:

// actualnode.dist = 0.
// A ()

// Returns true, iff actualnode.dist = unendlich.
// B () bool

// actualnode.dist = postactualnode.dist + 1
// actualnode.hinten = postactualnode.
// C ()
////////////////////////////////////////////////////////////////////////////////

// Pre: p is defined on nodes.
// Returns true, if x is empty or
// if p returns true for all nodes of x.
  True (p Pred) bool

// Pre: p is defined on nodes.
// Returns true, iff x is empty or
// if p returns true for all nodes in the actual subgraph of x.
  TrueAct (p Pred) bool

// Pre: p is defined on nodes.
// p is applied to all nodes of x.
// Postactual and actual node of x are the same as before;
// subgraph and nodestack of x are not changed.
  Trav (p Op)

// Pre: p is defined on nodes.
// p is applied to all nodes of x, where
// p is called with 2nd parameter "true", iff
// the corresponding node is contained in the actual subgraph of x.
// Postactual and actual node of x are the same as before;
// subgraph and nodestack of x are not changed.
  TravCond (p CondOp)

// Pre: p is defined on edges.
// If x has no edgetype, nothing had happened. Otherwise:
// p is applied to all edges of x.
// Postactual and actual node of x are the same as before;
// subgraph and nodestack of x are not changed.
  Trav1 (p Op)

// Pre: p is defined on edges.
// If x has no edgetype, nothing had happened. Otherwise:
// p is applied on all edges of x with 2nd parameter "true", iff
// the correspoding edge is contained in the actual subgraph of x.
// Postactual and actual node of x are the same as before;
// subgraph and nodestack of x are not changed.
  Trav1Cond (p CondOp)

// Pre: p is defined on tripels of the form (edge, node, node).
// If x has no edgetype, nothing had happened. Otherwise:
// For all edges e of x with their neighbour nodes n and n1
// (if x is directed, in the direction n -e-> n1)
// p is applied to the tripel (e, n, n1).
// Postactual and actual node of x are the same as before;
// subgraph and nodestack of x are not changed.
  Trav3 (B Op3)

// Pre: p is defined on tripels of the form (edge, node, node).
// If x has no edgetype, nothing had happened. Otherwise:
// For all edges e of x with their neighbour nodes n and n1
// p is applied to the triple (e, n, n1), where
// p is called with 4th parameter "true", iff
// e is contained in the actual subgraph of x.
// Postactual and actual node of x are the same as before;
// subgraph and nodestack of x are not changed.
  Trav3Cond (p CondOp3)

// Returns true, iff there are no cycles in x.
  Acyclic () bool

// Returns false, if x is not totally connected.
// Returns otherwise true, iff there is an Euler way or circle in x.
  Eulerian () bool

// If x is empty, nothing has happened. Otherwise:
// The following equivalence relation is defined on x:
// Two nodes e and e1 of x are equivalent, iff
// there is a way in x from e to e1 and vice versa (hence
// the set of equivalence classes is a directed graph without cycles).
  Isolate () // TODO name

// The actual subgraph of x consists of exactly those nodes, that are
// equivalent to the actual node and of exactly all edges between them.
// The actual way of x is now empty.
  IsolateAct () // TODO name

// Returns true, iff x is not empty and
// if the actual and the postactual node of x are equivalent,
// i.e. for both of them there is a way in x to the other one.
  Equiv () bool

// If x is directed, nothing has happened. Otherwise:
// The actual subgraph of x is a minimal spanning tree in
// the connected component, that contains the postactual node
// (minimal w.r.t. the values of the sum of its edges;
// hence, if x has no edgetype, w.r.t. the number of its nodes)
// The actual way is not changed. */
  Minimize () // TODO name

// If x is empty or undirected or
// if x is directed and has cycles, nothing has happened. Otherwise:
// The nodes of x are ordered s.t. at each subsequent traversal of x
// each node with outgoing edges is always handled before the nodes,
// at which those edges come in.
  Sort ()

// Pre: p is defined on nodes and p3 on tripels of the form (edge, node, node).
// p and p3 are the demofunctions for the Writing of nodes and edges of x.
  Install (p CondOp, p3 CondOp3)

// The demofunction for d is switched on, iff s[d] == true.
  Set (d Demo)
}
