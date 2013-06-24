package phil

// (c) Christian Maurer   v. 130526 - license see murus.go

// >>> Naive solution with deadlock:
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 96

import
  "sync"
type
  LockNaiv struct {
             fork []sync.Mutex
                  }


func NewLockNaiv () *LockNaiv {
//
  x:= new (LockNaiv)
  x.fork = make ([]sync.Mutex, nPhilos)
  return x
}


func (x *LockNaiv) Lock (p uint) {
//
  changeStatus (p, hungry)
  x.fork[left (p)].Lock()
  changeStatus (p, hasLeftFork)
  x.fork[p].Lock()
  changeStatus (p, dining)
}


func (x *LockNaiv) Unlock (p uint) {
//
  changeStatus (p, dining)
  x.fork[p].Unlock()
  x.fork[left (p)].Unlock()
}
