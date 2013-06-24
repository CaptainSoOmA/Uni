package phil

// (c) Christian Maurer   v. 130526 - license see murus.go

// >>> Unsymmetric case:
//     Störung der Symmetrie dadurch, dass manche (aber nicht alle)
//     Philosophen zuerst die linke Gabel aufnehmen
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 96

import
  "sync"
type
  LockUns struct {
            fork []sync.Mutex
                 }


func NewLockUns () *LockUns {
//
  x:= new (LockUns)
  x.fork = make ([]sync.Mutex, nPhilos)
  return x
}


func (x *LockUns) Lock (p uint) {
//
  changeStatus (p, hungry)
  if p % 2 == 1 {
//  if p == 0 {
    x.fork [left (p)].Lock()
    changeStatus (p, hasLeftFork)
    x.fork [p].Lock()
  } else {
    x.fork [p].Lock()
    changeStatus (p, hasRightFork)
    x.fork [left (p)].Lock()
  }
  changeStatus (p, dining)
}


func (x *LockUns) Unlock (p uint) {
//
  changeStatus (p, dining)
  changeStatus (p, satisfied)
  x.fork[p].Unlock()
  x.fork[left (p)].Unlock()
}
