package buf

// (c) Christian Maurer   v. 121030 - license see murus.go

import (
  "murus/ker"; . "murus/obj"
  "murus/fmon"
)
type
  ImpF struct {
       object Any
            b *buffer
            m *fmon.Imp
              }


func NewF (a Any, n uint, s string, p uint) *ImpF {
//
  if a == nil || n == 0 { return nil }
  x:= new (ImpF)
  x.object = Clone (a)
  x.b = newBuffer (a, n)
  c:= func (a Any, i uint) bool {
        if i == 0 { // Ins
          return true
        }
        return x.b.Num() > 0 // Get
      }
  f:= func (a Any, i uint) Any {
        switch i {
        case 1: // Get
          return x.b.Get ()
        }
        return x.object // Ins
      }
  x.m = fmon.New (a, 2, f, c, s, p)
  return x
}


func (x *ImpF) Num () uint {
//
  ker.Stop ("buf far", 1) // pointless to be called
  return 0
}


func (x *ImpF) Full () bool {
//
  ker.Stop ("buf far", 2) // pointless to be called
  return false
}


func (x *ImpF) Ins (a Any) {
//
  _ = x.m.F (a, 0)
}


func (x *ImpF) Get () Any {
//
  return x.m.F (x.object, 1)
}
