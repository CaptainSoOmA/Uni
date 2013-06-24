package piset

// (c) Christian Maurer   v. 130115 - license see murus.go

// >>> still some things TODO

import (
  . "murus/obj"; "murus/ker"; "murus/str"
  "murus/pseq"; "murus/qu"; "murus/set"
  "murus/piset/index"
)
const (
  pack = "piset"
  suffix = "seq"
)
type
  Imp struct {
      object Object
        name string
        file *pseq.Imp
       index *index.Imp
           f func (Object) Object
         pos uint
        tree *set.Imp
     posPool *qu.Imp
             }


func (x *Imp) check (a Any) {
//
  if ! TypeEq (x.object, a) { TypeNotEqPanic (x.object, a) }
}


func (x *Imp) imp (a Any) *Imp {
//
  y, ok:= a.(*Imp)
  if ! ok { TypeNotEqPanic (x, a) }
  x.check (y)
  return y
}


func (x *Imp) ob (a Any) Object {
//
  y, ok:= a.(Object)
  if ! ok { TypeNotEqPanic (x, a) }
  x.check (y)
  return y
}


func New (o Object, f ObjectFunc) *Imp {
//
  x:= new (Imp)
  x.object = o.Clone () // TODO type information behind o lost, e.g. Editor
  x.file = pseq.New (x.object)
  x.f, x.index = f, index.New (f (o))
  x.tree = set.New (x.index)
  x.posPool = qu.New (uint(0))
  return x
}


func (x *Imp) Terminate () {
//
  x.file.Terminate()
}


func (x *Imp) Offc () bool {
//
  return x.Empty()
}


func (x *Imp) Name (s string) {
//
  if str.Empty (s) { return }
  x.name = s
  x.file.Name (x.name + "." + suffix)
  x.tree.Clr ()
  x.posPool.Clr ()
  if x.file.Empty () { return }
  x.file.Trav (func (a Any) {
    x.object = a.(Object)
    if x.object.Empty() {
      x.posPool.Ins (x.pos)
    } else {
      x.index.Set (x.f (x.object), x.pos)
      x.tree.Ins (x.index)
    }
    x.pos ++
  })
  x.pos = 0
  x.Jump (false)
}


func (x *Imp) Rename (s string) {
//
  if str.Empty (s) || x.name == s { return }
  x.name = s
  x.file.Rename (x.name + "." + suffix)
}


func (x *Imp) Empty () bool {
//
  return x.tree.Empty ()
}


func (x *Imp) Clr () {
//
  x.file.Clr ()
  x.tree.Clr ()
  x.pos = 0
  x.posPool.Clr ()
  x.object.Clr ()
}


func (x *Imp) Less (Y Object) bool {
//
  return x.tree.Less (x.imp (Y).tree)
}


func (x *Imp) Num () uint {
//
  return x.tree.Num ()
}


func (x *Imp) NumPred (p Pred) uint {
//
  n:= uint(0)
  x.TravPred (p, func (a Any) { n++ })
  return n
}

/*
func (x *Imp) Equiv (C Collector, r Rel) bool {
//
  y:= x.imp (C)
  ker.Panic ("piset.Equiv not yet implemented")
  return y == nil
}
*/

func (x *Imp) Ex (a Any) bool {
//
  o:= x.ob (a)
  x.index.Set (x.f (o), 0)
  if x.tree.Ex (x.index) {
    if x.index.Eq (x.tree.Get().(*index.Imp)) {
      x.index = x.tree.Get().(*index.Imp)
      x.pos = x.index.Pos()
      x.file.Seek (x.pos)
      return true
    }
  }
  return false
}


func (x *Imp) Ins (a Any) {
//
  o:= x.ob (a)
  if o.Empty () { return } // oder doch leere Objekte als Element zulassen ?
  if x.Ex (o) {
    return
  }
  if x.posPool.Empty () {
    x.pos = x.file.Num ()
  } else {
    x.pos = x.posPool.Del().(uint)
  }
  x.file.Seek (x.pos)
  x.file.Put (o)
  x.index.Set (x.f (o), x.pos)
  x.tree.Ins (x.index)
  x.file.Seek (x.pos)
}


func (x *Imp) Step (forward bool) {
//
  if x.tree.Eoc (forward) { return }
  x.tree.Step (forward)
  x.index = x.tree.Get().(*index.Imp)
  x.pos = x.index.Pos()
  x.file.Seek (x.pos)
}


func (x *Imp) Jump (toEnd bool) {
//
  if x.tree.Empty () { return }
  x.tree.Jump (toEnd)
  x.index = x.tree.Get().(*index.Imp)
  x.pos = x.index.Pos()
  x.file.Seek (x.pos)
}


func (x *Imp) Eoc (atEnd bool) bool {
//
  return x.tree.Eoc (atEnd)
}


func (x *Imp) Get () Any {
//
  if x.tree.Empty () { return nil }
  x.file.Seek (x.pos)
  x.object = x.file.Get ().(Object)
  return x.object.Clone ()
}


func (x *Imp) Put (a Any) {
//
  o:= x.ob (a)
  if x.tree.Empty () { return }
  x.file.Put (o)
  x.index.Set (x.f (o), x.pos)
  x.tree.Put (x.index)
}


