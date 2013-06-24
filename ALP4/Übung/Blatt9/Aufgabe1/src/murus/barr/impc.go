package barr

// (c) Christian Maurer   v. 130527 - license see murus.go

// >>> Implementation with a monitor
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. // TODO

import
  "murus/cond"
type
  ImpC struct {
            m chan int
     involved,
      waiting uint
              *cond.Imp
              }


func NewC (n uint) *ImpC {
//
  if n < 2 { return nil }
  x:= new (ImpC)
  x.involved = n
  x.m = make (chan int, 1); x.m <- 0
  x.Imp = cond.New (&x.m)
  return x
}


func (x *ImpC) Wait () {
//
  <-x.m
  x.waiting++
  x.Imp.Wait (x.waiting == x.involved)
  x.waiting-- // standard solution
  x.Signal ()
/*
  x.waiting = 0 // solution with broadcast
  x.SignalAll ()
*/
  x.m <- 0
}
