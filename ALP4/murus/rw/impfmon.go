package rw

// (c) Christian Maurer   v. 130130 - license see murus.go

// >>> readers/writers problem: implementation with far monitors
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, Abschnitt 8.3

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
  var nR, nW int
  c:= func (a Any, k uint) bool {
        switch k { case rIn:
          return nW == 0
        case wIn:
          return nR == 0 && nW == 0
        }
        return true
      }
  f:= func (a Any, k uint) Any {
        switch k { case rIn:
          nR ++
        case rOut:
          nR --
        case wIn:
          nW = 1
        case wOut:
          nW = 0
        }
        return true
      }
  return &ImpFMon { fmon.New (true, 4, f, c, s, p) }
}


func (x *ImpFMon) ReaderIn () {
//
  x.F (true, rIn)
}


func (x *ImpFMon) ReaderOut () {
//
  x.F (true, rOut)
}


func (x *ImpFMon) WriterIn () {
//
  x.F (true, wIn)
}


func (x *ImpFMon) WriterOut () {
//
  x.F (true, wOut)
}
