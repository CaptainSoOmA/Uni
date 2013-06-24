package lr

// (c) Christian Maurer   v. 120920 - license see murus.go

// >>> left/right problem: implementation with critical sections
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 93

import (
  "murus/ker"
  . "murus/obj"
  "murus/cs"
)
type
  ImpB struct {
              cs.CriticalSection
              }


func NewB (l, r uint) *ImpB {
//
  if l == 0 { l = ker.MaxNat }
  if r == 0 { r = ker.MaxNat }
  var nL, nR, rL, rR, zL, zR uint
  var s cs.CriticalSection
  c:= func (k uint) bool {
        if k == left {
          return nR == 0 && (! s.Blocked (right) || rL < zL)
        }
        return nL == 0 && (! s.Blocked (left) || rR < zR) // right
      }
  e:= func (A Any, k uint) {
        if k == left {
          nL ++
          rL ++
          rR = 0
        } else { // right
          nR ++
          rR ++
          rL = 0
        }
      }
  a:= func (A Any, k uint) {
        if k == left {
          nL --
        } else { // right
          nR --
        }
      }
  zL, zR = l, r
  s = cs.New (2, c, e, a)
  return &ImpB { s }
}


func (x *ImpB) LeftIn () {
//
  x.Enter (left, nil)
}


func (x *ImpB) LeftOut () {
//
  x.Leave (left, nil)
}


func (x *ImpB) RightIn () {
//
  x.Enter (right, nil)
}


func (x *ImpB) RightOut () {
//
  x.Leave (right, nil)
}
