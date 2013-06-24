package phil

// (c) Christian Maurer   v. 130526 - license see murus.go

// >>> Solution with critical sections
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 98

import (
  . "murus/obj"
  . "murus/cs"
)
type
  LockCS struct {
                CriticalSection
                }


func NewLockCS () *LockCS {
//
  nForks:= make ([]uint, nPhilos)
  for n:= uint(0); n < nPhilos; n++ {
    nForks[n] = 2
  }
  c:= func (p uint) bool {
        return nForks[p] == 2
      }
  l:= func (a Any, p uint) {
        nForks[left(p)] --
        nForks[right(p)] --
      }
  u:= func (a Any, p uint) {
        nForks[left(p)] ++
        nForks[right(p)] ++
      }
  return &LockCS { New (nPhilos, c, l, u) }
}


func (x *LockCS) Lock (p uint) {
//
  changeStatus (p, hungry)
  x.Enter (p, nil)
  changeStatus (p, dining)
}


func (x *LockCS) Unlock (p uint) {
//
  changeStatus (p, satisfied)
  x.Leave (p, nil)
}
