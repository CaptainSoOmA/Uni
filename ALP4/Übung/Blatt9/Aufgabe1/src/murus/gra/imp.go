package gra

// (c) Christian Maurer   v. 130119 - license see murus.go

// >>> alpha-version; still quite a lot of things TODO

import (
  "sort"
  . "murus/obj"; "murus/ker"; "murus/str"; "murus/rand"
  "murus/kbd"; "murus/errh"
  "murus/pseq"; "murus/adj"
)

// CLR  = Cormen, Leiserson, Rivest        1990
// CLRS = Cormen, Leiserson, Rivest, Stein 2001

/*    node                                                                   node
                       neighbour                        neighbour
  [-----------]                                                          [-----------]
  [  content  ]           /--------------------------------------------->[  content  ]
  [-----------]<--------/----------------------------------------\       [-----------]
  [   nbPtr --]-------/---\                                  /-----\-----[-- nbPtr   ]
  [-----------]     /      |              edge              |        \   [-----------]
  [inSubgraph ]   /        V                                V         |  [inSubgraph ]
  [-----------]  |   [-----------]    [----------]    [-----------]   |  [-----------]
  [  marked   ]  |   [  edgePtr--]--->[  attrib  ]<---[--edgePtr  ]   |  [  marked   ]
  [-----|-----]  |   [-----------]    [----------]    [-----------]   |  [-----|-----]
  [dist |time ]<-----[-- from    ]<---[--nbPtr0  ]    [   from ---]----->[dist |time ]
  [-----|-----]  |   [-----------]    [----------]    [-----------]   |  [-----|-----]
  [predecessor]   \--[--- to     ]    [  nbPtr1--]--->[    to ----]--/   [predecessor]
  [-----------]      [-----------]    [----------]    [-----------]      [-----------]
  [    repr   ]      [  forward  ]    [inSubgraph]    [  forward  ]      [    repr   ]
  [-----------]      [-----------]    [----------]    [-----------]      [-----------]
  [   nextN --]->    [  nextNb---]->  [  nextE --]->  [  nextNb --]->    [   nextN --]->
  [-----------]      [-----------]    [----------]    [-----------]      [-----------]
<-[-- prevN   ]    <-[-- prevNb  ]  <-[-- prevE  ]  <-[-- prevNb  ]    <-[-- prevN   ]
  [-----------]      [-----------]    [----------]    [-----------]      [-----------]

The nodes of a graph are represented by structs,
whose field "content" represents the "real" node.
All nodes are connected in a doubly linked list with anchor cell,
that can be traversed to execute some operation on all nodes of the graph.

The edges are also represented by structs,
whose field "attrib" is either nil (with value 1)
or carries a variable of type Valuator.
Also all edges are connected in a doubly linked list with anchor cell.

For a node n one finds all outgoing and incoming edges
with the help of a further doubly linked ringlist of neighbour(hoodrelation)s
  nb1 = n.nbPtr, nb2 = n.nbPtr.nextNb, nb3 = n.nbPtr.nextNb.nextNb etc.
by the links outgoing from the nbi (i = 1, 2, 3, ...)
  nb1.edgePtr, nb2.edgePtr, nb3.edgePtr etc.
In directed graphs the edges outgoing from a node are exactly those ones
in the neighbourlist, for which forward == true.

For an edge e one finds its two nodes by the links
  e.nbPtr0.from = e.nbPtr1.to und e.nbPtr0.to = e.nbPtr1.from.

Semantics of some variables, that are "hidden" in fields of nodeAnchor:
  nodeAnchor.time0: in that the "time" is incremented for each search step
  nodeAnchor.marked: (after call of search) == true <=> graph has no cycles
*/

const (
  pack = "gra"
  inf = uint(1<<32 - 1)
)
type (
  node struct {
      content Any
        nbPtr *neighbour
   inSubgraph,      // characteristic function of the nodes in the actual subgraph
       marked bool  // for the development of design patterns by clients
         dist,      // for breadth first search/Dijkstra and use in En/Decode
 time0, time1 uint  // for applications of depth first search
  predecessor,      // for back pointers in depth first search and in ways
         repr,      // for the computation of connected components
 nextN, prevN *node
              }

  nodeSet []*node // to be able to apply sort.Sort to []*node

  nodeCell struct {
             nPtr *node
             next *nodeCell
                  }

  edge struct {
       attrib Any
       nbPtr0,
       nbPtr1 *neighbour
   inSubgraph bool // characteristic function of the nodes in the actual subgraph
 nextE, prevE *edge
              }

  edgeSet []*edge // to be able to apply sort.Sort to []*node

  neighbour struct {
           edgePtr *edge
          from, to *node
           forward bool
    nextNb, prevNb *neighbour
                   }

  Imp struct {
        name string
    directed bool
      nNodes,
      nEdges uint
  nodeAnchor,
  postactual,
      actual *node
  edgeAnchor *edge
         path []*node
    eulerPath []*neighbour
       ncPtr *nodeCell
        demo Demoset
       write CondOp
      write3 CondOp3
             }
)


func newNode (a Any) *node {
//
  n:= new (node)
  n.content = Clone (a)
  n.time1 = inf // for applications of depth first search
  n.dist = inf
  n.repr = n
  n.nextN, n.prevN = n, n
  return n
}


func (n *node) Clone () *node {
//
  return n
}


func insert (s []*node, n *node, i uint) []*node {
//
  l:= uint(len (s))
  if i > l { i = l }
  s1:= make ([]*node, l + 1)
  copy (s1[:i], s[:i])
  s1[i] = n
  copy (s1[i+1:], s[i:])
  return s1
}


func contains (s []*node, n *node) bool {
//
  l:= uint(len (s))
  c:= l
  for i, a:= range s {
    if a == n {
      c = uint(i)
      break
    }
  }
  return c < l
}


func exists (s []*node, p Pred) (*node, bool) {
//
  for _, a:= range s {
    if p (a) {
      return a, true
    }
  }
  return nil, false
}


func remove (s []*node, i uint) []*node {
//
  l:= uint(len (s))
  if l == 0 { return nil }
  if i >= l { return s }
  s1:= make ([]*node, l - 1)
  copy (s1[:i], s[:i])
  copy (s1[i:], s[i+1:])
  return s1
}


func (s nodeSet) Clone () nodeSet { // just for fun (proof of concept)
//
  t:= make (nodeSet, len (s))
  for i, x:= range (s) {
    t[i] = x
  }
  return t
}


func (s nodeSet) Len () int {
//
  return len (s)
}


func (s nodeSet) Swap (i, j int) {
//
  s[i], s[j] = s[j], s[i]
}


func (s nodeSet) Less (i, j int) bool {
//
  if s[i].dist == s[j].dist {
    if s[i] == s[j] {
      return false
    }
    return i < j
  }
  return s[i].dist < s[j].dist
}


func newEdge () *edge {
//
  e:= new (edge)
  e.nextE, e.prevE = e, e
  return e
}


func (s edgeSet) Len () int {
//
  return len (s)
}


func (s edgeSet) Swap (i, j int) {
//
  s[i], s[j] = s[j], s[i]
}


func (s edgeSet) Less (i, j int) bool {
//
  return Val (s[i]) < Val (s[j])
}


func newNeighbour (e *edge, n, n1 *node, f bool) *neighbour {
//
  nb:= new (neighbour)
//  errh.Error ("neuer Nachbar", uint(uintptr(unsafe.Pointer(nb))))
  nb.edgePtr = e
  nb.from, nb.to = n, n1
  nb.forward = f
  nb.nextNb, nb.prevNb = nb, nb
  return nb
}


