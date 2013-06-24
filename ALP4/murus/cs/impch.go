package cs

// (c) Christian Maurer   v. 130513 - license see murus.go

import (
  "murus/ker"
  . "murus/obj"
)
type
  ImpCh struct {
            nP uint // number of process classes
        ci, co []chan Any // channels for messages to enter and leave
               }


func NewCh (n uint, c CondSpectrum, e, l OpSpectrum) *ImpCh {
//
  if n == 0 { return nil }
  x:= new (ImpCh)
  x.nP = n
  x.ci, x.co = make ([]chan Any, x.nP), make ([]chan Any, x.nP)
  for i:= uint(0); i < x.nP; i++ {
    x.ci[i], x.co[i] = make (chan Any), make (chan Any)
  }
  go func () {
    for {
      for k:= uint(0); k < x.nP; k++ {
        select {
        case a:= <-When (c (k), x.ci[k]):
          e (a, k)
        case a:= <-x.co[k]:
          l (a, k)
        default:
        }
      }
      ker.Msleep (10)
    }
  }()
  return x
}


func (x *ImpCh) Blocked (k uint) bool {
//
  if k >= x.nP { return false }
  return false // TODO
}


func (x *ImpCh) Enter (k uint, a Any) {
//
  if k >= x.nP { return }
  x.ci[k] <- a
}


func (x *ImpCh) Leave (k uint, a Any) {
//
  if k >= x.nP { return }
  x.co[k] <- a
}
