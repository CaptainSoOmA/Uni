package phil

// (c) Christian Maurer   v. 130526 - license see murus.go

// >>> Unfair solution with semaphores, aushungerungsgefährdet
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 95 ff.

import
  "sync"
type
  LockSemuf struct {
             mutex sync.Mutex
             plate []sync.Mutex
                   }


func (x *LockSemuf) test (p uint) {
//
  if stat[p] == hungry &&
     stat[left(p)] != dining && stat[right(p)] != dining {
    changeStatus (p, dining)
    x.plate[p].Unlock()
  }
}


func NewLockSemuf () *LockSemuf {
//
  x:= new (LockSemuf)
  x.plate = make ([]sync.Mutex, nPhilos)
  for p:= uint(0); p < nPhilos; p++ {
    x.plate[p].Lock()
  }
  return x
}


func (x *LockSemuf) Lock (p uint) {
//
  x.mutex.Lock()
  changeStatus (p, hungry)
  x.test (p)
  x.mutex.Unlock()
  x.plate[p].Lock()
}


func (x *LockSemuf) Unlock (p uint) {
//
  x.mutex.Lock()
  changeStatus (p, satisfied)
  x.test (left (p))
  x.test (right (p))
  x.mutex.Unlock()
}
