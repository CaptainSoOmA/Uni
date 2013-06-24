package phil

// (c) Christian Maurer   v. 130526 - license see murus.go

// >>> Fair Algorithm of Dijkstra
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 98

import
  "sync"
type
  LockSemf struct {
            mutex sync.Mutex
            plate []sync.Mutex
                  }


func (x *LockSemf) test (p uint) {
//
  if stat[p] == hungry &&
     (stat[left(p)] == dining && stat[right(p)] == satisfied ||
      stat[left(p)] == satisfied && stat[right(p)] == dining) {
    changeStatus (p, starving)
  }
  if (stat[p] == hungry || stat[p] == starving) &&
   ! (stat[left(p)] == dining || stat[left(p)] == starving) &&
   ! (stat[right(p)] == dining || stat[right(p)] == starving) {
    changeStatus (p, dining)
    x.plate[p].Unlock()
  }
}


func NewLockSemf () *LockSemf {
//
  x:= new (LockSemf)
  x.plate = make ([]sync.Mutex, nPhilos)
  for p:= uint(0); p < nPhilos; p++ {
    x.plate[p].Lock()
  }
  return x
}


func (x *LockSemf) Lock (p uint) {
//
  x.mutex.Lock()
  changeStatus (p, hungry)
  x.test (p)
  x.mutex.Unlock()
  x.plate[p].Lock()
}


func (x *LockSemf) Unlock (p uint) {
//
  x.mutex.Lock()
  changeStatus (p, satisfied)
  x.test (left(p))
  x.test (right(p))
  x.mutex.Unlock()
}
