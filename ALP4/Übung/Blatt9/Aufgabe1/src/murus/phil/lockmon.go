package phil

// (c) Christian Maurer   v. 130526 - license see murus.go

// >>> monitor solution

import (
  . "murus/obj"
  "murus/mon"
)
type
  LockMon struct {
                 mon.Monitor
                 }


func NewLockMon () *LockMon {
//
  var m mon.Monitor
  nForks:= make ([]uint, nPhilos)
  for p:= uint(0); p < nPhilos; p++ {
    nForks[p] = 2
  }
  f:= func (a Any, i uint) Any {
        p:= a.(uint)
        if i == lock {
          changeStatus (p, starving)
          for nForks[p] < 2 {
            m.Wait (p)
          }
          nForks[left(p)] --
          nForks[right(p)] --
        } else { // k == unlock
          nForks[left(p)] ++
          nForks[right(p)] ++
          if nForks[left(p)] == 2 {
            m.Signal(left(p))
          }
          if nForks[right(p)] == 2 {
            m.Signal(right(p))
          }
        }
        return nil
      }
  m = mon.New (nPhilos, f)
  return &LockMon { m }
}


func (x *LockMon) Lock (p uint) {
//
  changeStatus (p, hungry)
  x.F (p, lock)
  changeStatus (p, dining)
}


func (x *LockMon) Unlock (p uint) {
//
  changeStatus (p, satisfied)
  x.F (p, unlock)
}
