package sem

// (c) Christian Maurer   v. 121030 - license see murus.go

import (
  . "murus/obj"
  "murus/fmon"
)
type
  ImpF struct {
              *fmon.Imp
              }


func NewF (n uint, s string, p uint) *ImpF {
//
  val:= n
  c:= func (a Any, i uint) bool {
        if i == 0 { return val > 0 } // P
        return true // V
      }
  f:= func (a Any, i uint) Any {
        if i == 0 { // P
          val--
        } else {
          val++ // V
        }
        return true
      }
  return &ImpF { fmon.New (false, 2, f, c, s, p) }
}


func (x *ImpF) P () {
//
  x.F (true, 0)
}


func (x *ImpF) V () {
//
  x.F (true, 1)
}
