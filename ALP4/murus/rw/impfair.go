package rw

// (c) Christian Maurer   v. 120909 - license see murus.go

// >>> readers/writers problem: fair solution
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 84

import (
  . "murus/obj"
  "murus/cs"
)
type
  ImpFair struct {
                 cs.CriticalSection
                 }


func NewFair () *ImpFair {
//
  var s *cs.Imp
  var nR, nW uint
  var lastR bool
  c:= func (k uint) bool {
        if k == reader {
          return          nW == 0 && (! s.Blocked (writer) || ! lastR)
        }
        return nR == 0 && nW == 0 && (! s.Blocked (reader) || lastR) // writer
      }
  e:= func (X Any, k uint) {
        if k == reader {
          nR ++
          lastR = true
        } else { // writer
          nW ++
          lastR = false
        }
      }
  l:= func (X Any, k uint) {
        if k == reader {
          nR --
        } else { // writer
          nW --
        }
      }
  s = cs.New (2, c, e, l)
  return &ImpFair { s }
}


func (x *ImpFair) ReaderIn () {
//
  x.Enter (reader, 0)
}


func (x *ImpFair) ReaderOut () {
//
  x.Leave (reader, 1)
}


func (x *ImpFair) WriterIn () {
//
  x.Enter (writer, 0)
}


func (x *ImpFair) WriterOut () {
//
  x.Leave (writer, 1)
}
