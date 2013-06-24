package lr

// (c) Christian Maurer   v. 121030 - license see murus.go

// >>> readers/writers problem: implementation with monitors
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. TODO

import (
  . "murus/obj"
  "murus/fmon"
)
type
  ImpFMon struct {
                 fmon.FarMonitor
                 }


func NewFMon (s string, p uint) *ImpFMon {
//
  var nL, nR int
  c:= func (a Any, k uint) bool {
        switch k { case lIn:
          return nR == 0
        case rIn:
          return nL == 0 && nR == 0
        }
        return true
      }
  f:= func (a Any, k uint) Any {
        switch k { case lIn:
          nL++
        case lOut:
          nL--
        case rIn:
          nR = 1
        case rOut:
          nR = 0
        }
        return true
      }
  return &ImpFMon { fmon.New (true, 4, f, c, s, p) }
}


func (x *ImpFMon) LeftIn () {
//
  x.F (true, lIn)
}


func (x *ImpFMon) LeftOut () {
//
  x.F (true, lOut)
}


func (x *ImpFMon) RightIn () {
//
  x.F (true, rIn)
}


func (x *ImpFMon) RightOut () {
//
  x.F (true, rOut)
}
