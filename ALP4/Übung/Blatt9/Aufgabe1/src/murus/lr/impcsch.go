package lr

// (c) Christian Maurer   v. 130513 - license see murus.go

// >>> left/right problem: implementation with critical sections
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 85 ff.

import (
  . "murus/obj"
  "murus/cs"
)
type
  ImpCSCh struct {
                 cs.CriticalSection
                 }


func NewCSCh () *ImpCSCh {
//
  var nL, nR uint
  c:= func (k uint) bool {
        if k == left {
          return nR == 0
        }
        return nL == 0 // k == right
      }
  e:= func (a Any, k uint) {
        if k == left {
          nL++
        } else { // right
          nR++
        }
      }
  l:= func (a Any, k uint) {
        if k == left {
          nL--
        } else { // right
          nR--
        }
      }
  return &ImpCSCh { cs.NewCh (2, c, e, l) }
}


func (x *ImpCSCh) LeftIn () {
//
  x.Enter (left, nil)
}


func (x *ImpCSCh) LeftOut () {
//
  x.Leave (left, nil)
}


func (x *ImpCSCh) RightIn () {
//
  x.Enter (right, nil)
}


func (x *ImpCSCh) RightOut () {
//
  x.Leave (right, nil)
}
