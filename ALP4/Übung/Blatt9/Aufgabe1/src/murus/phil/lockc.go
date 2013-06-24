package phil

// (c) Christian Maurer   v. 130527 - license see murus.go

// >>> solution with conditions

import (
//  . "murus/obj"
  "murus/cond"
)
type
  LockC struct {
        nForks []uint
             m chan int
             c []cond.Condition
               }


func NewLockC () *LockC {
//
  x:= new (LockC)
  x.nForks = make ([]uint, nPhilos)
  x.m = make (chan int, 1); x.m <- 0
  x.c = make ([]cond.Condition, nPhilos)
  for p:= uint(0); p < nPhilos; p++ {
    x.nForks[p] = 2
    x.c[p] = cond.New (&x.m)
  }
  return x
}


func (x *LockC) Lock (p uint) {
//
  <-x.m
  changeStatus (p, hungry)
  changeStatus (p, starving)
  x.c[p].Wait (x.nForks[p] == 2)
  x.nForks[left(p)] --
  x.nForks[right(p)] --
  changeStatus (p, dining)
  x.m <- 0
}


func (x *LockC) Unlock (p uint) {
//
  <-x.m
  changeStatus (p, satisfied)
  x.nForks[left(p)] ++
  x.nForks[right(p)] ++
  if x.nForks[left(p)] == 2 {
    if ! x.c[left(p)].Signal () {
      if x.nForks[right(p)] == 2 {
        x.c[right(p)].Signal ()
      }
    }
  }
  x.m <- 0
}
