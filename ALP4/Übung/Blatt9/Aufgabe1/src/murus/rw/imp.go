package rw

// (c) Christian Maurer   v. 120330 - license see murus.go

// >>> 1st readers/writers problem (readers' preference)
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. TODO

import
  . "sync"
type
  Imp struct {
          nR int
       mutex,
          rw Mutex
             }


func New () *Imp {
//
  return new (Imp)
}


func (x *Imp) ReaderIn () {
//
  x.mutex.Lock ()
  x.nR ++
  if x.nR == 1 {
    x.rw.Lock ()
  }
  x.mutex.Unlock ()
}


func (x *Imp) ReaderOut () {
//
  x.mutex.Lock ()
  x.nR --
  if x.nR == 0 {
    x.rw.Unlock ()
  }
  x.mutex.Unlock ()
}


func (x *Imp) WriterIn () {
//
  x.rw.Lock ()
}


func (x *Imp) WriterOut () {
//
  x.rw.Unlock ()
}
