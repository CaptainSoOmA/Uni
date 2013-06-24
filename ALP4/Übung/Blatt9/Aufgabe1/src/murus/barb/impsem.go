package barb

// (c) Christian Maurer   v. 130424 - license see murus.go

import
  "sync"
type
  ImpSem struct {
        waiting,
          mutex sync.Mutex
           n, k uint
                }


func NewSem () *ImpSem {
//
  x:= new (ImpSem)
  x.waiting.Lock ()
  return x
}


func (x *ImpSem) Customer () {
//
  x.mutex.Lock ()
  x.n ++
  if x.n == 1 {
    x.waiting.Unlock ()
  }
  x.mutex.Unlock ()
}


func (x *ImpSem) Barber () {
//
  if x.k == 0 {
    x.waiting.Lock ()
  }
  x.mutex.Lock ()
  x.n --
  x.k = x.n
  x.mutex.Unlock ()
}
