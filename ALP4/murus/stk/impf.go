package stk

// (c) Christian Maurer   v. 130125 - license see murus.go

import (
  . "murus/obj"
  "murus/fmon"
)
type
  ImpF struct {
       object Any
              *Imp
            m *fmon.Imp
              }


func NewF (a Any, s string, p uint) *ImpF {
//
  x:= new (ImpF)
  x.object = Clone (a)
  x.Imp = New (a)
  c:= func (a Any, i uint) bool {
        if i == 1 { // push
          return true
        }
        return ! x.Imp.Empty () // top, pop
      }
  f:= func (a Any, i uint) Any {
        if i == 2 { // top
          return x.Imp.Top ()
        }
        return x.object // push, pop
      }
  x.m = fmon.New (a, 3, f, c, s, p)
  return x
}


func (x *ImpF) Empty () bool {
//
  return false // pointless to be called
}


func (x *ImpF) Push (a Any) {
//
  x.m.F (a, 0)
}


func (x *ImpF) Pop () {
//
  x.m.F (x.object, 1)
}


func (x *ImpF) Top () Any {
//
  return x.m.F (x.object, 2)
}