func existsnb (s []*neighbour, p Pred) (*neighbour, bool) {
//
  for _, a:= range s {
    if p (a) {
      return a, true
    }
  }
  return nil, false
}


func New (d bool, n, e Any) *Imp {
//
  if n == nil { PanicIfNotOk (n) }
  x:= new (Imp)
  x.directed = d
  x.nodeAnchor, x.edgeAnchor = newNode (n), newEdge ()
  if e == nil {
    x.edgeAnchor.attrib = nil
  } else {
    if _, ok:= e.(Valuator); ! ok { TypePanic() }
    x.edgeAnchor.attrib = Clone (e)
  }
  x.postactual, x.actual = x.nodeAnchor, x.nodeAnchor
  x.write, x.write3 = CondNull1, CondNull3
  return x
}


func (x *Imp) Name (s string) {
//
  x.name = s
  str.Move (&x.name, true)
  str.RemSpaces (&x.name)
  n:= pseq.Length (x.name)
  if n > 0 {
    buf:= make ([]byte, n)
    f:= pseq.New (buf)
    f.Name (x.name)
    buf = f.Get ().([]byte)
    f.Terminate ()
    x.Decode (buf)
  }
}


func (x *Imp) Rename (s string) {
//
  x.name = s
  str.Move (&x.name, true)
  str.RemSpaces (&x.name)
// rest of implementation TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO
  n:= pseq.Length (x.name)
  if n > 0 {
    buf:= make ([]byte, n)
    f:= pseq.New (buf)
    f.Rename (x.name)
//    buf = f.Get ().([]byte)
    f.Terminate ()
//    x.Decode (buf)
  }
}


func (x *Imp) Terminate () {
//
  if ! str.Empty (x.name) {
    buf:= x.Encode ()
    f:= pseq.New (buf)
    f.Name (x.name)
    f.Clr ()
    f.Put (buf)
    f.Terminate ()
  }
//  x.Clr ()
}


func (x *Imp) Directed () bool {
//
  return x.directed
}


func (x *Imp) Empty () bool {
//
  return x.nodeAnchor.nextN == x.nodeAnchor
}


func delEdge (e *edge) {
//
  e = e.nextE
  e.prevE.nextE, e.nextE.prevE = e.nextE, e.prevE
  e.nbPtr0.prevNb.nextNb, e.nbPtr0.nextNb.prevNb = e.nbPtr0.nextNb, e.nbPtr0.prevNb // bug
  e.nbPtr1.prevNb.nextNb, e.nbPtr1.nextNb.prevNb = e.nbPtr1.nextNb, e.nbPtr1.prevNb
}


func delNode (n *node) {
//
  N:= n.nbPtr.nextNb
  for N != n.nbPtr {
    N = n.nbPtr
    N.to.predecessor = nil
    n.nbPtr = n.nbPtr.nextNb
  }
  n.prevN.nextN, n.nextN.prevN = n.nextN, n.prevN
  n = n.nextN
}


func (x *Imp) Clr () {
//
  for x.edgeAnchor.nextE != x.edgeAnchor {
    delEdge (x.edgeAnchor.nextE) // bug
  }
  x.nEdges = 0
  for x.nodeAnchor.nextN != x.nodeAnchor {
    delNode (x.nodeAnchor.nextN)
  }
  x.nNodes = 0
  x.postactual, x.actual = x.nodeAnchor, x.nodeAnchor
  x.path = nil
  x.eulerPath = nil
}


func (x *Imp) nE (n *node) uint {
//
  c:= uint(0)
  for nb:= n.nbPtr; nb.nextNb != n.nbPtr; nb = nb.nextNb {
    c ++
  }
  return c
}


func (x *Imp) Eq (Y Object) bool { // eklige Komplexität: x.nNodes * y.nNodes * ... * ...
//
  y, ok:= Y.(*Imp)
  if ! ok { return false }
  if x.nNodes != y.nNodes || x.nEdges != y.nEdges ||
     ! TypeEq (x.nodeAnchor.content, y.nodeAnchor.content) ||
     ! TypeEq (x.edgeAnchor.attrib, y.edgeAnchor.attrib) {
    return false
  }
  ya:= y.actual // save
  eq:= true
  loop:
  for xn:= x.nodeAnchor.nextN; xn != x.nodeAnchor; xn = xn.nextN {
    if ! y.Ex (xn.content) {
      eq = false
      break
    }
    yn:= y.actual // y.actual was changed
    if x.nE (xn) != y.nE (yn) {
      eq = false
      break
    }
    for xnb:= xn.nbPtr; xnb.nextNb != xn.nbPtr; xnb = xnb.nextNb {
      for ynb:= yn.nbPtr; ynb.nextNb != yn.nbPtr; ynb = ynb.nextNb {
        if ynb.to == xnb.to {
          aa:= true
          if x.edgeAnchor.attrib != nil {
            aa = Eq (xnb.edgePtr.attrib, ynb.edgePtr.attrib)
          }
          if aa {
            break // next xnb
          } else {
            eq = false
            break loop
          }
        }
      }
    }
  }
  y.actual = ya // restore
  return eq
}


func (x *Imp) Less (Y Object) bool {
//
  return false
}


func (x *Imp) Copy (Y Object) {
//
  y, ok:= Y.(*Imp)
  if ! ok { return }
  x.Decode (y.Encode ())
}


func (x *Imp) Clone () Object {
//
  var y *Imp = New (x.directed, x.nodeAnchor.content, x.edgeAnchor.attrib)
  y.Decode (x.Encode ())
  return y
}


func (x *Imp) Num () uint {
//
  return x.nNodes
}


func (x *Imp) NumAct () uint {
//
  c:= uint(0)
  for n:= x.nodeAnchor.nextN; n != x.nodeAnchor; n = n.nextN {
    if n.inSubgraph {
      c ++
    }
  }
  return c
}


func (x *Imp) NumPred (p Pred) uint {
//
  c:= uint(0)
  for n:= x.nodeAnchor.nextN; n != x.nodeAnchor; n = n.nextN {
    if p (n.content) {
      c ++
    }
  }
  return c
}


func (x *Imp) Num1 () uint {
//
  return x.nEdges
}


func (x *Imp) insertedNode ( /* n *node, */  a Any) *node {
//
  n:= newNode (a)
// TODO is the following sensefull ?
  n.nbPtr = newNeighbour (nil, n, nil, false) // was set to nil by newNode
// include by pointers in x.nodeAnchor:
  n.nextN, n.prevN = x.nodeAnchor, x.nodeAnchor.prevN // n, n
  n.prevN.nextN = n
  x.nodeAnchor.prevN = n
//
  return n
}


func (x *Imp) Ins (a Any) {
//
  if ! TypeEq (a, x.nodeAnchor.content) {
    return
  }
  if x.Ex (a) { // actual is set
    return
  }
  n:= x.insertedNode (a)
  x.nNodes ++
  if x.nNodes == 1 {
    x.postactual = n
  } else {
    x.postactual = x.actual
  }
  x.actual = n
}


func (x *Imp) Edge () {
//
  x.Edge1 (nil)
}


