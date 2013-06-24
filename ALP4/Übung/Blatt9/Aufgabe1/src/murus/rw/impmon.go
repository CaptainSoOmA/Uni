package rw

// (c) Christian Maurer   v. 121030 - license see murus.go

// >>> readers/writers problem: implementation with monitors
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. TODO

import (
  . "murus/obj"
  "murus/mon"
)
type
  ImpMon struct {
                mon.Monitor
                }


func NewMon () *ImpMon {
//
  var m mon.Monitor
  var nR, nW int
  f:= func (a Any, k uint) Any {
        switch k { case rIn:
          for nW > 0 { m.Wait (rIn) }
          nR ++
          m.Signal (rIn)
        case rOut:
          nR --
          if nR == 0 {
            m.Signal (wIn)
          }
        case wIn:
          for nR > 0 || nW > 0 { m.Wait (wIn) }
          nW = 1
        case wOut:
          nW = 0
          if m.Awaited (rIn) {
            m.Signal (rIn)
          } else {
            m.Signal (wIn)
          }
        }
        return nil
      }
  m = mon.New (4, f)
  return &ImpMon { m }
}


func (x *ImpMon) ReaderIn () {
//
  x.F (nil, rIn)
}


func (x *ImpMon) ReaderOut () {
//
  x.F (nil, rOut)
}


func (x *ImpMon) WriterIn () {
//
  x.F (nil, wIn)
}


func (x *ImpMon) WriterOut () {
//
  x.F (nil, wOut)
}
