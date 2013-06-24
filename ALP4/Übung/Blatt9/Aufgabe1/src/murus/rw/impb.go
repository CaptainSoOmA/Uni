package rw

// (c) Christian Maurer   v. 130130 - license see murus.go

// >>> 3rd readers/writers problem (number of concurrent readers bounded)

import (
  . "murus/obj"
  "murus/cs"
)
type
  ImpB struct {
              cs.CriticalSection
              }


func NewB (m uint) *ImpB {
//
  var s cs.CriticalSection
  if m < 1 { m = 1 }
  var nR, nW, rR uint
  c:= func (k uint) bool {
        if k == reader {
          return nW == 0 && (s.Blocked (writer) || rR < m)
        }
        return nR == 0 && nW == 0 /* && ! x.s.Blocked (readeer) */ // writer
      }
  e:= func (a Any, k uint) {
        if k == reader {
          nR ++
          rR ++
        } else { // writer
          nW ++
          rR = 0
        }
      }
  l:= func (a Any, k uint) {
        if k == reader {
          nR --
        } else { // writer
          nW --
        }
      }
  s = cs.New (2, c, e, l)
  return &ImpB { s }
}


func (x *ImpB) ReaderIn () {
//
  x.Enter (reader, nil)
}


func (x *ImpB) ReaderOut () {
//
  x.Leave (reader, nil)
}


func (x *ImpB) WriterIn () {
//
  x.Enter (writer, nil)
}


func (x *ImpB) WriterOut () {
//
  x.Leave (writer, nil)
}