// Pre: n and n1 are nodes in the same graph.
// Returns true, iff there is no edge from n to n1;
// returns otherwise the corresponding pointer.
func connection (n, n1 *node) *edge {
//
  nb:= n.nbPtr.nextNb
  for {
    if nb == n.nbPtr {
      break
    }
    if nb.forward && nb.to == n1 {
      return nb.edgePtr
    }
    nb = nb.nextNb
  }
  return nil
}


// Pre: nb.from == e.
// nb is appended in n.nbPtr
func insertNeighbour (nb *neighbour, n *node) {
//
  nb.nextNb, nb.prevNb = n.nbPtr, n.nbPtr.prevNb
  nb.prevNb.nextNb = nb
  n.nbPtr.prevNb = nb
}


// TODO Spec
func (x *Imp) insertEdge (a Any) {
//
  if ! TypeEq (a, x.edgeAnchor.attrib) {
    return
  }
  e:= newEdge ()
  if a == nil {
    e.attrib = nil
  } else {
    e.attrib = Clone (a)
  }
  e.nbPtr0 = newNeighbour (e, x.postactual, x.actual, true)
  insertNeighbour (e.nbPtr0, x.postactual)
  e.nbPtr1 = newNeighbour (e, x.actual, x.postactual, ! x.directed)
  insertNeighbour (e.nbPtr1, x.actual)
  e.nextE, e.prevE = x.edgeAnchor, x.edgeAnchor.prevE
  e.prevE.nextE = e
  x.edgeAnchor.prevE = e
}


func (x *Imp) Edge1 (a Any) {
//
  if x.Empty () || ! TypeEq (a, x.edgeAnchor.attrib) ||
    x.postactual == x.actual {
    return
  }
  // simple case: actual and postactual are not yet adjacent:
  if connection (x.postactual, x.actual) == nil &&
    connection (x.actual, x.postactual) == nil {
    x.insertEdge (a)
    x.nEdges ++
    return
  }
  // otherwise: an existing edge must not be cleared:
  if a == nil { return }
// if there is an edge from postactual to actual, it is looked for:
  nb:= x.postactual.nbPtr.nextNb
  for nb.to != x.actual {
    nb = nb.nextNb
    if nb == x.postactual.nbPtr { ker.Stop (pack, 1) } // not found, contradiction
  }
// and its attrib is replaced:
  nb.edgePtr.attrib = Clone (a)
  nb.forward = true
// in the directed case the edge goes from postactual to actual,
// but not the other way:
  if x.directed {
    nb = x.actual.nbPtr.nextNb
    for nb.to != x.postactual {
      nb = nb.nextNb
      if nb == x.actual.nbPtr { ker.Stop (pack, 2) }
    }
    nb.forward = false
  }
}


func (x *Imp) Edge2 (a, a1 Any) {
//
  x.Edge3 (a, nil, a1)
}


func (x *Imp) Edge3 (a, b, a1 Any) {
//
  if x.Empty () || // x.edgeAnchor.attrib == nil ||
    Eq (a, a1) ||
    ! TypeEq (a, a1) ||
    ! TypeEq (a, x.nodeAnchor.content) ||
    ! TypeEq (b, x.edgeAnchor.attrib) ||
    false {
    return
  }
  if n, ok:= x.found (a); ! ok {
    return
  } else {
    x.postactual = n
  }
  if n1, ok:= x.found (a1); ! ok {
    return
  } else {
    x.actual = n1
  }
  if x.postactual == x.actual ||
    connection (x.postactual, x.actual) != nil {
    return
  }
  x.Edge1 (b)
}


func (x *Imp) Matrix () *adj.Imp {
//
  if x.edgeAnchor.attrib != nil { return nil }
  m:= x.Num ()
  any:= make ([]Any, m)
  i:= uint(0)
  for n:= x.nodeAnchor.nextN; n != x.nodeAnchor; n = n.nextN {
    any[i] = Clone (n.content)
    n.time0 = i
    i++
  }
  mat:= adj.New (any, x.edgeAnchor.attrib)
  i = 0
  for n:= x.nodeAnchor.nextN; n != x.nodeAnchor; n = n.nextN {
    for nb:= n.nbPtr.nextNb; nb != n.nbPtr; nb = nb.nextNb {
      if nb.forward {
        mat.Put (i, nb.to.time0, nb.edgePtr.attrib)
      }
    }
    i++
  }
  return mat
}


func (x *Imp) Def (m *adj.Imp) {
//
  x = New (m.Symmetric (), m.Node (0), nil)
  n:= m.Num ()
  for i:= uint(0); i < n; i++ {
    x.Ins (m.Node (i))
  }
  for i:= uint(0); i < n; i++ {
    for k:= uint(0); k < n; k++ {
      x.Edge3 (m.Node (i), m.Edge (i, k), m.Node (k))
    }
  }
}


func (x *Imp) Edged () bool {
//
  return connection (x.postactual, x.actual) != nil
}


func (x *Imp) EdgedInv () bool {
//
  return connection (x.actual, x.postactual) != nil
}


// Returns (nil, false), iff a there is no node in x with content a;
// returns otherwise (n, true), where n is the pointer to that node
// (which is unique because of effect of Ins).
func (x *Imp) found (a Any) (*node, bool) {
//
  for n:= x.nodeAnchor.nextN; n != x.nodeAnchor; n = n.nextN {
    if Eq (n.content, a) {
// errh.Error ("gra.found found node with content", a.(uint))
      return n, true
    }
  }
  return nil, false
}


func (x *Imp) Ex (a Any) bool {
//
  if n, ok:= x.found (a); ok {
    x.actual = n
    return true
  }
  return false
}


func (x *Imp) Ex2 (a, a1 Any) bool {
//
  if Eq (a, a1) {
// errh.Error ("#", 20)
    return false
  }
  if n, ok:= x.found (a); ok {
// errh.Error ("Ex2 found from", a.(uint))
    if n1, ok1:= x.found (a1); ok1 {
// errh.Error ("Ex2 found to", a1.(uint))
      x.postactual = n
      x.actual = n1
      return true
    }
  }
// errh.Error2 ("Ex2 ! found from", a.(uint), "to", a1.(uint))
  return false
}


// Returns true, iff there is no node in x, for which p returns true,
// Returns otherwise a pointer to such a node.
func (x *Imp) foundPred (p Pred) (*node, bool) {
//
  n:= x.nodeAnchor.nextN
  for n != x.nodeAnchor {
    if /* e != x.nodeAnchor.predecessor && */ p (n.content) {
      return n, true
    } else {
      n = n.nextN
    }
  }
  return nil, false
}


func (x *Imp) ExPred (p Pred) bool {
//
//  x.nodeAnchor.predecessor = nil
  if n, ok:= x.foundPred (p); ok {
    x.actual = n
    return true
  }
  return false
}


func (x *Imp) ExPred1 (p Pred) bool {
//
  e:= x.edgeAnchor.nextE
  for e != x.edgeAnchor {
    if p (e.attrib) {
      if e.nbPtr0.forward && e.nbPtr1.forward {
        x.postactual = e.nbPtr0.from
        x.actual = e.nbPtr1.from
      } else {
        x.postactual = e.nbPtr1.from
        x.actual = e.nbPtr0.from
      }
      return true
    }
    e = e.nextE
  }
  return false
}


