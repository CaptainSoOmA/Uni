package rw

// (c) Christian Maurer   v. 130506 - license see murus.go

// >>> 1st readers/writers problem with additive semaphores
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 75

import
  "murus/asem"
const
  m = 19
type
  ImpAsem struct {
                 asem.AddSemaphore
                 }


func NewAsem () *ImpAsem {
//
  return &ImpAsem { asem.New (m) }
}


func (x *ImpAsem) ReaderIn () {
//
  x.AddSemaphore.P (1)
}


func (x *ImpAsem) ReaderOut () {
//
  x.AddSemaphore.V (1)
}


func (x *ImpAsem) WriterIn () {
//
  x.AddSemaphore.P (m)
}


func (x *ImpAsem) WriterOut () {
//
  x.AddSemaphore.V (m)
}
