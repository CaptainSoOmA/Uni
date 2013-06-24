package rw

// (c) Christian Maurer   v. 130130 - license see murus.go

// >>> 1st readers/writers problem (readers' preference)
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 75

import
  "murus/sem"
type
  ImpSem struct {
             nR int
          mutex,
             rw sem.Semaphore
                }


func NewSem () *ImpSem {
//
  return &ImpSem { mutex: sem.New (1), rw: sem.New (1)  }
}


func (x *ImpSem) ReaderIn () {
//
  x.mutex.P ()
  x.nR ++
  if x.nR == 1 {
    x.rw.P ()
  }
  x.mutex.V ()
}


func (x *ImpSem) ReaderOut () {
//
  x.mutex.P ()
  x.nR --
  if x.nR == 0 {
    x.rw.V ()
  }
  x.mutex.V ()
}


func (x *ImpSem) WriterIn () {
//
  x.rw.P ()
}


func (x *ImpSem) WriterOut () {
//
  x.rw.V ()
}