func (x *Imp) ExPred2 (p, p1 Pred) bool {
//
  if n, ok:= x.foundPred (p); ok {
    if n1, ok1:= x.foundPred (p1); ok1 {
      if n == n1 {
        tmp:= x.nodeAnchor.nextN
        x.nodeAnchor.nextN = n
        n1, ok1 = x.foundPred (p1) // n1 != n
        x.nodeAnchor.nextN = tmp
        if ! ok1 {
          return false
        }
      }
      x.postactual = n
      x.actual = n1
      return true
    }
  }
  return false
}


func (x *Imp) Get () Any {
//
  if x.actual == x.nodeAnchor { return nil }
  return Clone (x.actual.content)
}


func (x *Imp) Get2 () (Any, Any) {
//
  if x.actual == x.nodeAnchor { return nil, nil }
  return Clone (x.postactual.content), Clone (x.actual.content)
}


func (x *Imp) Get1 () Any {
//
  if x.actual == x.nodeAnchor { return nil }
  if x.actual == x.postactual { return nil }
  nb:= x.postactual.nbPtr.nextNb
  for {
    if nb == x.postactual.nbPtr { break }
    if nb.forward && nb.to == x.actual {
      break
    }
    nb = nb.nextNb
  }
  return Clone (nb.edgePtr.attrib)
}


func (x *Imp) Get3 () (Any, Any, Any) {
//
  if x.actual == x.nodeAnchor { return nil, nil, nil }
  if x.actual == x.postactual { return nil, nil, nil }
  nb:= x.postactual.nbPtr.nextNb
  for {
    if nb == x.postactual.nbPtr { ker.Stop (pack, 3) }
    if nb.forward && nb.to == x.actual {
      if nb.from != x.postactual { ker.Stop (pack, 4) }
      break
    }
    nb = nb.nextNb
  }
  return Clone (nb.from.content), Clone (nb.to.content), Clone (nb.edgePtr.attrib)
}


func (x *Imp) Put (a Any) {
//
  if x.nodeAnchor == x.nodeAnchor.nextN { return }
  x.actual.content = Clone (a)
}


func (x *Imp) Put1 (a Any) {
//
  if x.nodeAnchor == x.nodeAnchor.nextN { return }
  if x.postactual == x.actual { return }
  x.postactual.nbPtr.edgePtr.attrib = Clone (a)
}


func (x *Imp) Put2 (a, a1 Any) {
//
  if x.nodeAnchor == x.nodeAnchor.nextN { return }
  if x.postactual == x.actual { return }
  x.postactual.content = Clone (a)
  x.actual.content = Clone (a1)
}


func (x *Imp) Put3 (a, a1, a2 Any) {
//
  if x.nodeAnchor == x.nodeAnchor.nextN { return }
  if x.postactual == x.actual { return }
  x.postactual.content = Clone (a)
  x.actual.content = Clone (a1)
  x.postactual.nbPtr.edgePtr.attrib = Clone (a2)
}


func (x *Imp) Del () {
//
  if x.nodeAnchor == x.nodeAnchor.nextN { return }
  if x.actual == x.nodeAnchor { return }
// delete all edges and their neighbour lists
  for x.actual.nbPtr.nextNb != x.actual.nbPtr {
    delEdge (x.actual.nbPtr.nextNb.edgePtr)
    x.nEdges --
  }
  x.path = nil
  x.clearSubgraph ()
  n:= x.actual
  delNode (x.actual)
  x.nNodes --
  x.actual = n.prevN
  if x.actual == x.nodeAnchor {
    x.actual = x.nodeAnchor.nextN
  }
  x.postactual = x.actual
}


func (x *Imp) Del1 () {
//
  if x.postactual == x.nodeAnchor || x.postactual == x.actual {
    return
  }
  nb:= x.postactual.nbPtr.nextNb
  for nb.to != x.actual {
    if nb == x.postactual.nbPtr {
      return // actual no neighbour of postactual
    } else {
      nb = nb.nextNb
    }
  }
  delEdge (nb.edgePtr)
  x.nEdges --
}


func wait () {
//
  kbd.Wait (true)
}


func (x *Imp) preDepth () {
//
  x.nodeAnchor.time0 = 0
  n:= x.nodeAnchor.nextN
  for n != x.nodeAnchor {
    n.time0, n.time1 = 0, 0
    n.predecessor = nil
    n.repr = n
    n = n.nextN
  }
}


// For all nodes n, that are accessible from n0 by a way, n.repr == n0.
// nodeAnchor.marked == true, if x has no cycles.
func (x *Imp) search (n0, n *node, p Pred) {
//
  x.nodeAnchor.time0 ++
  n.time0 = x.nodeAnchor.time0
  n.repr = n0
  nb:= n.nbPtr.nextNb
  if x.demo [Depth] {
    x.write (n.content, true)
    wait()
  }
  for nb != n.nbPtr {
    if nb.forward && nb.to != n.predecessor && p (nb.to.content) {
      if nb.to.time0 == 0 {
        if x.demo [Depth] {
          x.write3 (nb.edgePtr.attrib, n.content, nb.to.content, true)
        }
        nb.to.predecessor = n
        x.search (n0, nb.to, p)
        if x.demo [Depth] {
          x.write3 (nb.edgePtr.attrib, n.content, nb.to.content, false)
          wait()
        }
      } else if nb.to.time1 == 0 {
        x.nodeAnchor.marked = false // found cycle
        if x.demo [Cycle] { // also x.demo [Depth], see Set
          x.write3 (nb.edgePtr.attrib, n.content, nb.to.content, true)
//          errh.Error ("Kreis gefunden", 0)
          x.write3 (nb.edgePtr.attrib, n.content, nb.to.content, false)
          wait()
        }
      }
    }
    nb = nb.nextNb
  }
  x.nodeAnchor.time0 ++
  n.time1 = x.nodeAnchor.time0
  if x.demo [Depth] {
    x.write (n.content, false)
  }
}


func (x *Imp) Conn () bool {
//
  return x.ConnCond (True)
}


func (x *Imp) ConnCond (p Pred) bool {
//
  if x.nodeAnchor == x.nodeAnchor.nextN { return true }
  if x.postactual == x.actual { return true }
  x.preDepth ()
  x.search (x.postactual, x.postactual, p)
  return x.actual.repr == x.postactual
  return x.actual.time0 > 0 // Alternative
}


func (x *Imp) preBreadth () {
//
  n:= x.nodeAnchor.nextN
  for n != x.nodeAnchor {
    n.dist = inf
    n.predecessor = nil
    n = n.nextN
  }
  x.postactual.dist = 0
}


// Lit.: CLR 23.2, CLRS 22.2
// TODO spec
func (x *Imp) breadthfirstSearch (p Pred) {
//
  var qu []*node
  qu = append (qu, x.postactual)
  for len (qu) > 0 {
    n:= qu[0]
    if len (qu) == 1 {
      qu = nil
    } else {
      qu = qu [1:]
    }
    nb:= n.nbPtr.nextNb
    for nb != n.nbPtr {
      if nb.forward && nb.to.dist == inf && p (nb.to.content) {
        if x.demo [Breadth] {
          var nb1 *neighbour
          if nb.to.predecessor != nil {
            nb1 = nb.to.predecessor.nbPtr.nextNb
            for nb1.from != nb.to.predecessor {
              nb1 = nb1.nextNb
              if nb1.nextNb == nb1 { ker.Stop (pack, 5) }
            }
            x.write3 (nb1.edgePtr.attrib, nb.to.predecessor.content, nb.to.content, false)
            x.write (nb.to.content, true)
          }
          x.write3 (nb1.edgePtr.attrib, n.content, nb.to.content, false)
          x.write (nb.to.content, true)
          wait()
        }
        nb.to.dist = n.dist + 1
        nb.to.predecessor = n
        qu = append (qu, nb.to)
      }
      nb = nb.nextNb
    }
  }
}


