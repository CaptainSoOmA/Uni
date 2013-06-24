package set

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  . "murus/obj"
  "murus/ker"
)
type
  node struct {
          ptr *tree
         next *node
              }


// x.path is the list of nodes from actual up to root. 
// min/max == true, iff the actual node w.r.t. Less
// is the smallest/largest object in x.
func (x *Imp) defPath () (bool, bool) {
//
  t:= x.anchor
  x.path = &node { t, nil }
  min, max:= true, true
  for {
    if t == x.actual {
      break
    }
    if x == nil { ker.Panic ("murus/set/path.go defPath: x == nil") }
    if t.root == nil { ker.Panic ("murus/set/path.go defPath: t.root == nil") }
    if Less (x.actual.root, t.root) { // TODO avoid crash
      t, max = t.left, false
    } else {
      t, min = t.right, false
    }
    x.path = &node { t, x.path }
  }
  return min && t.left == nil, max && t.right == nil
}


func (x *Imp) below (f bool) *tree {
//
  if f {
    return x.path.ptr.right
  }
  return x.path.ptr.left
}


func (x *Imp) somethingBelow (f bool) bool {
//
  return x.below (f) != nil
}


func (x *Imp) abovePointsToCurrent (f bool) bool {
//
  t:= x.path.next.ptr.right
  if ! f {
    t = x.path.next.ptr.left
  }
  return t == x.path.ptr
}


func (x *Imp) up () {
//
  x.path = x.path.next
}


func (x *Imp) pointer () *tree {
//
  return x.path.ptr
}
