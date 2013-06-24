package lr

// (c) Christian Maurer   v. 130527 - license see murus.go

// >>> Left/Right problem: Solution with conditions
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. TODO

import
  "murus/cond"
type
  ImpC struct {
       nL, nR int
            m chan int
     okL, okR cond.Condition
              }


func NewC () *ImpC {
//
  x:= new (ImpC)
  x.m = make (chan int, 1); x.m <- 0
  x.okL, x.okR = cond.New (&x.m), cond.New (&x.m)
  return x
}


func (x *ImpC) LeftIn () {
//
  <-x.m
  x.okL.Wait (x.nR == 0)
  x.nL ++
  x.okL.Signal ()
  x.m <- 0
}


func (x *ImpC) LeftOut () {
//
  <-x.m
  x.nL --
  if x.nL == 0 {
    if ! x.okR.Signal () {
      x.okL.Signal ()
    }
  }
  x.m <- 0
}


func (x *ImpC) RightIn () {
//
  <-x.m
  x.okR.Wait (x.nL == 0)
  x.nR ++
  x.okR.Signal ()
  x.m <- 0
}


func (x *ImpC) RightOut () {
//
  <-x.m
  x.nR --
  if x.nR == 0 {
    if ! x.okL.Signal () {
      x.okR.Signal ()
    }
  }
  x.m <- 0
}
