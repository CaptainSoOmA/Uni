package lr

// (c) Christian Maurer   v. 130409 - license see murus.go

// >>> left/right problem: implementation with critical sections
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 85 ff.

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
  var nL, nR uint
  c:= func (k uint) bool {
        if k == left {
          return nR == 0
        }
        return nL == 0 // right
      }
  e:= func (a Any, k uint) {
        if k == left {
          nL++
        } else { // right
          nR++
        }
      }
  l:= func (X Any, k uint) {
        if k == left {
          nL--
        } else { // right
          nR--
        }
      }
  return &ImpCS { cs.New (2, c, e, l) }
}


func (x *ImpCS) LeftIn () {
//
  x.Enter (left, nil)
}


func (x *ImpCS) LeftOut () {
//
  x.Leave (left, nil)
}


func (x *ImpCS) RightIn () {
//
  x.Enter (right, nil)
}


func (x *ImpCS) RightOut () {
//
  x.Leave (right, nil)
}
