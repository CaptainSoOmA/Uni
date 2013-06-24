package rw

// (c) Christian Maurer   v. 121030 - license see murus.go

// >>> readers/writers problem: implementation with monitors
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. TODO

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
  var nR, nW int
  f:= func (a Any, i uint) Any {
        switch i { case rIn:
          nR++
        case rOut:
          nR--
        case wIn:
          nW++
        case wOut:
          nW--
        }
        return nil
      }
  p:= func (a Any, i uint) bool {
        switch i { case rIn:
          return nW == 0
        case wIn:
          return nR == 0 && nW == 0
        }
        return true
      }
  return &ImpCMon { mon.NewC (4, f, p) }
}


func (x *ImpCMon) ReaderIn () {
//
  x.F (nil, rIn)
}


func (x *ImpCMon) ReaderOut () {
//
  x.F (nil, rOut)
}


func (x *ImpCMon) WriterIn () {
//
  x.F (nil, wIn)
}


func (x *ImpCMon) WriterOut () {
//
  x.F (nil, wOut)
}
