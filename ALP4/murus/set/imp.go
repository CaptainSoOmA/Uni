package set

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  "murus/ker"
  . "murus/obj"
)
type
  Imp struct {
 emptyObject Any
      anchor,
      actual *tree
         num uint
        path *node
             }


func (x *Imp) check (a Any) {
//
  if ! TypeEq (x.emptyObject, a) { TypeNotEqPanic (x.emptyObject, a) }
}


func (x *Imp) imp (Y Any) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  x.check (y.emptyObject)
  return y
}


func New (a Any) *Imp {
//
  PanicIfNotOk (a)
  x:= new (Imp)
  x.emptyObject = Clone (a)
  x.anchor, x.actual = nil, nil
  return x
}


func (x *Imp) Offc () bool {
//
  return x.actual == nil
}


func (x *Imp) Empty () bool {
//
  return x.anchor == nil
}


func (x *Imp) Clr () {
//
  x.anchor, x.actual = nil, nil
  x.num = 0
  x.path = nil
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  x.Clr ()
  y.Trav (func (a Any) { x.Ins (a) })
  x.num = y.num
}


func (x *Imp) Clone () Object {
//
  y:= New (x.emptyObject)
  x.Trav (func (a Any) { y.Ins (a) })
  y.num = x.num
  return y
}


func (x *Imp) e (y *Imp, r Rel) bool {
//
  if x.num != y.num { return false }
  if x.anchor == nil { return true }
  xact, yact:= x.actual, y.actual
  x.Jump (false)
  y.Jump (false)
  for {
    if r (x.actual.root, y.actual.root) {
      if x.Eoc (true) {
        x.actual, y.actual = xact, yact
        return true
      } else {
        x.Step (true)
        y.Step (true)
      }
    } else {
      break
    }
  }
  x.actual, y.actual = xact, yact
  return false
}


func (x *Imp) Eq (Y Object) bool {
//
  y:= x.imp (Y)
  return x.e (y, Eq)
}


func (x *Imp) Less (Y Object) bool {
//
  y:= x.imp (Y)
  if ! Less (x.num, y.num) { return false }
  if x.anchor == nil { return true }
  return x.All (func (a Any) bool { return y.Ex (a) } )
}


func (x *Imp) Num () uint {
//
  return x.num
}


func (x *Imp) NumPred (p Pred) uint {
//
  return x.anchor.numPred (p)
}


func (x *Imp) Equiv (Y Iterator, r Rel) bool {
//
  y:= x.imp (Y)
  return x.e (y, r)
}


func (x *Imp) Ex (a Any) bool {
//
  x.check (a)
  if t, c:= x.anchor.contained (a); c {
    x.actual = t
    return true
  }
  return false
}


func (x *Imp) All (p Pred) bool {
//
  return x.anchor.all (p)
}


func (x *Imp) Sort () {
//
  y:= New (x.emptyObject)
  x.Trav (func (a Any) { y.Ins (a) } )
  x.anchor, x.num = y.anchor, y.num
  y.anchor = nil
//  x.actual = y.actual
//  x.Jump (false)
  x.actual = x.anchor
  for x.actual.left != nil {
    x.actual = x.actual.left
  }
}


func (x *Imp) Step (forward bool) {
//
  if x == nil { return }
  min, max:= x.defPath ()
  if forward {
    if max { return }
  } else {
    if min { return }
  }
  if x.somethingBelow (forward) {
    x.actual = x.below (forward)
    for {
      if forward {
        if x.actual.left == nil {
          break
        }
        x.actual = x.actual.left
      } else {
        if x.actual.right == nil {
          break
        }
        x.actual = x.actual.right
      }
    }
  } else {
    for {
      if ! x.abovePointsToCurrent (forward) {
        x.up ()
        x.actual = x.pointer ()
        return
      }
      x.up ()
    }
  }
}


func (x *Imp) Jump (toEnd bool) {
//
  x.actual = x.anchor
  for {
    if toEnd {
      if x.actual.right == nil {
        break
      }
      x.actual = x.actual.right
    } else {
      if x.actual.left == nil {
        break
      }
      x.actual = x.actual.left
    }
  }
}


func (x *Imp) Eoc (forward bool) bool {
//
  t:= x.anchor
  for t != nil {
    if t == x.actual {
      if forward {
        return t.right == nil
      } else {
        return t.left == nil
      }
    }
    if forward {
      t = t.right
    } else {
      t = t.left
    }
  }
  return false
}


func (x *Imp) Get () Any {
//
  if x.anchor == nil { return nil }
  if x.actual == nil { ker.Panic ("set.Get error: x.actual == nil") }
  return Clone (x.actual.root)
}


func (x *Imp) Put (a Any) {
//
  x.check (a)
  x.Del ()
  x.Ins (a)
}


