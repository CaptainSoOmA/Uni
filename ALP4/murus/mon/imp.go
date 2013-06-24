package mon

// (c) Christian Maurer   v. 130527 - license see murus.go

// >>> Nichtsequentielle Programmierung mit Go 1 kompakt, S. 154 ff.

import (
//  "sync"
  . "murus/obj"; "murus/perm"
)
type
  Imp struct {
        nFns uint // number of monitor functions
//          me sync.Mutex // monitor entry queue
          me chan int // monitor entry queue
//           s []sync.Mutex // condition variable queues
           s []chan int // condition variable queues
          nB []uint // numbers of goroutines blocked on s
//      urgent sync.Mutex // urgent queue
      urgent chan int // urgent queue
          nU uint // number of goroutines blocked on urgent
           f FuncSpectrum
           p PredSpectrum
        cond bool // true, iff monitor is conditioned
          pm *perm.Imp // indeterminism
             }


func New (n uint, f FuncSpectrum /* , p PredSpectrum */) *Imp {
//
  if n == 0 { return nil }
  x:= new (Imp)
  x.nFns = n
  x.me = make (chan int, 1); x.me <- 0
//  x.s = make ([]sync.Mutex, x.nFns)
  x.s = make ([]chan int, x.nFns)
  for i:= uint(0); i < x.nFns; i++ {
//    x.s[i].Lock()
    x.s[i] = make (chan int, 1)
  }
  x.nB = make ([]uint, x.nFns)
//  x.urgent.Lock()
  x.urgent = make (chan int, 1)
  x.f = f
//  if p == nil {
    x.p = TrueSp
//  } else {
//    x.cond = true
//    x.p = p
//  }
  x.pm = perm.New (x.nFns)
  return x
}


func NewC (n uint, f FuncSpectrum, p PredSpectrum) *Imp {
//
  x:= New (n, f)
  x.p = p
  x.cond = true
  return x
}


func (x *Imp) Wait (i uint) {
//
  if i >= x.nFns { WrongUintParameterPanic ("mon.Wait", x, i) }
  x.nB[i] ++
  if x.nU > 0 {
//    x.urgent.Unlock()
    x.urgent <- 0
  } else {
//    x.me.Unlock()
    x.me <- 0
  }
//  x.s[i].Lock()
  <-x.s[i]
  x.nB[i] --
}


func (x *Imp) Awaited (i uint) bool {
//
  if i >= x.nFns { WrongUintParameterPanic ("mon.Awaited", x, i) }
  return x.nB[i] > 0
}


func (x *Imp) Signal (i uint) {
//
  if i >= x.nFns { WrongUintParameterPanic ("mon.Signal", x, i) }
  if x.nB[i] > 0 {
    x.nU ++
//    x.s[i].Unlock()
    x.s[i] <- 0
//    x.urgent.Lock()
    <-x.urgent
    x.nU --
  }
}


func (x *Imp) SignalAll (i uint) {
//
  if i >= x.nFns { return }
  if i >= x.nFns { WrongUintParameterPanic ("mon.SignalAll", x, i) }
  for {
    if x.nB[i] == 0 { break }
    x.nU ++
//    x.s[i].Unlock()
    x.s[i] <- 0
//    x.urgent.Lock()
    <-x.urgent
    x.nU --
  }
}


func (x *Imp) F (a Any, i uint) Any {
//
  if i >= x.nFns { WrongUintParameterPanic ("mon.F", x, i) }
//  x.me.Lock ()
  <-x.me
  if x.cond {
    for ! x.p (a, i) {
      x.Wait (i)
    }
  }
  b:= x.f (a, i)
  if x.cond {
    x.pm.Permute ()
    for j:= uint(0); j < x.nFns; j++ {
      x.Signal (x.pm.F (j))
    }
  }
  if x.nU > 0 {
//    x.urgent.Unlock()
    x.urgent <- 0
  } else {
//    x.me.Unlock ()
    x.me <- 0
  }
  return b
}


// experimental
func (x *Imp) S (a Any, i uint, c chan Any) {
//
  c <- x.F (a, i)
}
