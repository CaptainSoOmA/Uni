package cs

// (c) Christian Maurer   v. 120910 - license see murus.go

import (
  "sync"
  . "murus/obj"
  "murus/perm"
)
type
  Imp struct {
          nP uint // number of process classes
       baton sync.Mutex
           s []sync.Mutex // on which goroutines are blocked, if ! c
          nB []uint // numbers of goroutines, that are blocked on these semaphores
           c CondSpectrum // conditions to enter the critical section
     in, out OpSpectrum // operations in the entry and exit protocols
           p *perm.Imp // random permutation
             }


func New (n uint, c CondSpectrum, e, l OpSpectrum) *Imp {
//
  if n == 0 { return nil }
  x:= new (Imp)
  x.nP = n
  x.s = make ([]sync.Mutex, x.nP)
  x.nB = make ([]uint, x.nP)
  for k:= uint(0); k < x.nP; k++ {
    x.s [k].Lock()
  }
  x.c, x.in, x.out = c, e, l
  x.p = perm.New (x.nP)
  return x
}


func (x *Imp) vall () {
//
  x.p.Permute ()
  var k uint
  for i:= uint(0); i < x.nP; i++ {
    k = x.p.F (i)
    if x.c (k) && x.nB [k] > 0 {
      x.nB [k] --
      x.s [k].Unlock()
      return
    }
  }
  x.baton.Unlock()
}


func (x *Imp) Blocked (k uint) bool {
//
  if k >= x.nP { return false }
  return x.nB [k] > 0
}


func (x *Imp) Enter (k uint, a Any) {
//
  if k >= x.nP { return }
  x.baton.Lock()
  if ! x.c (k) {
    x.nB [k] ++
    x.baton.Unlock()
    x.s [k].Lock()
  }
  x.in (a, k)
  x.vall ()
}


func (x *Imp) Leave (k uint, a Any) {
//
  if k >= x.nP { return }
  x.baton.Lock()
  x.out (a, k)
  x.vall ()
}