func (x *Imp) Ins (a Any) {
//
  x.check (a)
  if x.anchor == nil {
    x.anchor = leaf (a)
    x.actual = x.anchor
    x.num = 1
  } else {
    var t *tree
    x.anchor, t = x.anchor.ins (a)
    if t != nil {
      x.actual = t
      x.num ++
    }
  }
}


func (x *Imp) Del () Any {
//
  if x.anchor == nil {
    return nil
  }
  act:= x.actual
  toDel:= x.actual.root
  x.Step (true) // to set "actual" one step ahead
  var tmp Any = nil
  if act == x.actual { // the root to remove is the last right node in x,
                       // "actual" must be reset one position or set to nil, see below
  } else {
    tmp = Clone (toDel)
  }
  oneLess:= false
  x.anchor, oneLess = x.anchor.del (toDel)
  if oneLess {
    if act == x.actual { // the root to remove was the last right node of x
      if x.num == 1 {    // see above
        x.actual = nil   // x is now empty
      } else {
        x.Jump (true)
      }
    } else {
      if x.Ex (tmp) { // thus the above copy-action to "tmp": "actual" might have been
                      // rotated off while deleting, with this trick is it restored !
      }
    }
    x.num --
  }
  return Clone (act.root)
}


func (x *Imp) ExPred (p Pred, f bool) bool {
//
  t:= x.anchor.exPred (p)
  if t == nil {
    return false
  }
  x.actual = t
  return true
}


func (x *Imp) StepPred (p Pred, f bool) bool {
//
  if x == nil { return false }
  xact:= x.actual
  for ! x.Eoc (f) {
    x.Step (f)
    if x.Eoc (f) { break }
    if p (x.actual.root) {
      return true
    }
  }
  x.actual = xact
  return false
}


func (x *Imp) ExGeq (a Any) bool {
//
  t:= x.anchor.first (a)
  if t == nil {
    return false
  }
  x.actual = t
  return true
}


func (x *Imp) Trav (op Op) {
//
  x.anchor.trav (op)
}


func (x *Imp) TravPred (p Pred, op Op) {
//
  x.anchor.travPred (p, op)
}


func (x *Imp) TravCond (p Pred, op CondOp) {
//
  x.anchor.travCond (p, op)
}


func (x *Imp) Filter (Y Iterator, p Pred) {
//
  y:= x.imp (Y)
  if x.anchor == nil { return }
  y.Clr ()
  x.Trav (func (a Any) { if p (a) { y.Ins (a) } })
  y.Jump (false)
}


func (x *Imp) Split (Y Iterator) {
//
  y:= x.imp (Y)
  y.Clr ()
  if x.anchor == nil { return }
  x1:= New (x.emptyObject)
  b:= x.actual.root
  x.Trav (func (a Any) { if Less (a, b) { x1.Ins (a) } else { y.Ins (a) } })
  x.anchor, x.num = x1.anchor, x1.num
  x.Jump (false)
  y.Jump (false)
}


func (x *Imp) Cut (Y Iterator, p Pred) {
//
  y:= x.imp (Y)
  y.Clr ()
  if x.anchor == nil { return }
  x1:= New (x.emptyObject)
  x.Trav (func (a Any) { if p (a) { y.Ins (a) } else { x1.Ins (a) } })
  x.anchor, x.num = x1.anchor, x1.num
  x.Jump (false)
  y.Jump (false)
}


func (x *Imp) ClrPred (p Pred) {
//
  if x.anchor == nil { return }
  y:= New (x.emptyObject)
  x.Trav (func (a Any) { if ! p (a) { y.Ins (a) } })
  x.anchor, x.num = y.anchor, y.num
  x.Jump (false)
}


func (x *Imp) Join (Y Iterator) {
//
  y:= x.imp (Y)
  y.Trav (func (a Any) { x.Ins (a) })
  y.Clr ()
  x.Jump (false)
}


func (x *Imp) Codelen () uint {
//
  n:= uint(4)
  x.Trav (func (a Any) { n += 4 + Codelen (a) })
  return n
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  copy (b[:4], Encode (x.num))
  i:= uint(4)
  x.Trav (func (a Any) {
            k:= Codelen (a)
            copy (b[i:i+4], Encode (k))
            i += 4
            copy (b[i:i+k], Encode (a))
            i += k
          })
  return b
}


func (x *Imp) Ordered () bool {
//
  if x.num <= 1 { return true }
  x.Jump (false)
  result, first, o:= true, true, x.actual.root
  x.Trav (func (a Any) {
            if first {
              first = false
            } else {
              if ! Less (o, a) {
                result = false
              }
            }
            o = a
          })
  return result
}


func (x *Imp) Decode (b []byte) {
//
  x.Clr ()
  n:= Decode (uint(0), b[:4]).(uint)
  i:= uint(4)
  for j:= uint(0); j < n; j++ {
    k:= Decode (uint(0), b[i:i+4]).(uint)
    i += 4
    a:= Decode (x.emptyObject, b[i:i+k])
    i += k
    x.Ins (a)
  }
}


// func init () { var _ Set = New (int(0)) }