// Algorithm of Dijkstra, Lit.: CLR 25.1-2, CLRS 24.2-3
// Pre: dist == inf, predecessor == nil for all nodes.
// TODO spec
func (x *Imp) searchShortestPath (p Pred) {
//
  n:= x.nodeAnchor.nextN
  set:= make (nodeSet, x.nNodes)
  for i, n:= 0, x.nodeAnchor.nextN; n != x.nodeAnchor; i, n = i+1, n.nextN {
    set[i] = n
  }
  sort.Sort (set)
  for len (set) > 0 {
    n = set[0]
    if len (set) == 1 {
      set = nil
    } else {
      set = set[1:]
    }
    nb:= n.nbPtr.nextNb
    for nb != n.nbPtr {
      if nb.forward && nb.to != n.predecessor && p (nb.to.content) {
        var d uint
        if n.dist == inf {
          d = inf
        } else {
          d = n.dist + Val (nb.edgePtr.attrib)
        }
        if d < nb.to.dist {
          if x.demo [Breadth] {
            var nb1 *neighbour
            if nb.to.predecessor != nil {
              nb1 = nb.to.predecessor.nbPtr.nextNb
              for nb1.from != nb.to.predecessor {
                nb1 = nb1.nextNb
                if nb1.nextNb == nb1 { ker.Stop (pack, 6) }
              }
              x.write3 (nb1.edgePtr.attrib, nb.to.predecessor.content, nb.to.content, false)
              x.write (nb.to.content, false)
            }
            x.write3 (nb.edgePtr.attrib, n.content, nb.to.content, true)
            x.write (nb.to.content, true)
            wait ()
          }
          nb.to.dist = d
          nb.to.predecessor = n
// put the changed nb.to into the right position in set:
          sort.Sort (set)
        }
      }
      nb = nb.nextNb
    }
  }
}


func (x *Imp) defineSubgraph (n *node) {
//
  n1:= n
  for n1 != x.postactual {
    if n1.predecessor == nil {
      return
    }
    n1 = n1.predecessor
  }
  for {
    n.inSubgraph = true
    if n == x.postactual { return }
    nb:= n.nbPtr.nextNb
    for nb.to != n.predecessor {
      nb = nb.nextNb
      if nb == n.nbPtr { ker.Stop (pack, 7) }
    }
    nb.edgePtr.inSubgraph = true
    n = n.predecessor
  }
}


func (x *Imp) Actualize () {
//
  x.ActualizePred (True)
}


func (x *Imp) clearSubgraph () {
//
  n:= x.nodeAnchor.nextN
  for n != x.nodeAnchor {
    n.inSubgraph = false
    n = n.nextN
  }
  e:= x.edgeAnchor.nextE
  for e != x.edgeAnchor {
    e.inSubgraph = false
    e = e.nextE
  }
}


func (x *Imp) ActualizePred (p Pred) {
//
  n:= x.nodeAnchor.nextN
  if n == x.nodeAnchor {
errh.Error ("#", 10)
    return
  }
  if ! p (x.actual.content) {
// errh.Error ("#", 11)
    return
  }
  x.clearSubgraph ()
  if ! x.ConnCond (p) {
// errh.Error ("#", 12)
    return
  }
  x.preBreadth ()
  if x.edgeAnchor.attrib == nil {
    x.breadthfirstSearch (p)
  } else {
// errh.Error ("shortest Path", 0)
    x.searchShortestPath (p)
  }
  x.path = nil
  n = x.actual
  for n != nil {
    x.path = insert (x.path, n, 0)
    n = n.predecessor
  }
  x.defineSubgraph (x.actual)
}


func (x *Imp) Len () uint {
//
  l:= uint(0)
  if x.nodeAnchor == x.nodeAnchor.nextN {
    return l
  }
  e:= x.edgeAnchor.nextE
  for e != x.edgeAnchor {
    l += Val (e.attrib)
    e = e.nextE
  }
  return l
}


func (x *Imp) LenAct () uint {
//
  l:= uint(0)
  if x.nodeAnchor == x.nodeAnchor.nextN {
    return l
  }
  e:= x.edgeAnchor.nextE
  for e != x.edgeAnchor {
    if e.inSubgraph {
      l += Val (e.attrib)
    }
    e = e.nextE
  }
  return l
}


func (x *Imp) NumLocal () uint {
//
  c:= uint(0)
  nb:= x.actual.nbPtr.nextNb
  for nb != x.actual.nbPtr {
    if nb.forward {
      c ++
    }
    nb = nb.nextNb
  }
  return c
}


func (x *Imp) Inv () {
//
  if x.directed {
    e:= x.edgeAnchor.nextE
    for e != x.edgeAnchor {
      e.nbPtr0.forward = ! e.nbPtr0.forward
      e.nbPtr1.forward = ! e.nbPtr1.forward
      e = e.nextE
    }
  }
}


func (x *Imp) InvAct () {
//
  if x.actual != x.nodeAnchor {
    if x.directed {
      nb:= x.actual.nbPtr.nextNb
      for nb != x.actual.nbPtr {
        nb.edgePtr.nbPtr0.forward = ! nb.edgePtr.nbPtr0.forward
        nb.edgePtr.nbPtr1.forward = ! nb.edgePtr.nbPtr1.forward
        nb = nb.nextNb
      }
    }
  }
}


func (x *Imp) Reposition () {
//
  x.clearSubgraph ()
  n:= x.actual
  x.actual = x.postactual
  x.postactual = n
  x.postactual.inSubgraph = true
//  x.path = make ([]*node, 1)
//  x.path[0] = x.postactual
  x.path = nil
  x.path = insert (x.path, x.postactual, 0)
}


func (x *Imp) Position0 (postactual2actual bool) {
//
  x.clearSubgraph ()
  if postactual2actual {
    x.postactual = x.actual
  } else {
    x.actual = x.postactual
  }
  x.path = nil
}


func (x *Imp) Position (postactual2actual bool) {
//
  x.Position0 (postactual2actual)
  x.postactual.inSubgraph = true
  x.path = append (x.path, x.postactual)
}


func (x *Imp) Positioned () bool {
//
  if x.nodeAnchor == x.nodeAnchor.nextN {
    return true
  }
  return x.actual == x.postactual
}


func (x *Imp) Step (i uint, vorwaerts bool) {
//
  if x.nodeAnchor == x.nodeAnchor.nextN {
    return
  }
  if vorwaerts {
    if i >= x.NumLocal () { return }
    if x.path == nil {
      x.postactual.inSubgraph = true
      x.path = make ([]*node, 1)
      x.path[0] = x.postactual
      x.actual = x.postactual
    } else {
      P:= x.path[0]
      if x.postactual != P { ker.Stop (pack, 8) }
    }
    c:= uint(len (x.path))
    n:= x.path[c - 1]
    if x.actual != n { ker.Stop (pack, 9) }
    nb:= x.actual.nbPtr.nextNb
    for {
      if nb.forward {
        if i == 0 {
          break
        } else {
          i --
        }
      }
      nb = nb.nextNb
    }
    nb.edgePtr.inSubgraph = true
    nb.to.inSubgraph = true
    x.actual = nb.to
    x.path = append (x.path, x.actual) // insert (x.path, x.actual, c)
  } else {
    c:= uint(len (x.path))
    if c <= 1 { return }
    x.actual = x.path[c - 2]
    n:= x.path[c - 1]
    c --
    x.path = remove (x.path, c)
    if ! contains (x.path, n) {
      n.inSubgraph = false
    }
    e:= connection (x.actual, n)
    if e == nil { ker.Stop (pack, 10) }
    e.inSubgraph = false
    i = uint(0)
    for {
      if i + 1 == c { break }
      n = x.path[i]
      n1:= x.path[i+1]
      if e == connection (n, n1) {
        e.inSubgraph = true
        break
      } else {
        i ++
      }
    }
  }
}


