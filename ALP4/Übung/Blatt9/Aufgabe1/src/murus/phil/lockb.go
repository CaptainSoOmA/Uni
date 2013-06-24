package phil

// (c) Christian Maurer   v. 130526 - license see murus.go

// >>> Bounded case:
//     At most m - 1 philosophers are allowed to take place at the table
//     at the same time, where m is the number of participating philsophers.
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 96

import (
  "sync"
  . "murus/sem"
)
type
  LockB struct {
      takeSeat Semaphore
          fork []sync.Mutex
               }


func NewLockB () *LockB {
//
  x:= new (LockB)
  x.takeSeat = New (nPhilos - 1)
  x.fork = make ([]sync.Mutex, nPhilos)
  return x
}


func (x *LockB) Lock (p uint) {
//
  changeStatus (p, hungry)
  x.takeSeat.P ()
  x.fork[left (p)].Lock ()
  changeStatus (p, hasRightFork)
  x.fork[p].Lock ()
  changeStatus (p, dining)
}


func (x *LockB) Unlock (p uint) {
//
  changeStatus (p, satisfied)
  x.fork[p].Unlock()
  x.fork[left (p)].Unlock()
  x.takeSeat.V ()
}
