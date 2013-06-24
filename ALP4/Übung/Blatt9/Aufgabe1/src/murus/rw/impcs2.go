package rw

// (c) Christian Maurer   v. 120909 - license see murus.go

// >>> 2nd readers/writers problem: Implementation with critical sections

import (
  . "murus/obj"
  "murus/cs"
)
type
  ImpCS2 struct {
                cs.CriticalSection
                }


func NewCS2 () *ImpCS2 {
//
  var s cs.CriticalSection
  var nR, nW uint
  c:= func (k uint) bool {
        if k == reader {
          return nW == 0 && ! s.Blocked (writer)
        }
        return nR == 0 && nW == 0 // writer
      }
  e:= func (a Any, k uint) {
        if k == reader {
          nR ++
        } else { // writer
          nW = 1
        }
      }
  l:= func (a Any, k uint) {
        if k == reader {
          nR --
        } else { // writer
          nW = 0
        }
      }
  s = cs.New (2, c, e, l)
  return &ImpCS2 { s }
}


func (x *ImpCS2) ReaderIn () {
//
  x.Enter (reader, nil)
}


func (x *ImpCS2) ReaderOut () {
//
  x.Leave (reader, nil)
}


func (x *ImpCS2) WriterIn () {
//
  x.Enter (writer, nil)
}


func (x *ImpCS2) WriterOut () {
//
  x.Leave (writer, nil)
}