func (x *Imp) Save () {
//
  x.ncPtr = new (nodeCell)
  x.ncPtr.nPtr = x.actual
  x.ncPtr.next = x.ncPtr
}


func (x *Imp) Restore () {
//
  if x.ncPtr == nil { return }
  x.actual = x.ncPtr.nPtr
  x.ncPtr = x.ncPtr.next
}


func (x *Imp) Mark (m bool) {
//
  if x.actual == nil {
    return
  }
  x.actual.marked = m
}


func (x *Imp) Marked () bool {
//
  if x.nodeAnchor == x.nodeAnchor.nextN {
    return false
  }
  return x.actual.marked
}


func (x *Imp) A () {
//
  x.actual.dist = 0
}


func (x *Imp) B () bool {
//
  return x.actual.dist == 1 << 32 - 1
}


func (x *Imp) C () {
//
  x.actual.dist = x.postactual.dist + 1
  x.actual.predecessor = x.postactual
}


func (x *Imp) True (p Pred) bool {
//
  n:= x.nodeAnchor.nextN
  for n != x.nodeAnchor {
    if ! p (n.content) {
      return false
    }
    n = n.nextN
  }
  return true
}


func (x *Imp) TrueAct (p Pred) bool {
//
  n:= x.nodeAnchor.nextN
  for n != x.nodeAnchor {
    if n.inSubgraph {
      if ! p (n.content) {
        return false
      }
    }
    n = n.nextN
  }
  return true
}


func (x *Imp) Trav (op Op) {
//
  for n:= x.nodeAnchor.nextN; n != x.nodeAnchor; n = n.nextN {
    op (n.content)
  }
}


func (x *Imp) TravCond (op CondOp) {
//
  for n:= x.nodeAnchor.nextN; n != x.nodeAnchor; n = n.nextN {
    if ! n.inSubgraph {
      op (n.content, false)
    }
  }
  for n:= x.nodeAnchor.nextN; n != x.nodeAnchor; n = n.nextN {
    if n.inSubgraph {
      op (n.content, true)
    }
  }
}


func (x *Imp) Trav1 (op Op) {
//
  for e:= x.edgeAnchor.nextE; e != x.edgeAnchor; e = e.nextE {
    op (e.attrib)
  }
}


func (x *Imp) Trav1Cond (op CondOp) {
//
  for e:= x.edgeAnchor.nextE; e != x.edgeAnchor; e = e.nextE {
    if ! e.inSubgraph { op (e.attrib, false) }
  }
  for e:= x.edgeAnchor.nextE; e != x.edgeAnchor; e = e.nextE {
    if e.inSubgraph { op (e.attrib, true) }
  }
}


func (x *Imp) Trav3 (op Op3) {
//
  for e:= x.edgeAnchor.nextE; e != x.edgeAnchor; e = e.nextE {
    if e.nbPtr0.forward {
      op (e.attrib, e.nbPtr0.from.content, e.nbPtr1.from.content)
    }
    if x.directed {
     if e.nbPtr1.forward {
        op (e.attrib, e.nbPtr1.from.content, e.nbPtr0.from.content)
      }
    }
  }
}


func (x *Imp) Trav3Cond (op CondOp3) {
//
  for e:= x.edgeAnchor.nextE; e != x.edgeAnchor; e = e.nextE {
    if ! e.inSubgraph {
      if e.nbPtr0.forward {
        op (e.attrib, e.nbPtr0.from.content, e.nbPtr1.from.content, false)
      }
      if x.directed && e.nbPtr1.forward {
        op (e.attrib, e.nbPtr1.from.content, e.nbPtr0.from.content, false)
      }
    }
  }
  for e:= x.edgeAnchor.nextE; e != x.edgeAnchor; e = e.nextE {
    if e.inSubgraph {
      if e.nbPtr0.forward {
        op (e.attrib, e.nbPtr0.from.content, e.nbPtr1.from.content, true)
      }
      if x.directed && e.nbPtr1.forward {
        op (e.attrib, e.nbPtr1.from.content, e.nbPtr0.from.content, true)
      }
    }
  }
}


// Returns true, iff every node of x is accessible from every other one by a path.
func (x *Imp) TotallyConnected () bool {
//
  if x.nNodes <= 1 {
    return true
  }
  if x.directed {
    x.Isolate ()
  } else {
    x.depthfirstSearch ()
  }
  n:= x.nodeAnchor.nextN
  e0:= n.repr
  for n != x.nodeAnchor {
    if n.repr != e0 {
      return false
    }
    n = n.nextN
  }
  return true
}


// CLR 23.3, CLRS 22.3
func (x *Imp) depthfirstSearch () {
//
  x.preDepth ()
  if x.demo [Depth] {
    errh.Hint ("weiter mit Eingabetaste")
  }
  n:= x.nodeAnchor.nextN
  for n != x.nodeAnchor {
    if n.time0 == 0 {
      x.search (n, n, True)
    }
    n = n.nextN
  }
  if x.demo [Depth] {
    errh.DelHint()
  }
}


func (x *Imp) Acyclic () bool {
//
  if x.nodeAnchor.nextN == x.nodeAnchor { return true }
  x.nodeAnchor.marked = true
  x.depthfirstSearch ()
  return x.nodeAnchor.marked
}


func undurchlaufenerNachbar (n *node) *neighbour {
//
  nb:= n.nbPtr.nextNb
  for nb != n.nbPtr {
    if nb.edgePtr.inSubgraph || ! nb.forward {
      nb = nb.nextNb
    } else {
      return nb
    }
  }
  return nil
}


func undurchlaufen (a Any) bool {
//
  n:= a.(*neighbour)
  return undurchlaufenerNachbar (n.from) != nil
}


