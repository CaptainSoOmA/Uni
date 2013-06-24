package rw

// (c) Christian Maurer   v. 130130 - license see murus.go

// >>> readers/writers problem: implementation with critical sections
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 92

import (
  . "murus/obj"
  "murus/cs"
)
type
  ImpCS struct {
               cs.CriticalSection
               }


func NewCS () *ImpCS {
//
  var nR, nW uint
  c:= func (k uint) bool {
        if k == reader {
          return nW == 0
        }
        return nR == 0 && nW == 0 // writer
      }
  e:= func (a Any, k uint) {
        if k == reader {
          nR++
        } else { // writer
          nW = 1
        }
      }
  a:= func (a Any, k uint) {
        if k == reader {
          nR --
        } else { // writer
          nW = 0
        }
      }
  return &ImpCS { cs.New (2, c, e, a) }
}


func (x *ImpCS) ReaderIn () {
//
  x.Enter (reader, nil)
}


func (x *ImpCS) ReaderOut () {
//
  x.Leave (reader, nil)
}


func (x *ImpCS) WriterIn () {
//
  x.Enter (writer, nil)
}


func (x *ImpCS) WriterOut () {
//
  x.Leave (writer, nil)
}
