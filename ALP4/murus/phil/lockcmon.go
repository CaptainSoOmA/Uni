package phil

// (c) Christian Maurer   v. 130526 - license see murus.go

// >>> Solution with conditioned monitor
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 164

import (
  . "murus/obj"
  "murus/mon"
)
type
  LockCMon struct {
                  mon.Monitor
                  }


func NewLockCMon () *LockCMon {
//
  nForks:= make ([]uint, nPhilos)
  for p:= uint(0); p < nPhilos; p++ {
    nForks[p] = 2
  }
  c:= func (a Any, i uint) bool {
        if i == lock {
          p:= a.(uint) // p-th philosopher
          return nForks[p] == 2
        }
        return true // unlock
      }
  f:= func (a Any, i uint) Any {
        p:= a.(uint) // p-th philosopher
        if i == lock {
          nForks[left(p)] --
          nForks[right(p)] --
        } else { // unlock
          nForks[left(p)] ++
          nForks[right(p)] ++
        }
        return p
      }
  return &LockCMon { mon.NewC (nPhilos, f, c) }
}


func (x *LockCMon) Lock (p uint) {
//
  changeStatus (p, hungry)
  x.F (p, lock)
  changeStatus (p, dining)
}


func (x *LockCMon) Unlock (p uint) {
//
  changeStatus (p, satisfied)
  x.F (p, unlock)
}