func (x *Imp) Eulerian () bool {
//
  if ! x.TotallyConnected () {
    return false // TODO Fleury's algorithm
  }
  p:= x.postactual
  a:= x.actual
  x.postactual = x.nodeAnchor
  x.actual = x.nodeAnchor
  e:= uint(0)
  n:= x.nodeAnchor.nextN
  for n != x.nodeAnchor {
// Prüfung auf Existenz von Eulerkreisen (genau dann, falls Graph
// ungerichtet:
//   wenn jede Ecke eine gerade Nachbarzahl hat,
// gerichtet:
//   wenn an jeder Ecke die Anzahl der forwarden edge mit der der ankommenden übereinstimmt)
// oder von Eulerwegen (genau dann, falls Graph
// ungerichtet:
//   wenn es genau zwei Ecken mit ungerader Nachbarzahl gibt,
// gerichtet:
//   wenn genau eine Ecke eine abgehende Kanten mehr als ankommende
//   und genau eine Ecke eine ankommende Kante mehr als abgehende hat)
    z:= uint(0)
    z1:= uint(0)
    nb:= n.nbPtr.nextNb
    for nb != n.nbPtr {
      if nb.forward {
        z ++
      } else {
        z1 ++
      }
      nb = nb.nextNb
    }
    if x.directed {
      if z == z1 + 1 {
        if x.postactual == x.nodeAnchor {
          x.postactual = n
          e ++
        } else {
          x.postactual = p
          x.actual = a
          return false
        }
      } else if z1 == z + 1 {
        if x.actual == x.nodeAnchor {
          x.actual = n
          e ++
        } else {
          x.postactual = p
          x.actual = a
          return false
        }
      }
    } else { // ! x.directed
      if z % 2 == 1 {
        if x.postactual == x.nodeAnchor {
          x.postactual = n
        } else if x.actual == x.nodeAnchor {
          x.actual = n
        } else {
          x.postactual = p
          x.actual = a
          return false
        }
        e ++
      }
    }
    n = n.nextN
  }
  switch e { case 0: // Euler cycle with random starting node
    x.postactual = x.nodeAnchor.nextN
    n:= rand.Natural (x.nNodes)
    for n > 0 {
      x.postactual = x.postactual.nextN
      n --
    }
    x.actual = x.postactual
  case 1:
    x.postactual = p
    x.actual = a
    return false
  case 2: // Euler way from postactual to actual node
    ;
  default:
    ker.Stop (pack, 100 + e)
  }
  x.clearSubgraph ()
  x.eulerPath = nil
  x.postactual.inSubgraph = true
  n = x.postactual
  n.inSubgraph = true
//  for j:= 0; j <= 9; j** { for a:= false TO true { write (E.content, a); ker.Msleep (100) } }
// attempt, to find an Euler way/cycle "by good luck":
  var nb *neighbour
  for {
    nb = undurchlaufenerNachbar (n)
    if nb == nil { ker.Stop (pack, 11) }
    // write3 (N.edgePtr.attrib, E.content, N.to.content, true)
    //  for j:= 0; j <= 9; j++ { for a:= false TO true { write (N.to.content, a); ker.Msleep (100) } } };
    nb.edgePtr.inSubgraph = true
    n = nb.to
    n.inSubgraph = true
    x.eulerPath = append (x.eulerPath, nb)
    if n == x.actual { break }
  }
// errh.Error ("erster Wegabschnitt gefunden", 0);
// as long there are undurchlaufene edges,
// den bisherigen Teil des Eulerwegs auf Ecken durchsuchen,
// von denen undurchlaufene Kanten abgehen,
// und von ihnen aus weitere Kreise finden und
// in den bisher gefundenen Teil des Eulerwegs einfügen:
  for {
    nb, ok:= existsnb (x.eulerPath, undurchlaufen)
    if ! ok { break }
    // for j:= 0; j <= 9; j++ { for a:= false; a <= true; a++ { // nonsense
    //   x.write3 (nb.edgePtr.attrib, nb.edgePtr.nbPtr0.from.content, nb.edgePtr.nbPtr1.from.content, a); ker.Msleep (100) } }
    n = nb.from
    n1:= n
    for {
      nb = undurchlaufenerNachbar (n)
      if nb == nil { ker.Stop (pack, 12) }
    // write3 (N.edgePtr.attrib, E.content, N.to.content, true)
    // for j:= 0 TO 9 { for a:= false TO true { write (N.to.content, a); ker.Msleep (100) } }
      nb.edgePtr.inSubgraph = true
      n = nb.to
      n.inSubgraph = true
      x.eulerPath = append (x.eulerPath, nb)
      if n == n1 { break } // zusätzlicher Kreis gefunden
    // errh.Error ("weiterer Teil eines Eulerwegs gefunden", 0)
    }
  }
  if x.demo [Euler] {
    x.write (x.postactual.content, true)
    wait ()
    for i:= uint(0); i < uint(len (x.eulerPath)); i++ {
      nb = x.eulerPath[i]
      x.write3 (nb.edgePtr.attrib, nb.edgePtr.nbPtr0.from.content, nb.edgePtr.nbPtr1.from.content, true)
      if nb.edgePtr.nbPtr0 == nb {
        x.write (nb.edgePtr.nbPtr1.from.content, true)
      } else {
        x.write (nb.edgePtr.nbPtr0.from.content, true)
      }
      if i + 1 < uint(len (x.eulerPath)) {
        wait()
      }
    }
  }
  return true
}


// Kruskal's algorithm, CLR 24.1-2, CLRS 23.1-2
func (x *Imp) Minimize () {
//
  if x.nNodes < 2 || x.directed || x.edgeAnchor.nextE == x.edgeAnchor {
    return
  }
  var set edgeSet = make ([]*edge, x.nEdges)
  n:= x.nodeAnchor.nextN
  for n != x.nodeAnchor {
    n.predecessor = nil
    n.repr = n
    n.inSubgraph = false
    n = n.nextN
  }
  e:= x.edgeAnchor.nextE
  for e != x.edgeAnchor {
    e.inSubgraph = false
    e = e.nextE
  }
  if x.nNodes == 1 {
    x.actual = x.nodeAnchor.nextN
    x.actual.inSubgraph = true
    return
  }
  e = x.edgeAnchor.nextE
  for e != x.edgeAnchor {
    set = append (set, e)
    sort.Sort (set)
    e.inSubgraph = false
    e = e.nextE
  }
  for len (set) > 0 {
    k:= set[0]
    set = set [1:]
    n = k.nbPtr0.from
    n1:= k.nbPtr1.from
    if x.demo [SpanTree] {
      x.write3 (e.attrib, n.content, n1.content, true)
      x.write (n.content, true)
      x.write (n1.content, true)
      wait()
    }
    if n.repr != n1.repr {
      n.inSubgraph = true
      n1.inSubgraph = true
      k.inSubgraph = true
      for n.predecessor != nil {
        n = n.predecessor
      }
      n1 = n1.repr
      n.predecessor = n1
      n = n.repr
      for n1.predecessor != nil {
        n1.repr = n
        n1 = n1.predecessor
      }
      n1.repr = n
    } else {
      if x.demo [SpanTree] {
        x.write3 (e.attrib, n.content, n1.content, false)
        wait()
      }
    }
  }
}


// topological Sort, CLR 23.4, CLRS 22.4
// TODO spec
func (x *Imp) Sort () {
//
  if x.nNodes < 2 || ! x.directed { return }
  x.depthfirstSearch ()
// sort list of nodes due to decrementing times, for which we supply a slice von *node:
  f:= make ([]*node, 2 * x.nNodes)
  for i:= uint(0); i < 2 * x.nNodes; i++ {
    f[i] = nil
  }
// partial function f: [0 .. 2 * nNodes - 1] -> *node with
// f[i]:= the node with time1 = i, if there is such, otherwise nodeAnchor
  n:= x.nodeAnchor.nextN
  for i:= uint(0); i < x.nNodes; i++ {
    f[n.time1 - 1] = n
    n = n.nextN
  }
// sort list of nodes by
// von vorne nach hinten jeweils die Ecke mit Zeit i an den Anfang der Liste holen:
  for i:= uint(0); i < 2 * x.nNodes; i++ {
    n = f[i]
    if n != nil { // put n to the head of the list:
      n.nextN.prevN, n.prevN.nextN = n.prevN, n.nextN
      n.nextN, n.prevN = x.nodeAnchor.nextN, x.nodeAnchor
      n.nextN.prevN = n
      x.nodeAnchor.nextN = n
    }
  }
}


