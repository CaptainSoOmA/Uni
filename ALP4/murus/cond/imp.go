package cond

// (c) Christian Maurer   v. 130527 - license see murus.go

type
  Imp struct {
           c chan int
           m *chan int
             }


func New (m *chan int) *Imp {
//
  return &Imp { make (chan int), m }
}


func (x *Imp) Wait (b bool) {
//
  if ! b { // for ! b { // ? TODO
    *x.m <- 0
    <-x.c
    <-*x.m
  }
}


func (x *Imp) Signal () bool {
//
  select {
  case x.c <- 0:
    return true
  default:
  }
  return false
}


func (x *Imp) SignalAll() {
//
  for x.Signal () {
  }
}


/* func (x *Imp) Awaited() bool { // ? TODO
//
  select {
  case n:= <-x.c: // U ???
    x.c <- n
    return true
  default:
  }
  return false
} */


// func init () { var m chan int = make (chan int); var _ Condition = New (&m) }