func (x *Imp) Del () Any {
//
  if x.tree.Empty () { return nil }
  x.file.Seek (x.pos)
  o:= x.file.Get ().(Object)
  x.object.Clr ()
  x.file.Put (x.object)
  x.tree.Del ()
  x.posPool.Ins (x.pos)
  if x.tree.Empty () {
    x.pos = 0
  } else {
    x.index = x.tree.Get().(*index.Imp)
    x.pos = x.index.Pos()
    x.file.Seek (x.pos)
  }
  return o.Clone()
}


func (x *Imp) Sort () {
//
  x.tree.Sort ()
  x.Jump (false)
  x.index = x.tree.Get().(*index.Imp)
  x.pos = x.index.Pos()
}


func (x *Imp) Ordered () bool {
//
  return true // TODO
}


func (x *Imp) ExGeq (a Any) bool {
//
  y, ok:= a.(Object)
  if ! ok { return false } // ker.Panic ("piset.ExGeq ")
  if x.tree.Empty () { return false }
  x.index.Set (x.f (y), 0)
  if x.tree.ExGeq (x.index) {
    x.index = x.tree.Get().(*index.Imp)
    x.pos = x.index.Pos()
    x.file.Seek (x.pos)
    return true
  }
  return false
}


func (x *Imp) All (p Pred) bool {
//
  if x.tree.Empty () { return true }
  defer x.Jump (false)
  x.tree.Jump (false)
  for {
    x.index = x.tree.Get().(*index.Imp)
    x.pos = x.index.Pos()
    x.file.Seek (x.pos)
    if ! p (x.file.Get ().(Object)) {
      return false
    }
    if x.Eoc (true) {
      break
    }
  }
  return true
}


func (x *Imp) ExPred (p Pred, f bool) bool {
//
  defer x.Jump (false)
  x.tree.Jump (f)
  for {
    x.index = x.tree.Get().(*index.Imp)
    x.file.Seek (x.index.Pos())
    if p (x.file.Get ().(Object)) {
      x.pos = x.index.Pos()
      return true
    }
    if x.Eoc (f) {
      break
    }
    x.tree.Step (f)
  }
  return false
}


func (x *Imp) StepPred (p Pred, f bool) bool {
//
  defer x.Jump (false)
  x.tree.Jump (f)
  for {
    x.index = x.tree.Get().(*index.Imp)
    x.file.Seek (x.index.Pos())
    if p (x.file.Get ().(Object)) {
      x.pos = x.index.Pos()
      return true
    }
    if x.Eoc (f) {
      break
    }
    x.tree.Step (f)
  }
  return false
}


func (x *Imp) Trav (op Op) {
//
  if x.tree.Empty () { return }
  ok:= false
  x.tree.Trav (func (a Any) {
    x.index, ok = a.(*index.Imp)
    if ! ok { return }
    x.file.Seek (x.index.Pos())
    x.object = x.file.Get ().(Object)
    op (x.object)
  })
  x.Jump (false)
}


func (x *Imp) TravPred (p Pred, op Op) {
//
  if x.tree.Empty () { return }
  ok:= false
  x.tree.Trav (func (a Any) {
    x.index, ok = a.(*index.Imp)
    if ! ok { return }
    x.file.Seek (x.index.Pos())
    x.object = x.file.Get ().(Object)
    if p (x.object) {
      op (x.object)
    }
  })
  x.Jump (false)
}


func (x *Imp) TravCond (p Pred, op CondOp) {
//
  if x.tree.Empty () { return }
  ok:= false
  x.tree.Trav (func (a Any) {
    x.index, ok = a.(*index.Imp)
    if ! ok { return }
    x.file.Seek (x.index.Pos())
    x.object = x.file.Get ().(Object)
    op (x.object, p (x.object))
  })
  x.Jump (false)
}


func (x *Imp) Filter (Y Iterator, p Pred) {
//
  y:= x.imp (Y)
  if x.tree.Empty () { return }
  y.Clr ()
  x.file.Trav (func (a Any) {
    x.object = a.(Object)
    if ! x.object.Empty() && p (x.object) {
      y.Ins (x.object)
    }
  })
  y.pos = 0
  y.Jump (false)
}


func (x *Imp) Cut (Y Iterator, p Pred) {
//
  y:= x.imp (Y)
  y.Terminate ()
  y = New (x.object, x.f)
  if x.tree.Empty () { return }
  x.tree.Clr ()
  x.pos = 0
  x.file.Trav (func (a Any) {
    x.object = a.(Object)
    if x.object.Empty() {
      if p (x.object) {
        y.Ins (x.object)
        a = x.object
        x.posPool.Ins (x.pos)
      } else {
        x.index.Set (x.f (x.object), x.pos)
        x.tree.Ins (x.index)
      }
      x.pos++
    }
  })
  x.Jump (false)
  y.Jump (false)
}


func (x *Imp) ClrPred (p Pred) {
//
  ker.Panic ("piset.ClrPred not yet implemented")
}


func (x *Imp) Split (Y Iterator) {
//
//  y:= x.imp (Y)
  ker.Panic ("piset.Split not yet implemented")
}


func (x *Imp) Join (Y Iterator) {
//
  y:= x.imp (Y)
  if y.tree.Empty () { return }
  y.file.Trav (func (a Any) {
    x.object.Decode (a.([]byte))
    if ! x.object.Empty () {
      x.Ins (x.object)
    }
  })
  x.Jump (false)
  y.Clr ()
}