// strongly connected components, CLR 23.5, CLRS 22.5
// TODO spec
func (x *Imp) Isolate () {
//
  if x.nNodes < 1 || ! x.directed {
    return
  }
// depth first search with sorting of the list of nodes by decrementing times:
  x.Sort ()
// essence of the algorithm: invert directions of all edges:
  x.Inv ()
// and now once more depth first search,
// starting with the highest time of the first depth first search:
  x.depthfirstSearch ()
// the depth first search trees are now the strongly connected components with common repr
// finally again invert the directions of all edges
  x.Inv ()
// all nodes in the actual subgraph:
  n:= x.nodeAnchor.nextN
  for n != x.nodeAnchor {
    n.inSubgraph = true
    n = n.nextN
  }
// and furthermore all edges, that connect two nodes in the same strongly connected component:
  e:= x.edgeAnchor.nextE
  for e != x.edgeAnchor {
    e.inSubgraph = e.nbPtr0.from.repr == e.nbPtr1.from.repr
    e = e.nextE
  }
}


func (x *Imp) IsolateAct () {
//
  x.Isolate ()
// nur genau die Ecken in dem aktuellen Untergraph, die in der gleichen starken Zusammenhangskomponente wie die aktuelle Ecke liegen:
  n:= x.nodeAnchor.nextN
  for n != x.nodeAnchor {
    n.inSubgraph = n.repr == x.actual.repr
    n = n.nextN
  }
// und dazu genau die Kanten, die diese Ecken verbinden:
  e:= x.edgeAnchor.nextE
  for e != x.edgeAnchor {
    e.inSubgraph = e.nbPtr0.from.inSubgraph && e.nbPtr1.from.inSubgraph
    e = e.nextE
  }
}


func (x *Imp) Equivalent () bool {
//
  if x.Empty () {
    return false
  }
  x.Isolate ()
  return x.actual.repr == x.postactual.repr
}


const
  cluint = uint(4) // Codelen (uint(0))


func (x *Imp) Codelen () uint {
//
  c:= cluint
  if x.nNodes > 0 {
    for n:= x.nodeAnchor.nextN; n != x.nodeAnchor; n = n.nextN {
      c += cluint + Codelen (n.content)
    }
    c += 3 * cluint
    if x.nEdges > 0 {
      for e:= x.edgeAnchor.nextE; e != x.edgeAnchor; e = e.nextE {
        if x.edgeAnchor.attrib != nil {
          c += cluint + Codelen (e.attrib)
        }
        c += 2 * (cluint + Codelen (true))
      }
    }
  }
  return c
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  i, a:= uint(0), cluint
  copy (b[i:i+a], Encode (x.nNodes))
  i += a
  if x.nNodes == 0 {
    return b
  }
  z:= uint(0)
  for n:= x.nodeAnchor.nextN; n != x.nodeAnchor; n = n.nextN {
    k:= Codelen (n.content)
    copy (b[i:i+a], Encode (k))
    i += a
    copy (b[i:i+k], Encode (n.content))
    i += k
    n.dist = z
    z ++
  }
  copy (b[i:i+a], Encode (x.postactual.dist))
  i += a
  copy (b[i:i+a], Encode (x.actual.dist))
  i += a
  copy (b[i:i+a], Encode (x.nEdges))
  i += a
  if x.nEdges == 0 { return b }
  for e:= x.edgeAnchor.nextE; e != x.edgeAnchor; e = e.nextE {
    a = cluint
    if x.edgeAnchor.attrib != nil {
      k:= Codelen (e.attrib)
      copy (b[i:i+a], Encode (k))
      i += a
      copy (b[i:i+k], Encode (e.attrib))
      i += k
    }
    copy (b[i:i+a], Encode (e.nbPtr0.from.dist))
    i += a
    a = Codelen (true)
    copy (b[i:i+a], Encode (e.nbPtr0.forward))
    i += a
    a = cluint
    copy (b[i:i+a], Encode (e.nbPtr1.from.dist))
    i += a
    a = Codelen (true)
    copy (b[i:i+a], Encode (e.nbPtr1.forward))
    i += a
  }
  return b
}


func (x *Imp) Decode (b []byte) {
//
  i, a:= uint(0), cluint
  x.nNodes = Decode (uint(0), b[i:i+a]).(uint)
  i += a
  if x.nNodes == 0 {
    return
  }
  for n:= uint(0); n < x.nNodes; n++ {
    k:= Decode (uint(0), b[i:i+a]).(uint)
    i += a
    cont:= Decode (x.nodeAnchor.content, b[i:i+k])
    x.insertedNode (cont)
    i += k
  }
  a = cluint
  p:= Decode (uint(0), b[i:i+a]).(uint)
  i += a
  c:= Decode (uint(0), b[i:i+a]).(uint)
  i += a
  n:= x.nodeAnchor.nextN
  z:= uint(0)
  for n != x.nodeAnchor {
    if z == p {
      x.postactual = n
    }
    if z == c {
      x.actual = n
    }
    z ++
    n = n.nextN
  }
  x.nEdges = Decode (uint(0), b[i:i+a]).(uint)
  i += a
  if x.nEdges == 0 { return }
  for z:= uint(0); z < x.nEdges; z++ {
    e:= newEdge ()
    a = cluint
    if x.edgeAnchor.attrib == nil {
      e.attrib = nil
    } else {
      k:= Decode (uint(0), b[i:i+a]).(uint)
      i += a
      e.attrib = Decode (x.edgeAnchor.attrib, b[i:i+k])
      i += k
    }
    z1:= Decode (uint(0), b[i:i+a]).(uint)
    if z1 > x.nNodes { ker.Stop (pack, 13) }
    i += a
    n0:= x.nodeAnchor.nextN
    for z1 > 0 {
      n0 = n0.nextN
      z1 --
    }
    a = Codelen (true)
    bo:= Decode (true, b[i:i+a]).(bool)
    i += a
    e.nbPtr0 = newNeighbour (e, n0, nil, bo) // e.nbPtr0.to see below
    insertNeighbour (e.nbPtr0, n0)
    a = cluint
    z1 = Decode (uint(0), b[i:i+a]).(uint)
    if z1 > x.nNodes { ker.Stop (pack, 14) }
    i += a
    n0 = x.nodeAnchor.nextN
    for z1 > 0 {
      n0 = n0.nextN
      z1 --
    }
    e.nbPtr0.to = n0
    a = Codelen (true)
    bo = Decode (true, b[i:i+a]).(bool)
    i += a
    d:= e.nbPtr0.forward != bo
    if d != x.directed { ker.Stop (pack, 15) }
    e.nbPtr1 = newNeighbour (e, n0, e.nbPtr0.from, bo)
    insertNeighbour (e.nbPtr1, n0)
    e.inSubgraph = false
    e.nextE = x.edgeAnchor
    e.prevE = x.edgeAnchor.prevE
    e.prevE.nextE = e
    x.edgeAnchor.prevE = e
  }
  x.path = nil
  x.eulerPath = nil
  x.ncPtr = nil
}


func (x *Imp) Install (p CondOp, p3 CondOp3) {
//
  x.write, x.write3 = p, p3
}


func (x *Imp) Set (d Demo) {
//
  x.demo[d] = true
  if d == Cycle { x.demo[Depth] = true } // Cycle without Depth is pointless
}
