package buf

// (c) Christian Maurer   v. 130516 - license see murus.go

// >>> Implementation with conditioned monitor

import (
  . "murus/obj"
  "murus/mon"
)
type
  ImpMon struct {
                Buffer
                mon.Monitor
                }


func NewMon (a Any, n uint) *ImpMon {
//
  if a == nil || n == 0 { return nil }
  x:= new (ImpMon)
  x.Buffer = newBuffer (a, n)
  p:= func (a Any, i uint) bool {
        if i == 0 { // Ins
          return ! x.Buffer.Full ()
        }
        return x.Buffer.Num() > 0 // Get
      }
  f:= func (a Any, i uint) Any {
        if i == 0 { // Ins
          x.Buffer.Ins (a)
          return a
        } // Get
        return x.Buffer.Get ()
      }
  x.Monitor /* x.m */ = mon.NewC (3, f, p)
  return x
}


func (x *ImpMon) Num () uint {
//
  return x.Buffer.Num ()
}


func (x *ImpMon) Full () bool {
//
  return x.Buffer.Full ()
}


func (x *ImpMon) Ins (a Any) {
//
  _ = x.F (a, 0)
}


func (x *ImpMon) Get () Any {
//
  var a Any // belanglos
  return x.F (a, 1)
}
