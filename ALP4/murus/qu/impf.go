package qu

// (c) Christian Maurer   v. 130506 - license see murus.go

import (
  . "murus/obj"
  "murus/fmon"
)
type
  ImpF struct {
              fmon.FarMonitor
              }


func NewF (a Any, s string, p uint) *ImpF {
//
  q:= New (a)
  c:= func (a Any, i uint) bool {
        switch i {
        case 2: // Get
          return ! q.Empty()
        }
        return true // Num, Ins
      }
  f:= func (a Any, i uint) Any {
        switch i {
        case 0:
          return q.Num()
        case 1:
          q.Ins (a)
        case 2:
          return q.Get()
        }
        return a
      }
  return &ImpF { fmon.New (q, 3, f, c, s, p) }
}


func (x *ImpF) Terminate () {
//
  x.Terminate ()
}


func (x *ImpF) Num () uint {
//
  var a Any // belanglos
  return x.F (a, 0).(uint)
}


func (x *ImpF) Ins (a Any) {
//
  x.F (a, 1)
}


func (x *ImpF) Get () Any {
//
  var a Any // belanglos
  return x.F (a, 2)
}
