package lr

// (c) Christian Maurer   v. 130409 - license see murus.go

// >>> left/right problem: implementation with conditional monitors
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 155 ff.

import (
  . "murus/obj"
  "murus/mon"
)
type
  ImpCMon struct {
                 mon.Monitor
                 }


func NewCMon () *ImpCMon {
//
  var nR, nL int
  c:= func (a Any, k uint) bool {
        switch k { case lIn:
          return nR == 0
        case rIn:
          return nL == 0
        }
        return true
      }
  f:= func (a Any, k uint) Any {
        switch k { case lIn:
          nL++
        case lOut:
          nL--
        case rIn:
          nR++
        case rOut:
          nR--
        }
        return nil
      }
  return &ImpCMon { mon.NewC (4, f, c) }
}


func (x *ImpCMon) LeftIn () {
//
  x.F (nil, lIn)
}


func (x *ImpCMon) LeftOut () {
//
  x.F (nil, lOut)
}


func (x *ImpCMon) RightIn () {
//
  x.F (nil, rIn)
}


func (x *ImpCMon) RightOut () {
//
  x.F (nil, rOut)
}
