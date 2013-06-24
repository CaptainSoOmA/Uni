package rw

// (c) Christian Maurer   v. 130527 - license see murus.go

// >>> Readers/Writers problem: Solution with conditions
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. TODO

import
  "murus/cond"
type
  ImpC struct {
       nR, nW int
            m chan int
     okR, okW cond.Condition
              }


func NewC () *ImpC {
//
  x:= new (ImpC)
  x.m = make (chan int, 1); x.m <- 0
  x.okR, x.okW = cond.New (&x.m), cond.New (&x.m)
  return x
}


func (x *ImpC) ReaderIn () {
//
  <-x.m
  x.okR.Wait (x.nW == 0)
  x.nR ++
  x.okR.Signal ()
  x.m <- 0
}


func (x *ImpC) ReaderOut () {
//
  <-x.m
  x.nR --
  if x.nR == 0 {
    x.okW.Signal ()
  }
  x.m <- 0
}


func (x *ImpC) WriterIn () {
//
  <-x.m
  x.okW.Wait (x.nR == 0 && x.nW == 0)
  x.nW = 1
  x.m <- 0
}


func (x *ImpC) WriterOut () {
//
  <-x.m
  x.nW = 0
  if ! x.okR.Signal () {
    x.okW.Signal ()
  }
  x.m <- 0
}
