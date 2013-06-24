package csp

// (c) Christian Maurer   v. 130525 - license see murus.go

import (
  . "murus/obj"
  "murus/perm"
  "murus/sem"
)
type
  Imp struct {
          nP uint // number of process classes
       baton sem.Semaphore
           s []sem.Semaphore // on which goroutines are blocked, if ! c
          nB []uint // numbers of goroutines, that are blocked on these semaphores
           c PredSpectrum // predicates to enter the critical section
     in, out OpSpectrum // operations in the entry and exit protocols
           p *perm.Imp // random permutation
             }


func New (n uint, p PredSpectrum, e, l OpSpectrum) *Imp {
//
  if n == 0 { return nil }
  x:= new (Imp)
  x.nP = n
  x.s = make ([]sem.Semaphore, x.nP)
  x.nB = make ([]uint, x.nP)
  for k:= uint(0); k < x.nP; k++ {
    x.s[k] = sem.NewGSel (0)
  }
  x.c, x.in, x.out = p, e, l
  x.p = perm.New (x.nP)
  return x
}


func (x *Imp) vall (a Any) {
//
  x.p.Permute ()
  var k uint
  for i:= uint(0); i < x.nP; i++ {
    k = x.p.F (i)
    if x.c (a, k) && x.nB[k] > 0 {
      x.nB[k] --
      x.s[k].V()
      return
    }
  }
  x.baton.V()
}


func (x *Imp) Blocked (k uint) bool {
//
  if k >= x.nP { return false }
  return x.nB[k] > 0
}


func (x *Imp) Enter (k uint, a Any) {
//
  if k >= x.nP { return }
  x.baton.P()
  if ! x.c (a, k) {
    x.nB[k] ++
    x.baton.V()
    x.s[k].P()
  }
  x.in (a, k)
  x.vall (a)
}


func (x *Imp) Leave (k uint, a Any) {
//
  if k >= x.nP { return }
  x.baton.P()
  x.out (a, k)
  x.vall (a)
}
