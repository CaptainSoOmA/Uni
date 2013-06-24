package asem

// (c) Christian Maurer   v. 120909 - license see murus.go

import (
  . "sync"
  "murus/sem"
)
const
  M = 20
type
  Imp struct {
     val int // number of processes allowed to use the critical section,
             // which shall be protected by the semaphore
          me Mutex
           b [M]sem.Semaphore
          nB [M]int
             }


func New (n uint) *Imp {
//
  x:= new (Imp)
  x.val = int(n)
  for i:= 0; i < M; i++ {
    x.b[i] = sem.New (0)
  }
  return x
}


func (x *Imp) P (n uint) {
//
  x.me.Lock()
  if x.val >= int(n) {
    x.val -= int(n)
    x.me.Unlock()
  } else {
    x.nB[n]++
    x.me.Unlock()
    x.b[n].P()
  }
}


func (x *Imp) V (n uint) {
//
  x.me.Lock()
  x.val += int(n)
  i:= x.val
  for i > 0 {
    for x.nB[i] > 0 && i < x.val {
      x.val -= i // x.val--
      x.nB[i]--
      x.b[i].V()
    }
    i --
  }
  x.me.Unlock()
}
